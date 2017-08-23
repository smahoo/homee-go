package homee

import "github.com/connctd/homee/model"

type HomeeEvent struct {
	EventType int
	Client *Client

}

type HomeeNodeEvent struct {
	HomeeEvent
	Node *model.Node
}

type HomeeAttributeEvent struct {
	HomeeEvent
	Attribute *model.Attribute
}


const (
	EVENT_TYPE_NODE_ADDED			= 1001

	EVENT_TYPE_ATTRIBUTE_VALUE_CHANGE 	= 2002

	EVENT_TYPE_ERROR			= 3000

	EVENT_TYPE_WARNING			= 4000
)