package main
import (
    "fmt"
)
type fare interface {
    fare_calculation() string
}
type car struct {
    carname string
}
type bus struct {
    busname string
}
func new_car() *car {
    return &car{carname: "skoda"}
}
func new_bus() *bus {
    return &bus{busname: "tata"}
}
func (c car) fare_calculation() string {
    return "3000 per day" + "-carname->" + c.carname
}
func (b bus) fare_calculation() string {
    return "6000 per day" + "-busname->" + b.busname
}
func Factory(no_of_person int) fare {
    var a fare
    if no_of_person > 6 {
        a = new_bus()
    } else {
        a = new_car()
    }
    return a
}
func main() {
    f := Factory(4)
    fmt.Println(f.fare_calculation())
}