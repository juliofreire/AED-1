package onsort

// import "fmt"

func MergeSort(v []int) []int {

	size := len(v)

	if size > 1 {
		mid := size / 2

		v1 := make([]int, mid)
		v2 := make([]int, size-mid)

		for i := 0; i < mid; i++ {
			v1[i] = v[i]
		}
		for i := mid; i < size; i++ {
			v2[i-mid] = v[i]
		}

		MergeSort(v1)
		MergeSort(v2)
		merge(v, v1, v2)
	}

	return v

}

func merge(v []int, vE []int, vD []int) []int {
	sizeE := len(vE)
	sizeD := len(vD)

	indexE := 0
	indexD := 0
	indexV := 0

	for indexE < sizeE && indexD < sizeD {

		if vE[indexE] <= vD[indexD] {
			v[indexV] = vE[indexE]
			indexE++

		} else {
			v[indexV] = vD[indexD]
			indexD++
		}
		indexV++
	}
	for indexE < sizeE {
		v[indexV] = vE[indexE]
		indexE++
		indexV++

	}
	for indexD < sizeD {
		v[indexV] = vD[indexD]
		indexD++
		indexV++

	}

	return v
}
