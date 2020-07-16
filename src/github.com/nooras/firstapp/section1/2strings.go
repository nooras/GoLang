package main

//All strings in go are mutabel that means you can change it any time
//rune
import ( //Multiple import
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//string - double quotes  single word - single quotes
	var m1 string
	m1 = "Nooras,Fatima"
	fmt.Println(m1)

	m2 := "Fatima" //Direct Initialization
	fmt.Println(m1 + " " + m2)

	fmt.Println(strings.Contains(m1, m2)) //checks m2 string is in m1 or not and return boolian
	fmt.Println(strings.ReplaceAll(m1, "oo", "u"))
	fmt.Println(strings.Split(m1, ","))

	var i int = 42
	fmt.Printf("%v,%T\n", i, i)
	var j string
	j = string(i) //Prints *
	fmt.Printf("%v,%T\n", j, j)
	j = strconv.Itoa(i) //Integers to asterisk
	fmt.Printf("%v,%T\n", j, j)

	s := "Hello"
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", string(s[2]), s[2])
	//s[2] = "h" //cannot update
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	r := 'a'
	fmt.Printf("%v, %T\n", r, r)
	var r2 rune = 'a'
	fmt.Printf("%v, %T\n", r2, r2)
}
