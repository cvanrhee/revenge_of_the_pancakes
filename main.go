package main

import "fmt"

func main() {

    //file, err := os.Open("input.in")
    //
    //if err != nil {
    //    log.Fatal(err)
    //}
    //defer file.Close()
    //
    //scanner := bufio.NewScanner(file)
    //
    //for scanner.Scan() {
    //    sortPancakes(scanner.Text())
    //}


    fmt.Println(sortPancakes("5"))
}


func sortPancakes(pancakes string) int{

    flipCount := 0
    prev := pancakes[0]

    for i := range pancakes {
        if pancakes[i] != prev {
            flipCount++
            prev = pancakes[i]
        }
    }


    if prev != '+' {
        flipCount++
    }

    return flipCount
}
