package main

import "fmt"

func main() {
	f := true
	flag := &f
	//In if else it is compulasory to use curly braces and no need to put brackets in condition
	fmt.Println("If Else ------->>>")
	if flag == nil {
		fmt.Println("Flag is nill")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	ss := map[string]int{
		"a": 1,
		"b": 2,
	}
	if pop, ok := ss["b"]; ok {
		fmt.Println(pop)
	}

	fmt.Println("For ------->>>")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	arr := []string{"Nooras", "Fatima"}
	for x, y := range arr { //x is index and y is value
		fmt.Println(x, " ", y)
	}

	fmt.Println("Map ------->>>")
	mymap := make(map[string]interface{})

	mymap["name"] = "Nooras"
	mymap["age"] = 12
	for x, y := range mymap {
		fmt.Printf("Key: %s and Value: %v", x, y)
		fmt.Println()
	}

	fmt.Println("Continue ------->>>")
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("Break ------->>>")
	flagg := true
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			flagg = false
			break
		}
		fmt.Println(i)
	}
	fmt.Println(flagg)

	fmt.Println("Switch ------->>>")
	day := "Fri"
	switch day {
	case "Fri": //case day == Fri:
		fmt.Println("Good Day")
	case "Mon", "Tue", "Wed":
		fmt.Println("Boring")
	default:
		fmt.Println("Default")
	}
}
