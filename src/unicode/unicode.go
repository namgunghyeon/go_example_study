package main

import "fmt"

func main() {
    for i, r := range "가나다" {
        fmt.Println(i, r)
    }
    for _, r := range "가나다" {
        fmt.Println(string(r))
    }
}
