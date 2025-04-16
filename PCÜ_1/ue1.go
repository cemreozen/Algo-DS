package main
import "fmt"
func doSomething(p *int) {
	*p = 10
}

func main() {
	value:=5
	fmt.Println(value)
	doSomething(&value)
	fmt.Println(value)
    }
