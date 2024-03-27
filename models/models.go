package main

import "fmt"

//some basic structs

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person // Embedded type (anonymous field)
	Salary float64
}

type IPAddr [4]byte

func main() {

	emp := Employee{
		Person: Person{Name: "John", Age: 30},
		Salary: 50000,
	}

	fmt.Println("Name:", emp.Name)
	fmt.Println("Age:", emp.Age)
	fmt.Println("Salary:", emp.Salary)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	fmt.Println(hosts)
	for name, ip := range hosts {
		fmt.Printf("%s: %s:\n", name, ip)
	}

}

// changed the printing like a toString in JAVA
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}
