package main

// 节点状态枚举
type nodeEnum int

const (
	// 不存在
	nodeEnum_NoExist nodeEnum = iota

	// 在openList
	nodeEnum_Open

	// 在closeList
	nodeEnum_Close
)

type node struct {
	// 到起点的步数
	g int

	// 到终点的步数
	h int

	// 当前节点坐标
	pos *vec

	// 当前节点状态
	status nodeEnum

	// 父节点
	parent *node
}

// 计算节点f值
func (this *node) f() int {
	return this.g + this.h
}

// 构造函数
func newNode(pos vec) *node {
	return &node{
		g:      0,
		h:      0,
		pos:    &pos,
		status: nodeEnum_NoExist,
		parent: nil,
	}
}
