package def

type Entity interface {
	GetType() EntityType
	GetContent() map[EntityType][]string
}

type EntityType int

const (
	Eflow EntityType = iota
	Enode
	Eedge
)
