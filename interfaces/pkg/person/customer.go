package person

type Customer struct {
	Id   uint
	Name string
	Age  uint
}

func (c *Customer) GetAge() uint {
	return c.Age
}
