package invoice

import (
	"github.com/sebasromero/composition/pkg/customer"
	"github.com/sebasromero/composition/pkg/invoiceitem"
)

//Invoice is the invoice's struct
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   []invoiceitem.Item
}

//New returns a new invoice
func New(country, city string, client customer.Customer, items []invoiceitem.Item) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

//SetTotal is the Invoice.total setter
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}
