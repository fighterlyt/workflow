package entity

import (
	"time"

	"github.com/fighterlyt/gographviz"
	"github.com/fighterlyt/workflow/def"
)

type Node struct {
	id         string
	definition *FlowDefinition
	name       string
	status     Status
	createTime time.Time
}

const (
	InvalideNodeName = "非法状态"
)

func (n Node) GetType() def.EntityType {
	return def.Enode
}
func (n Node) GetContent() map[def.EntityType][]string {
	return nil
}

func (n Node) GetId() string {
	return n.id
}

func (n Node) GetLabel() string {
	if node := n.definition.nodes.Lookup[n.name]; node != nil {
		if node.Attrs != nil {
			if label, ok := node.Attrs[gographviz.Label]; ok {
				return label
			}
		}
		return n.name
	} else {
		return InvalideNodeName
	}

}
func (n Node) GetName() string {
	return n.name
}

func (n Node) GetStatus() Status {
	return n.status
}

func newNode(def *FlowDefinition, name string, status Status) *Node {
	return &Node{
		id:         "",
		definition: def,
		name:       name,
		status:     status,
	}
}
