package main

import (
	"sort"
)

type aStar struct {
	// 地图宽
	width int

	// 地图高
	height int

	// 直走一步需要的消耗
	straightVal int

	// 倾斜走一步需要的消耗
	obliqueVal int

	// 地图映射
	mapping []*node

	// openList
	openlist []*node

	// node 能否走的判断函数
	canPassFunc func(vec) bool
}

// 计算当前节点的g值
func (this *aStar) gValue(current, parent *node) int {
	var gValue int
	if current.pos.distance(parent.pos) == 2 {
		gValue = this.obliqueVal
	} else {
		gValue = this.straightVal
	}
	return gValue + parent.g
}

// 计算当前节点的h值
func (this *aStar) hValue(current, end *node) int {
	hValue := end.pos.distance(current.pos)
	return hValue * this.straightVal
}

// 节点是否存在于开启列表
func (this *aStar) inOpenlist(pos vec) (bool, *node) {
	outNode := this.mapping[pos.y*this.width+pos.x]
	if outNode != nil && outNode.status == nodeEnum_Open {
		return true, outNode
	}
	return false, outNode
}

// 节点是否存在于关闭列表
func (this *aStar) inClosedlist(pos vec) bool {
	node := this.mapping[pos.y*this.width+pos.x]
	if node != nil && node.status == nodeEnum_Close {
		return true
	}
	return false
}

// 判断当前节点是否可行走
func (this *aStar) canWalk(pos vec) bool {
	if pos.x >= 0 && pos.x < this.width && pos.y >= 0 && pos.y < this.height {
		return this.canPassFunc(pos)
	}
	return false
}

// 判断当前点是否可到达目标点
func (this *aStar) canPass(current, destination vec, allowCorner bool) bool {
	if destination.x >= 0 && destination.x < this.width && destination.y >= 0 && destination.y < this.height {
		if this.inClosedlist(destination) {
			return false
		}

		if destination.distance(&current) == 1 {
			return this.canPassFunc(destination)
		} else if allowCorner {
			return this.canPassFunc(destination) && (this.canWalk(vec{current.x + destination.x - current.x, current.y}) && this.canWalk(vec{current.x, current.y + destination.y - current.y}))
		}
	}
	return false
}

// 查找附近可通过的节点
func (this *aStar) findNearbyNodes(current vec, corner bool) []*node {
	// 附近的节点
	nearbyNodes := make([]*node, 0)

	maxRow := current.y + 1
	maxCol := current.x + 1

	rowIndex := current.y - 1
	if rowIndex < 0 {
		rowIndex = 0
	}

	for {
		if rowIndex > maxRow {
			break
		}

		colIndex := current.x - 1
		if colIndex < 0 {
			colIndex = 0
		}

		for {
			if colIndex > maxCol {
				break
			}
			destination := vec{colIndex, rowIndex}
			if this.canPass(current, destination, corner) {
				nearbyNodes = append(nearbyNodes, newNode(destination))
			}
			colIndex++
		}

		rowIndex++
	}
	return nearbyNodes
}

// 如果在openlist中找到这个邻居节点，则重新计算目标点的g值，如果g值小于原始值，重置
func (this *aStar) foundNode(current, destination *node) {
	// 重新继续destination相对于当前点的G值
	newGval := this.gValue(current, destination)
	if newGval < destination.g {
		destination.g = newGval
		destination.parent = current
	}
}

// 如果在openlist中没有找到这个邻居节点,计算邻居节点的g,h值，同时用当前节点设置为父节点,加入openlist
func (this *aStar) notFoundNode(current, destination, end *node) {
	destination.parent = current
	destination.h = this.hValue(destination, end)
	destination.g = this.gValue(current, destination)

	this.mapping[destination.pos.y*this.width+destination.pos.x] = destination
	destination.status = nodeEnum_Open

	this.openlist = append(this.openlist, destination)
}

// 开始自动寻路
func (this *aStar) find(start, end *node) []*node {
	// 返回的路径
	paths := make([]*node, 0)

	// 将起点放入openList
	this.openlist = append(this.openlist, start)

	// 寻路操作
	for {
		if len(this.openlist) == 0 {
			break
		}

		// openlist中的Node按照f值，从小到大排序,（排序可以优化，数据结构使用二叉堆）
		sort.Slice(this.openlist, func(i, j int) bool {
			return this.openlist[i].f() < this.openlist[j].f()
		})

		// 找出f值最小节点
		curNode := this.openlist[0]
		this.openlist = this.openlist[1:]

		// 当前点的状态置为Close
		this.mapping[curNode.pos.y*this.width+curNode.pos.x].status = nodeEnum_Close

		// 是否找到终点
		if curNode.pos.equal(end.pos) {
			for {
				if curNode.parent == nil {
					break
				}
				paths = append(paths, curNode)
				curNode = curNode.parent
			}
			return paths
		}

		// 查找相邻可通过的节点
		nearbyNodes := this.findNearbyNodes(*curNode.pos, false)

		// 遍历相邻节点
		index := 0
		for {
			if index >= len(nearbyNodes) {
				break
			}

			// 如果在openlist中找到这个邻居节点，则重新计算邻居点的g值，如果g值小于原始值，重置
			if exist, nextNode := this.inOpenlist(*nearbyNodes[index].pos); exist {
				this.foundNode(curNode, nextNode)
			} else {
				// 如果在openlist中没有找到这个邻居节点,计算邻居节点的g,h值，同时用当前节点设置为父节点
				this.notFoundNode(curNode, nearbyNodes[index], end)
			}
			index++
		}
	}

	return paths
}

// 构造函数
func newaStar(width, height int, canPassFunc func(vec) bool) *aStar {
	obj := &aStar{
		width:       width,
		height:      height,
		canPassFunc: canPassFunc,
		mapping:     make([]*node, 0),
		openlist:    make([]*node, 0),
		straightVal: 10,
		obliqueVal:  14,
	}

	// 初始化地图
	for i := 0; i < width*height; i++ {
		obj.mapping = append(obj.mapping, new(node))
	}
	return obj
}
