package homee

import (
	"net/http"
	log "github.com/golang/glog"
	"strconv"
	"github.com/connctd/homee/model"
)




type ConnectionError struct{
	Url string
	StatusCode int
	StatusMessage string
	Response *http.Response
}

type Client struct {
	Config          HomeeConfiguration
	connection      *apiConnection
	// nodes and attributes
	Attributes      map[int]model.Attribute
	Nodes           map[int]model.Node
	NodeEvents      chan HomeeNodeEvent
	AttributeEvents chan HomeeAttributeEvent
	//AttributeEvents chan
}



func New(config HomeeConfiguration) *Client {
	result := Client{Config:config, connection:nil}
	result.Nodes = make(map[int]model.Node)
	result.Attributes = make(map[int]model.Attribute)
	result.NodeEvents = make(chan HomeeNodeEvent)
	result.AttributeEvents = make (chan HomeeAttributeEvent)
	//result.Things = make(map[int]model.HomeeThingMapping)
	return &result
}

func (c *Client)Connect() {
	log.Infof("establish connection")
	if (c.connection == nil){
		c.connection = newConnection(c.Config, true)
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
	if _,ok := c.Nodes[node.Id]; !ok {
		c.addNode(node)
	}
}

func (c *Client) addNode(node model.Node){
	log.Infof("Adding node for '%s' - %v",c.Config.UID,node)
	c.Nodes[node.Id] = node
	for _,attr := range node.Attributes {
		c.addAttribute(attr)
	}
	event := HomeeNodeEvent{Node:&node}
	event.EventType = EVENT_TYPE_NODE_ADDED
	event.Client = c
	c.NodeEvents <- event
}


func (c *Client) addAttribute(attribute model.Attribute){
	c.Attributes[attribute.Id] = attribute
}


func (c *Client) handleAttributeMessage(Attribute model.Attribute){
	c.Attributes[Attribute.Id] = Attribute
	// when node is unknown, request that node
	if _,ok := c.Nodes[Attribute.NodeId];!ok {
		c.requestNode(Attribute.NodeId)
		// property change does not need to be handled here
		// after requesting the node, all properties will be transferred as well
	} else {
		event :=HomeeAttributeEvent{}
		event.EventType = EVENT_TYPE_ATTRIBUTE_VALUE_CHANGE
		event.Client = c
		event.Attribute = &Attribute
		c.AttributeEvents <- event
	}
}


func (c *Client) requestNode(nodeId int){
	log.Infof("requesting node information for node %v",nodeId)
	c.connection.sendMessage("get:nodes/"+strconv.Itoa(nodeId))
}
