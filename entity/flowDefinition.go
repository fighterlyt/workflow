package entity

import (
	"io"
	"time"

	"io/ioutil"

	"fmt"
	"strconv"

	"github.com/fighterlyt/gographviz"
	"github.com/pkg/errors"
)

const (
	attrVersion       = "attr"
	attrCreateTime    = "createTime"
	attrCanWithDraw   = "canWithDraw"
	attrKind          = "kind"
	attrCanMultiSend  = "canMultiSend"
	attrSingleForward = "singleForward"
	attrDir           = "attrDir"
)

var (
	f FlowDefinitionGenerator
)

//FlowDefinition 流程定义
type FlowDefinition struct {
	name          string
	version       int
	createTime    time.Time
	canWithDraw   bool
	canMultiSend  bool
	singleForward bool
	edges         *gographviz.Edges
	nodes         *gographviz.Nodes
}

func (f FlowDefinitionGenerator) NewFlowDefinition(reader io.Reader) (*FlowDefinition, error) {
	if data, err := ioutil.ReadAll(reader); err != nil {
		return nil, errors.Wrap(err, "读取流程定义出错")
	} else {
		if graph, err := gographviz.Read(data); err != nil {
			return nil, errors.Wrap(err, "解析流程定义出错")
		} else {
			flow := f.NewFloWDefinition()
			flow.name, flow.edges, flow.nodes = graph.Name, graph.Edges, graph.Nodes

			if err = flow.GetVersion(graph); err != nil {
				return nil, errors.Wrap(err, "生成流程定义出错")
			}
			return flow, nil
		}
	}
}
func (f *FlowDefinition) GetVersion(graph *gographviz.Graph) error {
	var err error
	if graph.Attrs != nil {
		if v, ok := graph.Attrs[attrVersion]; ok {
			if f.version, err = strconv.Atoi(v); err != nil {
				return fmt.Errorf("流程版本错误:%s", err.Error())
			}

		}
	}
	return nil
}
func (f *FlowDefinition) NewEdge(from, to *Node, dir Dir) (*Edge, error) {
	if from == nil || to == nil {
		return nil, errors.New("边对应的两个节点不能为空")
	}

	if edge := f.GetEdge(from.GetName(), to.GetName(), dir); edge == nil {
		return nil, errors.New("边未定义")
	} else {
		return NewEdge(from.GetId(),to.GetId(),getEdgeLabel(f,),dir)
	}

}
func (f *FlowDefinition) GetNodeByName(name string) *gographviz.Node {
	return f.nodes.Lookup[name]
}
func (f *FlowDefinition) GetEdge(from, to string, dir Dir) *gographviz.Edge {
	if fromEdges := f.edges.SrcToDsts[from]; fromEdges != nil {
		if endEdges := fromEdges[to]; len(endEdges) != 0 {
			for _, edge := range endEdges {
				if diraAttr, exist := edge.Attrs[attrDir]; exist && Dir(diraAttr) == dir {
					return edge
				}
			}

		}
	}
	return nil
}

//FlowDefinitionGenerator 一个流程定义构造器,定义了所有扩展属性的默认值
type FlowDefinitionGenerator struct {
}

/*      下方是扩展属性的默认值     */
func (f FlowDefinitionGenerator) getDefaultVersion() int {
	return 1
}

func (f FlowDefinitionGenerator) getDefaultCreateTime() time.Time {
	return time.Now()
}

func (f FlowDefinitionGenerator) getDefaultCanWithDraw() bool {
	return true
}

func (f FlowDefinitionGenerator) getDefaultKind() Kind {
	return Kprocessing
}

func (f FlowDefinitionGenerator) getDefaultCanMultiSend() bool {
	return false
}
func (f FlowDefinitionGenerator) getDefaultSingleForward() bool {
	return true
}

/*      扩展属性的默认值方法结束    */

func (f FlowDefinitionGenerator) NewFloWDefinition() *FlowDefinition {
	return &FlowDefinition{
		version:       f.getDefaultVersion(),
		createTime:    f.getDefaultCreateTime(),
		canWithDraw:   f.getDefaultCanWithDraw(),
		canMultiSend:  f.getDefaultCanMultiSend(),
		singleForward: f.getDefaultSingleForward(),
	}
}
