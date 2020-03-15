package bubble

func Sort(data *[]int) {
	for i, l := 0, len(*data); i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if (*data)[j] > (*data)[j+1] {
				(*data)[j], (*data)[j+1] = (*data)[j+1], (*data)[j]
			}
		}
	}
}

// 针对于冒泡排序的优化算法，当一趟冒泡没有发生交换时，意味着排序已经完成，可以结束掉整个排序了
// 这个优化对于数据分布均匀的数集没什么用，反而会增加n次的运算
func MtSort(data *[]int) {
	for i, l := 0, len(*data); i < l-1; i++ {
		flag := false
		for j := 0; j < l-i-1; j++ {
			if (*data)[j] > (*data)[j+1] {
				(*data)[j], (*data)[j+1] = (*data)[j+1], (*data)[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

