package main

import (
	"encoding/json"
	"fmt"
)

type Human interface {
	PrintData(Data []int)
}

type Person struct {
	Name  string
	Age   int
	Email string
}

func (p Person) PrintData(Data []int) {
	for _, value := range Data {
		fmt.Println("Slice Data :- ", value)
	}
	fmt.Println("this is a person..")
}

func main() {
	p := Person{Name: "shiv", Age: 29, Email: "ershivkumarrana@gmail.com"}
	slice := []int{2, 1, 3, 4}
	// var h Human = p
	byteData, err := json.Marshal(p)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("json Marshal message", string(byteData))
	per := &Person{}
	err = json.Unmarshal(byteData, per)
	if err != nil {
		fmt.Println(" UnMarshal Message :-", err)
	}
	fmt.Println(*per)
	fmt.Println(per.Name, per.Age, per.Email)
	fmt.Printf("Unmarshalled Person: %+v\n", *per)

	p.PrintData(slice)

	fmt.Println("second interface...")
}

//-------------------------------
// package main

// import "fmt"

// type USB interface {
// 	MouseAction()
// 	Connection()
// }

// type Zebronic struct{}

// func (z Zebronic) MouseAction() {
// 	fmt.Println("zebronic starts mouseAction")
// }

// func (z Zebronic) Connection() {
// 	fmt.Println("zebronic start connecction..")
// }

// type Iball struct{}

// func (i Iball) MouseAction() {
// 	fmt.Println("this is starting mouseAction")
// }

// func (i Iball) Connection() {
// 	fmt.Println("this is starting connecction")
// }

// func main() {
// 	var ball Iball
// 	ball.Connection()
// 	ball.MouseAction()
// 	var z Zebronic
// 	z.Connection()
// 	z.MouseAction()
// 	// var usb USB
// 	// ball.usb()
// 	// usb =z
// 	fmt.Println("start interface")
// }
