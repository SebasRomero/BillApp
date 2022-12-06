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

type Items []Item

func NewItems(items ...Item) Items {
	var it Items
	for _, item := range items {
		it = append(it, item)
	}
	return it
}

func (it Items) Total() float64 {
	var total float64
	for _, item := range it {
		total += item.value
	}
	return total
}
