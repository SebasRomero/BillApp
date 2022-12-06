package customer

//Customer is the client's struct
type Customer struct {
	name    string
	address string
	phone   string
}

//New returns a new customer
func New(name, address, phone string) Customer {
	return Customer{name, address, phone}
}
