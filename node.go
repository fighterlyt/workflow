package workflow

type Node struct {
	id         string        //节点id
	step       Step          //对应的业务步骤
	flowId     string        //所属的流程id
	business   BusinessInfo  //节点关联的业务信息
	fromNodeId string        //上一个节点Id
	target     transformInfo //转换信息

}

type transformInfo struct {
}

func newNode(flowId string,step Step) *Node{
	return &Node{
		id:idGenerator.Generate(),
		step:step,
		flowId:flowId,
	}
}
