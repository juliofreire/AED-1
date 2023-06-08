package onsort

func SelectionSort (v []int, method string) []int{
	// method can be 'outofplace' or 'inplace'
	
	 var tam int = len(v)

	if method == "outofplace"{
		
		w := make([]int, tam)
	
		idx := 0
		for j:= 0; j < tam; j++{
			
			menor_da_rodada := 999
			
			for i := 0; i < tam; i++{
	
				if v[i] < menor_da_rodada {
					menor_da_rodada = v[i]
					idx = i
				}				
			}
			w[j] = menor_da_rodada
			v[idx] = 999		
		}

		v = w
		return v
	} else if method == "inplace" {

		menor_da_rodada := v[0]

		for i := 0; i<tam; i++{
			for j := i; j<tam; j++{
				if v[j] < menor_da_rodada{
					v[i], v[j] = v[j], v[i]
				}
			}
		}

		return v
	} else {
		return v
	}

}