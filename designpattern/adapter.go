package main

import "fmt"

type old interface {
	print(string) string
}

type oldstruct struct{}

func (o *oldstruct) print(s string) string {
	fmt.Println("old")
	return s
}

type new interface {
	printnew() string
}

type newstruct struct {
	Oldone old
	text   string
}

func (n *newstruct) printnew() (msg string) {
	if n.Oldone != nil {
		msg = n.Oldone.print(n.text)
	} else {
		fmt.Println("new")
		msg = n.text
	}
	return msg
}

func main() {
	n := &newstruct{
		Oldone: nil,
		text:   "done",
	}
	n.printnew()

	o := &newstruct{
		Oldone: &oldstruct{},
		text:   "donee",
	}
	o.printnew()

}
