package entity

import (
	"github.com/fighterlyt/gographviz"
	"github.com/fighterlyt/workflow/db"
)

var (
	dbPlugin db.DBPlugin
)

func init() {
	gographviz.NewAttr(attrVersion)
}
