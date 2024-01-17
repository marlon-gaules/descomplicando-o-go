package main

import "fmt"

func dummyFunction(s string) string {
	return fmt.Sprintf("Ran dummy function. %s\n", s)
}

func DummyPublicFunction(s string) string {
	return dummyFunction(s)
}

func main() {
	s := dummyFunction("Hello World")
	fmt.Println(s)

	var u string
	u = dummyFunction("Hello World again")
	fmt.Println(u)

	var input int
	fmt.Scan(&input)
	if input == 1 {
		fmt.Println("É 1")
	} else if input == 2 {
		// do something
	} else {
		fmt.Println("Não é 1")
	}

	switch input {
	case 1:
		fmt.Println("É 1")
	case 2:
		fmt.Println("É 2")
	case 3:
		fmt.Println("Não é 1")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	running := true
	for running {
		// do something
		if input == 51 {
			running = false
		}
	}
	fmt.Println(input)
}
