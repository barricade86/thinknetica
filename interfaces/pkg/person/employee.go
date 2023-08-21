package person

type Employee struct {
	Id   uint
	Rank string
	Name string
	Age  uint
}

func (e *Employee) GetAge() uint {
	return e.Age
}
