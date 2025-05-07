package main

import "fmt"

func unimodalMax(u []int) int {
    if (len(u) == 1) {
        return u[0]
        }
    
    mitte := len(u)/2
    fmt.Println(mitte)
    if (u[mitte] < u[mitte-1]) {
	    unimodalMax(u[:mitte])
    } else {
	    unimodalMax(u[:mitte+1])
    }
    return u[mitte]
    }

func main() {
    u := []int {2, 4, 5, 6, 7, 8, 4, 3, 1}
    result := unimodalMax(u)
    fmt.Println(result)
    }
