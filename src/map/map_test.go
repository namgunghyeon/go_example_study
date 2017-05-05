package main

import "fmt"
import "sort"

func count(s string, codeCount map[rune]int) {
    for _, r := range s {
        codeCount[r]++
    }
}

func Example_Count1() {
    codeCount := map[rune]int{}
    count("가나다", codeCount)
    for _, key := range []rune{'가', '나', '다'} {
        fmt.Println(string(key), codeCount[key])
    }

    // Output:
    // 가 1
    // 나 1
    // 다 1
}

func Example_Count2() {
    codeCount := map[rune]int{}
    count("가나다", codeCount)
    var keys sort.IntSlice
    for key := range codeCount {
        keys = append(keys, int(key))
    }
    sort.Sort(keys)
    for _, key := range keys {
        fmt.Println(string(key), codeCount[rune(key)])
    }

    // Output:
    // 가 1
    // 나 1
    // 다 1
}

func hasDupeRune(s string) bool {
    runeSet := map[rune]bool{}
    for _, r := range s {
        if runeSet[r] {
            return true
        }
        runeSet[r] = true
    }
    return false
}

func hasDupeRune2(s string) bool {
    runeSet := map[rune]struct{}{}
    for _, r := range s {
        if _, exists := runeSet[r]; exists {
            return true
        }
        runeSet[r] = struct{}{}
    }
    return false
}
