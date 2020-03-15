package selected

func Sort(data *[]int) {
	for i, l := 0, len(*data); i < l-1; i++ {
		flag := i
		for j := i + 1; j < l; j++ {
			if (*data)[j] < (*data)[flag] {
				flag = j
			}
		}
		(*data)[flag], (*data)[i] = (*data)[i], (*data)[flag]
	}
}

func MtSort(data *[]int) {
	for i, l := 0, len(*data); i < l; i++ {
		min, max := i, i
		for j := i; j < l; j++ {
			if (*data)[j] < (*data)[min] {
				min = j
			}
			if (*data)[j] > (*data)[max] {
				max = j
			}
		}
		if i != min {
			(*data)[i], (*data)[min] = (*data)[min], (*data)[i]
		}
		if i != max {
			(*data)[l-1], (*data)[max] = (*data)[max], (*data)[l-1]
		}
		l--
	}
}
