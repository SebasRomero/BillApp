package invoiceitem

//Item contains information about invoiceitem
type Item struct {
	id      uint
	product string
	value   float64
}

//New returns a new item
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

func (i Item) Value() float64 {
	return i.value
}
