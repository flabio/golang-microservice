package domain

type CustomerRepositoryStub struct {
	customer []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customer, nil
}
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "10", Name: "Flabio", City: "Buenaventura", ZipCode: "12fdfd", DateofBirth: "1987-02-12", Status: "1"},
		{Id: "12", Name: "Santiago", City: "Bogota", ZipCode: "1111", DateofBirth: "1987-02-12", Status: "1"},
	}
	return CustomerRepositoryStub{customer: customers}
}
