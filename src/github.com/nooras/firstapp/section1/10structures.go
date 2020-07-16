package main

import "fmt"

//Structure is an abstract DT which consist closely relted data

//structure encapsulation
type Car struct {
	Name    string
	number  int
	ModelNo int
}
type Doctor struct {
	number     int
	actorName  string
	epsiode    []string
	campanions []string
}

//Inherit struct
type anotherCar struct {
	Car
	Speed int
}

//methods in structure
func (d Car) Print() {
	fmt.Println(d)
}

func (d Car) Drive() {
	fmt.Println("Drivinggg")
}

//Return in structure
func (d Car) GetName() string {
	return d.Name
}

func main() {
	c := Car{"Nooras", 20, 14} //Direct declaration
	d := Car{
		Name:    "BMW",
		number:  30,
		ModelNo: 10, //Dont forget to put last comma
	}
	fmt.Println(c)
	d.Print() // Calling Structure method
	d.Drive()
	fmt.Println(d.GetName())

	aDcotor := Doctor{
		number:    3,
		actorName: "Nooras",
		epsiode:   []string{},
		campanions: []string{
			"ABC",
			"XYZ",
		},
	}
	bDcotor := Doctor{
		3,
		"Nooras",
		[]string{},
		[]string{
			"ABC",
			"XYZ",
		},
	}
	fmt.Println(aDcotor)
	fmt.Println(aDcotor.actorName)
	fmt.Println(aDcotor.campanions[1])
	fmt.Println(bDcotor)

	x := struct{ name string }{name: "John"}
	fmt.Println(x)

	//struct inside struct
	b := anotherCar{}
	b.Name = "Bwm"
	b.number = 123
	b.ModelNo = 654
	b.Speed = 213
	fmt.Println(b)
	b = anotherCar{
		Car:   Car{Name: "Nano", number: 62, ModelNo: 32165},
		Speed: 21,
	}
	fmt.Println(b)
}
