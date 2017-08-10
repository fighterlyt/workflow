package entity

import "github.com/fighterlyt/workflow"

type Entity interface{
	GetType() workflow.EntityType
}


