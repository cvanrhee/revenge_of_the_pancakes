package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

func main() {

    fileLines := LinesInFile("input.in")

    // Loop and output correct format to show example.
    for i := 1; i < len(fileLines); i++ {
        fmt.Printf("Case #%d: %d", i, FlipPancakesCount(fileLines[i]))
        fmt.Println()
    }
}

// Opens 'fileName' file and returns each line as a string array
func LinesInFile(fileName string) []string {
    f, _ := os.Open(fileName)

    scanner := bufio.NewScanner(f)
    result := []string{}
    // Use Scan.
    for scanner.Scan() {
        line := scanner.Text()
        // Append line to result.
        result = append(result, line)
    }
    return result
}

// Returns the num of times needed to flip pancakes so that they are all right side up.
func FlipPancakesCount(pancakes string) int{

    // If the string contains invalid characters return 0
    if !IsValidPancake(pancakes) {
        return 0
    }

    flipCount := 0
    prevPancake := pancakes[0]

    for i := range pancakes {
        currentPancake := pancakes[i]

        // Flip stack if different so that every pancake thus far is aligned.
        if prevPancake != currentPancake {
            flipCount++
            prevPancake = currentPancake
        }
    }

    // If the pancakes are all wrong side up then flip them all.
    if prevPancake != '+' {
        flipCount++
    }

    return flipCount
}

// Returns false if str contains any character other than '-' or '+'.
func IsValidPancake(str string) bool {
    containsInvalid, _ := regexp.MatchString(`[^\-+]`, str)

    return !containsInvalid
}
