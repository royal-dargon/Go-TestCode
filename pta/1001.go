package main

import "fmt"

func main() {
        var a,b int
        fmt.Scanf("%d",&a)
        for n := 0; a != 1; {
                if a % 2 == 0 {
                        a = a/2
                } else {
                        a = (3 * a + 1)/2
                }
                n ++
                b = n
        }
        fmt.Printf("%d",b)
}

		

