package workflow

import "fmt"

type Status int

const (
	Draft      Status = iota //草稿
	Auditing                 //审批中
	Rejected                 //退回
	WithDrawed               //撤回
	TempSaved                //暂存
	Finished                 //完成
	Invalid                  //非法
)

var (
	statusMap = map[Status]string{
		Draft:      "草稿",
		Auditing:   "审批中",
		Rejected:   "退回",
		WithDrawed: "撤回",
		TempSaved:  "暂存",
		Finished:   "完成",
		Invalid:    "非法",
	}
)

func (s Status) String() string {
	if str, ok := statusMap[s]; ok {
		return str
	} else {
		return statusMap[Invalid]
	}
}
func (s Status) Kind() Kind {
	switch s {
	case Draft:
		return Kdraft
	case Auditing, Rejected, WithDrawed, TempSaved:
		return Kprocessing
	case Finished:
		return Kfinish
	default:
		return KInvalid
	}
}

/*
	--------------------------------------------------------------------
	Kind 表示流程步骤大状态
*/
type Kind int

const (
	Kdraft      Kind = iota //草稿
	Kprocessing             //处理中
	Kfinish                 //完成
	KInvalid                //非法
)

var (
	kindMap = map[Kind]string{
		Kdraft:      "草稿",
		Kprocessing: "处理中",
		Kfinish:     "完成",
		KInvalid:    "非法状态",
	}
)

func (k Kind) String() string {
	if str, ok := kindMap[k]; ok {
		return str
	} else {
		return kindMap[KInvalid]
	}
}

type Step struct {
	name   string
	status Status
}

func (s Step) Full() string {
	return fmt.Sprintf("%s-%s %s", s.name, s.Kind(), s.status)
}

func (s Step) Kind() Kind {
	return s.status.Kind()
}

func (s Step) Status() Status {
	return s.status
}

func (s Step) String() string {
	return s.name
}

type Steps struct {
	data []Step
}
