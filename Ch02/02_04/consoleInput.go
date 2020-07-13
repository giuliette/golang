package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {

    // var s string
    // fmt.Scanln(&s)
    // fmt.Println(s)

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("enter txt: ")
    str, _ := reader.ReadString('\n')
    fmt.Println(str)

    fmt.Print("enter no: ")
    str, _ = reader.ReadString('\n')
    f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.PrintLn("value of number", f)
    }
}