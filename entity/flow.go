package entity

import "time"

type Flow struct {
	nodeIds    []string
	edgeIds    []string
	createTime time.Time
}

func (f Flow) GetType() EntityType {
	return Eflow
}

func (f Flow) GetContent() map[EntityType][]string {
	return map[EntityType][]string{
		Enode: f.nodeIds,
		Eedge: f.edgeIds,
	}
}
