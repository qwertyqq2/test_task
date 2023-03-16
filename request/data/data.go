package data

import "fmt"

type Data struct {
	price string
	text  string
}

func NewData(price, text string) *Data {
	return &Data{
		price: price,
		text:  text,
	}
}

func (d *Data) String() string {
	return fmt.Sprintf("price: %s, \ninfo: %s", d.price, d.text)
}

func (d *Data) IsNil() bool {
	return d.price == "" && d.text == ""
}
