package main

import (
	"fmt"
)

func main() {
	// 地图映射数组
	maps := [10][10]int{
		{0, 1, 0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 1, 0, 0},
		{1, 1, 1, 1, 0, 1, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
		{0, 1, 0, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 1, 0},
	}

	// 初始化
	width := 10
	height := 10
	canPassFunc := func(pos vec) bool {
		return maps[pos.y][pos.x] == 0
	}
	obj := newaStar(width, height, canPassFunc)

	// 设置起点，终点,开始寻路
	end := newNode(vec{9, 9})
	start := newNode(vec{0, 0})
	path := obj.find(start, end)

	// 打印路径图形
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			found := false
			for index := 0; index < len(path); index++ {
				if path[index].pos.x == j && path[index].pos.y == i {
					found = true
					break
				}
			}

			if found {
				print("1 ")

			} else {
				print("* ")
			}
		}
		print("\n")
	}

	// 打印一共花费步数
	fmt.Printf("\npath step is %v\n", len(path))
}
