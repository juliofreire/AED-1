package lists

type IList interface{
	Add(value int) 
	AddOnIndex(value int, index int) error
	RemoveOnIndex(index int) error
	Get(indext int) error
	Set(value int, index int) error
	Size() int

}