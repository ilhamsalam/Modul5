package main

import "fmt"

func power(x, y int) int {
        if y == 0 {
                return 1
        } else if y > 0 {
                return x * power(x, y-1)
        } else {
                return 1 / power(x, -y)
        }
}

func main() {
        var base, exponent int
        fmt.Print("Masukkan bilangan pokok (x): ")
        fmt.Scan(&base)
        fmt.Print("Masukkan pangkat (y): ")
        fmt.Scan(&exponent)

        result := power(base, exponent)
        fmt.Println(base, "pangkat", exponent, "adalah", result)
}