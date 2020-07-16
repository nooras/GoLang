package main

import "fmt"

//Arrays : list of ordered DT items

func todo() {

	//var arr []int //Definaing an array
	arr := []int{1, 2, 3, 4, 5, 6} // Initializing an array
	//arr := [...]int{1,2,3}
	arr2 := []string{"Hii,", "I am"}
	arr2 = append(arr2, "Nooras")
	arr2 = append(arr2, "Fatima", "Ansari") //Can add multiple string at a time
	fmt.Println(arr)
	fmt.Println(arr2)

	var students [3]string
	students[0] = "nooras"
	students[1] = "fatima"
	students[2] = "ansari"
	fmt.Printf("Studets: %v\n", students)

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{1, 0, 0}, [3]int{1, 0, 0}}
	fmt.Println(identityMatrix)

	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 2, 3}
	identityMatrix2[1] = [3]int{1, 2, 3}
	identityMatrix2[2] = [3]int{1, 2, 3}
	fmt.Println(identityMatrix2)

	// a := [...]int{1, 2, 3}
	// b := &a
	// b[0] = 5
	// fmt.Println(a)
	// fmt.Println(b)

	// a:=[3]int{1,2,3}
	// a:=[...]int{1,2,3}
	// var a [3]int

	//Slice
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   //All elements
	c := a[3:]  //4th element to end
	d := a[:6]  //first 6 element
	e := a[3:6] //4th, 5th and 6th element
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	//Length and capacity
	//x := make([]int, 3) //type, length==capacity
	x := make([]int, 3, 100) //Type , length , Capacity
	fmt.Println(x)
	fmt.Printf("%v\n", len(x))
	fmt.Printf("%v\n", cap(x))

	y := []int{}
	y = append(y, 1)
	fmt.Printf("%v\n", len(y))
	fmt.Printf("%v\n", cap(y))   //On onilne it doubles the capacity everytime
	y = append(y, 2, 3, 4, 5, 6) // append(y,[]int{2,3,4,5,6}...)
	fmt.Printf("%v\n", len(y))
	fmt.Printf("%v\n", cap(y))

	fmt.Printf("%v\n", y)
	z := append(y[:2], y[3:]...)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", y)
}

func main() {
	todo() //Calling a function
}
