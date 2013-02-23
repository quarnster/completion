package parse

type nodeType int

type node struct {
	parent   *node
	children []*node
	pos      int
	endPos   int
	name     string
	nodeType nodeType
}
