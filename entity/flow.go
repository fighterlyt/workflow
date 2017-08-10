package entity

import "github.com/fighterlyt/workflow"

type FlowInstance struct {
	nodeIds []string
	edgeIds []string
}

func (f FlowInstance) GetType() workflow.EntityType {
	return workflow.Eflow
}

func (f FlowInstance) GetContent() []Entity {
	result:=make([]Entity,0,len(f.nodeIds)+len(f.edgeIds))

	result=append(result,f.nodeIds...)
	result
	for _,nodeId:=range f.nodeIds{

	}
}
