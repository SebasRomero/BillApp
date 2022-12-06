package main

import (
	"fmt"

	"github.com/sebasromero/composition/pkg/customer"
	"github.com/sebasromero/composition/pkg/invoice"
	"github.com/sebasromero/composition/pkg/invoiceitem"
)

func main() {
	invoice := invoice.New("Colombia",
		"Barranquilla",
		customer.New("Sebastian", "9021 sw", "3059922022"),
		[]invoiceitem.Item{
			invoiceitem.New(1, "Go course", 30),
			invoiceitem.New(2, "Js course", 20),
			invoiceitem.New(3, "Python course", 20),
		},
	)
	invoice.SetTotal()
	fmt.Printf("\n%+v", invoice)
}
