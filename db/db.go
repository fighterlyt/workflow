package db

import "github.com/fighterlyt/workflow/def"

/*
	DBPlugin 定义了整个流程的持久化插件
*/

/*
	DBPlugin 定义了整个流程的持久化插件
	Load    加载
	Save    保存
	Update  更新
	Delete  删除
*/
type DBPlugin interface {
	Load(id string, data def.Entity) error
	Save(data def.Entity) error
	Update(id string, data def.Entity) error
	Delete(data def.Entity) error
}
