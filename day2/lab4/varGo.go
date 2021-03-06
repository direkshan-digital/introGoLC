package main

import (
    "flag"
    "fmt"
    "time"
)

func main() {
    n := flag.Int("n", 20, "# goroutines")
    flag.Parse()
    count := *n
    fmt.Printf("Creating %d goroutines.\n", count)

   for i := 0; i < count; i++ {
       go func(x int) {
           fmt.Printf("%d ", x)
        }(i)
    }

    time.Sleep(30000)
    fmt.Println("\nExiting...")
}

