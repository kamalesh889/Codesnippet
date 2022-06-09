package main

import "fmt"

type details struct {
	name string
	roll int
}

type dept_details struct {
	dept string
	hod  string
	details
}

type sub_details struct {
	subject1 string
	subject2 string
	details
}

func (d details) show_details() {
	fmt.Println("Name--", d.name)
	fmt.Println("Roll--", d.roll)
}

func (x dept_details) Showdept_details() {
	fmt.Println("Department details------>")
	fmt.Println("dept--", x.dept)
	fmt.Println("Hod--", x.hod)
	x.show_details()
}

func (y sub_details) Showsub_details() {
	fmt.Println("Subject details-------->")
	fmt.Println("Sub1--", y.subject1)
	fmt.Println("Sub2--", y.subject2)
	y.show_details()
}

//here we have used composition to reuse our code i.e details struct

func main() {
	Student := details{"kamli", 102}
	Dept := dept_details{"It", "Abd", Student}
	Sub := sub_details{"OS", "SE", Student}

	Dept.Showdept_details()
	Sub.Showsub_details()

	//Department := dept_details{dept: "IT", hod: "Abd", Student}

}

// In go instead of inheriting features from a parent class/object, larger objects are composed of the other smaller objects and can thereby use their functionality.
