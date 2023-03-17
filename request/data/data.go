package data

import "fmt"

type Data struct {
	name   string
	number string
}

func NewData(number, text string) *Data {
	return &Data{
		number: number,
		name:   text,
	}
}

func (d *Data) String() string {
	return fmt.Sprintf("name: %s \nnumber: %s", d.name, d.number)
}

func (d *Data) IsNil() bool {
	return d.name == "" && d.number == ""
}
