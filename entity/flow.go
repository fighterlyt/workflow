package entity

import (
	"time"

	"github.com/fighterlyt/workflow/def"
	"github.com/fighterlyt/workflow/id"
	"github.com/pkg/errors"
)

const (
	systemDealer = "_system" //表示由系统内建的
	preparedId   = "preparedId"
)

type Action int

const (
	Send Action = iota
	Save
	WithDraw
	sendBack
	Finish
)

type Flow struct {
	id            string //流程实例id
	definition    *FlowDefinition
	nodeIds       []string
	nodeLoaded    map[string]bool
	edgeIds       []string
	edgetLoaded   map[string]bool
	createTime    time.Time
	currentNode   *Node
	currentDealer string
}

func (f Flow) GetType() def.EntityType {
	return def.Eflow
}

func (f Flow) GetContent() map[def.EntityType][]string {
	return map[def.EntityType][]string{
		def.Enode: f.nodeIds,
		def.Eedge: f.edgeIds,
	}
}

func newFlow(d *FlowDefinition) (*Flow, error) {
	if def == nil {
		return nil, errors.Errorf("不能使用空的流程定义生成流程")
	}
	f := &Flow{
		id:            id.Generate(def.Eflow),
		definition:    def,
		nodeIds:       make([]string, 0, len(d.nodes.Nodes)),
		nodeLoaded:    make(map[string]bool, len(d.nodes.Nodes)),
		edgeIds:       make([]string, 0, len(d.edges.Edges)),
		edgetLoaded:   make(map[string]bool, len(d.edges.Edges)),
		createTime:    time.Now(),
		currentDealer: systemDealer,
		currentNode:   preparedId,
	}
	return f, nil
}

/*
	开始流程,发起人保存或者发送

*/
func (f *Flow) Start(dealer string, targetId string, action Action, d *FlowDefinition, targetNodeName string) error {
	switch action {
	case Send:
		return f.Send()
	case Save:
		return f.Save()
	}
	n := newNode(f.definition, f.definition.start.Name, Draft)
	f.currentNode = n
	f.currentDealer = dealer
	f.nodeLoaded[n.GetId()] = true
	f.nodeIds = append(f.nodeIds, n.GetId())
	return nil
}

/*
	审批人保存
	节点: 添加一个处理节点
	边:  不需要添加边
	当前处理人:  不变
*/
func (f *Flow) Save(dealer string, d *FlowDefinition) error {
	n := newNode(d, f.currentNode.GetName(), f.currentNode.GetStatus())
	f.addNode(n)
	dbPlugin.Save(n)

}


func(f *Flow) addNode(node *Node){
	f.nodeLoaded[node.GetId()] = true
	f.nodeIds = append(f.nodeIds, node.GetId())
}
