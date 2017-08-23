package homee

import (
	"net/http"
	log "github.com/golang/glog"

	"strconv"
	"github.com/connctd/homee/model"
)


//would be better to establish something like an event mechanism: things can be created, deleted, edited, etc.. Would be
//stupid to create a function for each event

type ConnectionError struct{
	Url string
	StatusCode int
	StatusMessage string
	Response *http.Response
}

type Client struct {
	config     HomeeConfiguration
	connection *apiConnection

	// nodes
	//Attributes map[int]model.Attribute
//	Nodes map[int]model.Node
	//Things map[int]model.Hom eeThingMapping

}



func New(config HomeeConfiguration) *Client {
	result := Client{config:config, connection:nil}
	//result.Nodes = make(map[int]model.Node)
	//result.Attributes = make(map[int]model.Attribute)
	//result.Things = make(map[int]model.HomeeThingMapping)
	return &result
}

func (c *Client)Connect() {
	log.Infof("establish connection")
	if (c.connection == nil){
		c.connection = newConnection(c.config, true)
	}
	c.connection.connect()
	go c.handleHomeeMessages()
}


func (c *Client) handleHomeeMessages(){
	for {
		msg := <- c.connection.IncomingMessages
		if msg.Nodes != nil {
			c.handleNodesMessage(*msg.Nodes)
		}
		if msg.Node != nil {
			c.handleNodeMessage(*msg.Node)
		}
		if msg.Attribute != nil {
			c.handleAttributeMessage(*msg.Attribute)
		}
		if msg.Code != 0 {
			log.Infof("Received message of type '%s', but handling is not implemented yet","CODE")
		}
		if msg.Error != nil {
			log.Infof("Received message of type '%s', but handling is not implemented yet","ERROR")
		}
		if msg.Warning != nil {
			log.Infof("Received message of type '%s', but handling is not implemented yet","WARNING")
		}
	}
}

func (c *Client) handleNodesMessage(nodes []model.Node){
	for _,node := range nodes {
		c.handleNodeMessage(node)
	}
}

func (c *Client) handleNodeMessage(node model.Node){
	//if _,ok := c.Nodes[node.Id]; !ok {
	//	c.addNode(node)
	//}
}

func (c *Client) addNode(node model.Node){
	//log.Infof("Adding node for '%s' - %v",c.config.UID,node)
	//c.Nodes[node.Id] = node
	//if thing, err := model.GenerateThingMapping(node, c.config.UID); err == nil {
	//	log.Infof("Generated new thing mapping for node %v: %v",node.Id,thing)
	//	c.Things[node.Id] = *thing
	//
	//} else {
	//	log.Infof("Unable to generate thing mapping for node %v (uuid:%s) (node:%v)",node.Id,c.config.UID,node)
	//}
}


func (c *Client) handleAttributeMessage(Attribute model.Attribute){

	//c.Attributes[Attribute.Id] = Attribute
	//// when node is unknown, request that node
	//if _,ok := c.Nodes[Attribute.NodeId];!ok {
	//	c.requestNode(Attribute.NodeId)
	//	// property change does not need to be handled here
	//	// after requesting the node, all properties will be transferred as well
	//} else {
	//
	//	// TODO: handle property change mapping
	//}

}

func (c *Client) requestNode(nodeId int){
	log.Infof("requesting node information for node %v",nodeId)
	c.connection.sendMessage("get:nodes/"+strconv.Itoa(nodeId))
}
