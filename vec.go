package main

import (
	"math"
)

// 坐标向量
type vec struct {
	// x 坐标
	x int

	// y 坐标
	y int
}

// 计算两点之间的距离
func (this *vec) distance(other *vec) int {
	return int(math.Abs(float64(this.x)-float64(other.x)) + math.Abs(float64(this.y)-float64(other.y)))
}

// 判断两点是否相等
func (this *vec) equal(other *vec) bool {
	return this.x == other.x && this.y == other.y
}
