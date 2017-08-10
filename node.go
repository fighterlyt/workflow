package workflow

type Node struct {
	id       string        //节点id
	step     Step          //对应的业务步骤
	flowId   string        //所属的流程id
	business BusinessInfo  //节点关联的业务信息
	edge     transformInfo //转换信息

}

type transformInfo struct {
	fromId   string //起点
	targetId string //终点
}

func newNode(flowId string, step Step) *Node {
	return &Node{
		id:     idGenerator.Generate(),
		step:   step,
		flowId: flowId,
	}
}
