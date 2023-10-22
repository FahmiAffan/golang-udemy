package main

import "fmt"

func say(fName string, lName string) string {
	return "hello" + fName
}

func main() {
	name := [...]string{"Jendrod", "Bocil", "lano", "Nai", "Fahmi"}
	slice := name[0:4]

	person := map[string]string{
		"fname": "fahmi",
		"lname": "affan",
	}

	fmt.Println(say("fahmi", "affan"))

	var book map[string]string = make(map[string]string)
	book["title"] = "belajar"
	book["author"] = "Fahmi"
	// a := 20
	// max := 299 + a + b
	// var eString string = string(e)
	fmt.Println("ini", 1)
	fmt.Println("ini", 2)
	fmt.Println("ini", 3)
	// fmt.Println(max)
	fmt.Println(slice)

	fmt.Println(book)
	// delete(person, "lname")
	fmt.Println(person)

	if person["fname"] == "fahmi" && person["lname"] == "affn" {
		fmt.Println("hello " + person["fname"])
	} else {
		fmt.Println(len(person["fname"]))
	}

	switch person["lname"] {
	case "affan":
		fmt.Println("ini benar " + person["lname"])
	case "fahmi":
		fmt.Println("ini tidak " + person["lname"])
	default:
		fmt.Println("kenalan boleh tuh")
	}

	// fmt.Println(e, eString)

}
