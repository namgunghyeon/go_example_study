package main

import "fmt"

func Example_printBytes() {
    s := "가나다"
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x:", s[i])
    }
    fmt.Println()
    // Output:
    // ea:b0:80:eb:82:98:eb:8b:a4:
}

func Example_modifyBytes() {
    b := []byte("가나다")
    b[2]++
    fmt.Println(string(b))
    // Output:
    // 각나다
}

func Example_strCat() {
    s := "abc"
    ps := &s
    s += "def"
    fmt.Println(s)
    fmt.Println(*ps)
    // Output:
    // abcdef
    // abcdef
}
