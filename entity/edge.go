package entity

import "time"

type Dir int

const (
	Forward Dir = iota
	BackWard
	Self
)

type Edge struct {
	srcId      string
	targetId   string
	dir        Dir
	createTime time.Time
	name       string
	definition *FlowDefinition
}

func (e Edge) GetType() EntityType {
	return Eedge
}

func (e Edge) GetContent() map[EntityType][]string {
	return nil
}

func NewEdge(srcId, targetId, name string, dir Dir, createTime time.Time) *Edge {
	return &Edge{
		srcId:      srcId,
		targetId:   targetId,
		dir:        dir,
		name:       name,
		createTime: createTime,
	}
}
