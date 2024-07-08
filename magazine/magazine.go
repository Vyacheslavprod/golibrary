package magazine

type Subscriber struct {
	Name string
	Rate float64
	Active bool
	HomeAddress Address
}

type Employee struct {
	Name string
	Salary float64
	Address
}

type Address struct {
	Street string
	City string
	State string
	PostalCode string
}