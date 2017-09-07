package homee

import (

	log "github.com/golang/glog"
	"time"
	"net/url"
	"net/http"
	"github.com/gorilla/websocket"


	"encoding/json"

	"github.com/pkg/errors"
	"github.com/connctd/homee/model"
)

type apiConnection struct {
	stateChange chan int
	State int
	UID          string
	AccessToken string
	connection *websocket.Conn
	inByteMsg  chan []byte
	outByteMsg chan []byte
	IncomingMessages chan model.HomeeMessage
	AutoReconnect bool
	lastPong time.Time
	lastPing time.Time

	writeGopherRunning bool
	readGohperRunning bool
	messageGopherRunnging bool
	TimestampConnectionEstablished time.Time
	timestampConnectionClosed time.Time

	PingInterval		time.Duration
	ConnectionTimeout 	time.Duration
	ReconnectionInterval 	time.Duration
}

const (
	CONNECTION_STATE_CONNECTING 		= 1
	CONNECTION_STATE_CONNECTION_ESTABLISHED = 2
	CONNECTION_STATE_CONNECTED 		= 3
	CONNECTION_STATE_CLOSING 		= 4
	CONNECTION_STATE_CLOSED 		= 5
	CONNECTION_STATE_ERROR 			= -1
	CONNECTION_STATE_ERROR_AUTHORIZATION 	= -2
	CONNECTION_STATE_ERROR_TIMEOUT 		= -3

	DEFAULT_PING_INTERVAL 			= 30
	DEFAULT_CONNECTION_TIMEOUT 		= 20
	DEFAULT_RECONNECTION_INTERVAL 		= 10
)

func newConnection(config HomeeConfiguration, reconnect bool) *apiConnection {
	newInstance := &apiConnection{UID:config.UID,
		AccessToken:config.AccessToken,
		AutoReconnect:reconnect,
	}

	newInstance.writeGopherRunning = false
	newInstance.readGohperRunning =false
	newInstance.messageGopherRunnging = false
	newInstance.State = CONNECTION_STATE_CLOSED
	newInstance.inByteMsg 	= make(chan []byte, 100)
	newInstance.outByteMsg  = make(chan []byte, 100)
	newInstance.IncomingMessages = make(chan model.HomeeMessage, 100)
	newInstance.stateChange = make(chan int, 1)
	newInstance.ConnectionTimeout = time.Second * DEFAULT_CONNECTION_TIMEOUT
	newInstance.PingInterval = time.Second * DEFAULT_PING_INTERVAL
	newInstance.ReconnectionInterval = time.Second * DEFAULT_RECONNECTION_INTERVAL
	return newInstance
}

func (c *apiConnection)connect() {
	go c.handleConnection()
	c.stateChange <- CONNECTION_STATE_CONNECTING
}

func (c *apiConnection) Close(){
	c.AutoReconnect = false;
	c.connection.Close()
	c.stateChange <- CONNECTION_STATE_CLOSED
}

func (c *apiConnection) establishConnection() error{
	c.lastPing = time.Now()
	c.lastPong = time.Now()
	baseUrl := c.UID+".hom.ee"
	u := url.URL{Scheme: "wss", Host: baseUrl, Path:"connection",RawQuery:"access_token="+c.AccessToken}
	log.Infof("connecting to '%s'", baseUrl)
	// need to add some headers
	headers := http.Header{}
	headers.Add("Sec-WebSocket-Protocol","v2");
	// generating dialer
	dialer := websocket.DefaultDialer
	// try to connect
	if conn,httpResponse,err := dialer.Dial(u.String(), headers); err != nil {
		//arrr ficken man!
		log.Infof("Unable to connect to '%s', %v", baseUrl, err)
		log.Infof("%v",httpResponse)
		c.stateChange <- CONNECTION_STATE_ERROR
		if (httpResponse != nil) {
			if httpResponse.StatusCode == 302 {
				// 302 usually means redirect, but it indicates bad credentials for homee as well
				c.stateChange <- CONNECTION_STATE_ERROR_AUTHORIZATION
			}
		}
		return errors.New("Unable to connect to '"+c.UID+"'")
	} else {
		// well done!
		c.connection = conn
		c.connection.SetPongHandler(c.handlePongMessage)
		c.connection.SetCloseHandler(c.handleCloseMessage)
		c.stateChange <- CONNECTION_STATE_CONNECTION_ESTABLISHED
		return nil
	}
}

func (c *apiConnection) handleCloseMessage(code int, text string) error {
	log.Infof("Connection is closed. Error = %s (%v)",text,code)
	c.stateChange <- CONNECTION_STATE_CLOSED
	return nil
}

func (c *apiConnection) handlePongMessage(appData string) error {
	c.lastPong = time.Now()
	return nil
}



