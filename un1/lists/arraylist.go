package lists

import (
	"errors")

type ArrayList struct{
	values []int
	size int
}


func (arraylist *ArrayList) Init(size int) error{
	// if the size is acceptable make a slice with cap of the size
	if size > 0 {
		arraylist.values = make([]int, size)
		return nil
	} else if size == 0{
		return errors.New("Can't build a Array List with no size")
	} else{
		return errors.New("Can't build a Array List negative size")
	}
}

func (arraylist *ArrayList) double() {
	// creating a double sized array

	doublesizedarray := make([]int, len(arraylist.values)*2)

	for i, v := range(arraylist.values){
		doublesizedarray[i] = v
	}
	arraylist.values = doublesizedarray
}

func (arraylist *ArrayList) Add(value int) {
	if (arraylist.size) == len(arraylist.values){
		arraylist.double()
	}
	arraylist.values[arraylist.size] = value
	arraylist.size++
}

func (arraylist *ArrayList) Set(value int, index int) error{
	if index >= 0 && index <= arraylist.size{
		arraylist.values[index] = value
		return nil
	} else if index < 0{
		return errors.New("Can't set a value of a negative index of a Array List")
	} else {
		return errors.New("Can't set a value of a index that exceeds the size of a Array List")
	}
}

func (arraylist *ArrayList) AddOnIndex(value int, index int) error{
	
	if index >= 0 && index <= arraylist.size{
		if arraylist.size < len(arraylist.values) {
			for i := arraylist.size; i > index; i-- {
				arraylist.values[i] = arraylist.values[i-1]				
			}
			arraylist.Set(value, index)

		} else if arraylist.size == len(arraylist.values){
			arraylist.double()
			for i := arraylist.size; i > index; i-- {
				arraylist.values[i] = arraylist.values[i-1]				
			}
			arraylist.Set(value, index)
			
		}
		arraylist.size++
		return nil
	} else if index < 0{
		return errors.New("It's impossible add a value on a negative index")
	} else {
		return errors.New("It's impossible add a value on a index that exceeds the size of array list")
	}

}

func (arraylist *ArrayList) RemoveOnIndex(index int) error{
	if index > 0 && index <= arraylist.size{
		for i := arraylist.size; i < index; i--{
			arraylist.values[i] = arraylist.values[i-1]
		}
		arraylist.size--
	}
	return nil
}

func (arrayList *ArrayList) Get(index int) (int,error){
	if index >= 0 && index < arrayList.size {
		return arrayList.values[index], nil
	} else if index < 0 {
		return index, errors.New("It's impossible to get a value of a negative index")
	} else {
		return index, errors.New("The index exceeds the size of the ArrayList")
	}

}

func (arraylist *ArrayList) Size() (int, int) {
	return arraylist.size, len(arraylist.values)
}