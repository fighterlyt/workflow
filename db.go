package workflow

/*
	DBPlugin 定义了整个流程的持久化插件
*/

type EntityType int

const (
	Eflow EntityType = iota
	Enode
	Eddge
)

/*
	DBPlugin 定义了整个流程的持久化插件
	Load    加载
	Save    保存
	Update  更新
	Delete  删除
*/
type DBPlugin interface {
	Load(id string, data Entity) error
	Save(data Entity) error
	Update(id string, data Entity) error
	Delete(data Entity) error
}