func (c *apiConnection) handleConnection(){
	defer log.Infof("Connection handling for conncection '%s' has been stopped")
	for {
		select  {
			case state := <- c.stateChange:
				if (state != c.State){
					log.Infof("Connection to '%s' changed from '%s' to '%s'",c.UID,GetStateName(c.State),GetStateName(state))
					c.State = state
					switch state {
					case CONNECTION_STATE_CONNECTING:
						c.establishConnection()
					case CONNECTION_STATE_CONNECTION_ESTABLISHED:
						c.TimestampConnectionEstablished = time.Now()
						go c.write()
						go c.read()
						go c.handleIncomingByteMessages()
						c.stateChange <- CONNECTION_STATE_CONNECTED
					case CONNECTION_STATE_CONNECTED:
						// do nothing
					case CONNECTION_STATE_CLOSED:
						c.timestampConnectionClosed = time.Now()
					case CONNECTION_STATE_ERROR:
						if (c.connection != nil) {
							c.connection.Close()
						}
						c.stateChange <- CONNECTION_STATE_CLOSED
					case CONNECTION_STATE_ERROR_TIMEOUT:
						if (c.connection != nil) {
							c.connection.Close()
						}
						c.stateChange <- CONNECTION_STATE_CLOSED
					case CONNECTION_STATE_ERROR_AUTHORIZATION :
						if (c.connection != nil) {
							c.connection.Close()
						}
						c.stateChange <- CONNECTION_STATE_CLOSED
						return
					default:
						log.Infof("connection state '%v' (%s) will not be handled",c.State,GetStateName(c.State))

					}

				}
			case <-time.After(time.Second):
				if (c.State == CONNECTION_STATE_CLOSED){
					if c.AutoReconnect {
						if time.Now().Sub(c.timestampConnectionClosed) > c.ReconnectionInterval {
							log.Infof("Trying to reconnect to '%s'",c.UID)
							c.stateChange <- CONNECTION_STATE_CONNECTING
						}
					} else {
						// go routines are waiting until channel has new messages
						// send some empty ones that they can be stopped
						c.inByteMsg <- []byte{}
						c.outByteMsg <- []byte{}
						return
					}
				} else if c.State == CONNECTION_STATE_CONNECTED {
					// check whether last pong was received within timout intervall after ping was sent

					//if time.Now().Sub(c.lastPong) > time.Second * DEFAULT_CONNECTION_TIMEOUT {
					if c.lastPing.Sub(c.lastPong) > c.ConnectionTimeout{
						log.Infof("No Pong within %vs after last ping sent!",c.ConnectionTimeout)
						c.stateChange <- CONNECTION_STATE_ERROR_TIMEOUT
					}
					// check whether it's time for sending a ping
					if time.Now().Sub(c.lastPing) > c.PingInterval{
						c.ping()
					}
				}
		}

	}


}

func GetStateName(state int) string {
	switch state {
	case CONNECTION_STATE_CONNECTION_ESTABLISHED:
		return "CONNECTION_ESTABLISHED"
	case CONNECTION_STATE_CONNECTING:
		return "CONNECTING"
	case CONNECTION_STATE_CONNECTED:
		return "CONNECTED"
	case CONNECTION_STATE_CLOSING:
		return "CLOSING"
	case CONNECTION_STATE_CLOSED:
		return "CLOSED"
	case CONNECTION_STATE_ERROR:
		return "ERROR"
	case CONNECTION_STATE_ERROR_TIMEOUT:
		return "ERROR_TIMEOUT"
	case CONNECTION_STATE_ERROR_AUTHORIZATION:
		return "ERROR_AUTHORIZATION"
	default:
		return "UNKNOWN"
	}

}

func (c *apiConnection) sendMessage(message string){
	c.outByteMsg <- []byte(message)
}

func (c *apiConnection) write(){
	defer func() {
		c.writeGopherRunning = false;
	}()


	if (c.writeGopherRunning){
		return;
	}
	c.writeGopherRunning = true;
	for c.State != CONNECTION_STATE_CLOSED  {
		msg := <- c.outByteMsg
		if c.State == CONNECTION_STATE_CLOSED {
			return
		}
		log.Infof("sending to '%s': %s",c.UID,string(msg))
		if err := c.connection.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Errorf("Error during write. %v",err)
			c.stateChange <- CONNECTION_STATE_ERROR
			return
		}
	}
}

func (c *apiConnection) read(){
	defer func() {
		c.readGohperRunning = false;
	}()

	if (c.readGohperRunning){
		return;
	}
	c.readGohperRunning = true;
	for c.State != CONNECTION_STATE_CLOSED  {
		_, message, err := c.connection.ReadMessage()
		if err != nil {
			log.Errorf("Error during read. %v",err)
			if websocket.IsCloseError(err){
				log.Info("Connection is closed")
			}
			c.stateChange <- CONNECTION_STATE_ERROR

			return
		} else {
			c.inByteMsg <- message;
		}
	}
}

func (c *apiConnection) ping(){
	if err := c.connection.WriteControl(websocket.PingMessage, []byte{0x09}, time.Now().Add(time.Second*60)); err != nil {
		log.Errorf("Error during write. %v",err)
		c.stateChange <- CONNECTION_STATE_ERROR
		return
	}
	c.lastPing = time.Now()

}

func (c *apiConnection) handleIncomingByteMessages() {
	if (c.messageGopherRunnging){
		return;
	}
	c.messageGopherRunnging = true
	log.Infof("Message-Gopher for '%s' has been started",c.UID)
	for c.State != CONNECTION_STATE_CLOSED {
		msg := <-c.inByteMsg

		encoded := string(msg)
		if decoded, err := url.QueryUnescape(encoded); err == nil {
			log.Infof("received message : %s", decoded);
			var homeeMsg model.HomeeMessage
			if err := json.Unmarshal([]byte(decoded),&homeeMsg); err == nil {
				c.IncomingMessages <- homeeMsg
			} else {

			}
		} else {
			log.Infof("Failed to url decode message %s", encoded)
		}
	}
	c.messageGopherRunnging = false
}