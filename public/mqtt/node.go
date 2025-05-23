package mqtt

import meshtastic "github.com/kmpm/meshtastic-protobufs.go/v2/generated"

// Node implements a meshtastic node that connects only via MQTT
type Node struct {
	user *meshtastic.User
}

func NewNode(user *meshtastic.User) *Node {
	return &Node{
		user: user,
	}
}
