package onsort

func CountingSort(v []int) []int {

	// first step: get the min and max value
	min := v[0]
	max := v[0]
	for _, val := range v {
		if val <= min {
			min = val
		}
		if val >= max {
			max = val
		}
	}

	mm_value := (max - min) + 1

	// fmt.Println(min)
	// fmt.Println(max)

	// second step: create a vector c that will count the presence of each number
	// where the first value represents the min value
	c := make([]int, mm_value)

	// to map V in C, just subtract the V[i] of min value # V[i]-min

	// third step: count the numbers in v

	for _, val := range v {
		c[val-min] += 1
	}

	// fmt.Println(c)

	// fourth step: do a cumulative sum in c

	for i := 1; i <= len(c)-1; i++ {
		c[i] += c[i-1]
	}

	// fmt.Println(c)

	// fifth step: create a vector o with same lenght of v
	o := make([]int, len(v))

	for _, val := range v {
		o[c[val-min]-1] = val
	}

	return o
}
