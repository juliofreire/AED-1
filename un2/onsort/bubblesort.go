package onsort

func BubbleSort(v []int) []int{

	tam := len(v)

	for i := 0; i < tam ; i++{
		for j := 0; j < tam - i - 1; j++{
			if v[j] > v[j+1]{
				v[j], v[j+1] = v[j+1], v[j]
			}

		}
	}

	return v
}