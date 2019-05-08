package pregel

type Store interface {
	RegisterDataType(func() interface{})
	Put(nodes ...Node) error
	PutNodeData(string, Data) error
	PutEdges(string, ...*Edge) error
	PutEdgeData(string, string, Data) error
	Get(string) (Node, bool, error)
	Delete(string) error
	DeleteEdge(string, string) error
}
