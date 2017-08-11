package db

import "github.com/fighterlyt/workflow/entity"

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
	Load(id string, data entity.Entity) error
	Save(data entity.Entity) error
	Update(id string, data entity.Entity) error
	Delete(data entity.Entity) error
}
