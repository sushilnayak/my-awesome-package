package awesome

import "fmt"

func Add(a, b int) int {
	return a + b
}
func helper() {
	fmt.Println("Im an internal method, cant be called by module users")
}

type MyStruct struct {
	Name string // public
	age  int    // private
}
