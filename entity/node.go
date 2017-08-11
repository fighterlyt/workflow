package entity

import "github.com/fighterlyt/gographviz"

type Node struct {
	id         string
	definition *FlowDefinition
	name       string
}

const (
	InvalideNodeName = "非法状态"
)

func (n Node) GetType() EntityType {
	return Enode
}
func (n Node) GetContent() map[EntityType][]string {
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
