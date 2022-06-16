package main

import "fmt"

type Realinterface interface {
	Realdetails()
}

type Realobject struct {
	details string
}

// func (r *Realobject) Realdetails() {
// 	fmt.Println("Details are", r.details)
// }

type Accessor struct {
	Role string
}

type Proxy_obj struct {
	Realobject
	Accessor
}

func (p *Proxy_obj) Realdetails() {
	if p.Role != "Admin" {
		fmt.Println("Not allowed to read")
		return
	}
	fmt.Println(p.details)
}

func main() {
	Ro := Realobject{
		details: "Secret data",
	}
	Ac := Accessor{
		Role: "Admin",
	}
	Nac := Accessor{
		Role: "User",
	}

	Pao := &Proxy_obj{ //For admin
		Realobject: Ro,
		Accessor:   Ac,
	}

	Pno := &Proxy_obj{ //For non admin
		Realobject: Ro,
		Accessor:   Nac,
	}

	Pao.Realdetails()
	Pno.Realdetails()

}

//https://www.smartscribs.com/proxy-design-pattern-in-golang/

//In proxy design pattern, a resource (class/struct) acts or represents functionality
//similar to another resource on behalf of that resource
