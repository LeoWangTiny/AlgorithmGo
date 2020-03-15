package main

import (
	"algorithm/sort/bubble"
	"algorithm/sort/selected"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	CheckFuncEff("冒泡排序", bubble.Sort, 10000)
	CheckFuncEff("选择排序", selected.Sort, 10000)
	CheckFuncEff("选择优化", selected.MtSort, 10000)
}

func CheckFuncEff(name string, function func(*[]int), length int) {
	a := MakeMtRandArr(length)
	timeA := time.Now().UnixNano()
	function(&a)
	fmt.Println(name, "排序个数：", length, "共耗时:", (time.Now().UnixNano()-timeA)/1000, "μs")
}

func MakeMtRandArr(length int) []int {
	rand.Seed(time.Now().UnixNano())
	ret := make([]int, length)
	for i := 0; i < length; i++ {
		ret[i] = rand.Intn(length * 10)
	}
	return ret
}
