package main

import "fmt"

func Fib(n int) int {
    p, q := 0, 1
    for i := 0; i < n; i++ {
        fmt.Println(p, p+q)
        p, q = q, p+q

    }
    return p
}
func main() {
    fmt.Println(Fib(5))
}
