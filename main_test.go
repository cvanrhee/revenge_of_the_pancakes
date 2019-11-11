package main

import (
    "strconv"
    "testing"
)

func TestFlipPancakesCount_InvalidValue(t *testing.T) {

    // Should flag as invalid and not flip at all
    result := FlipPancakesCount("invalid_value")

    if  result != 0 {
       t.Error("Expected 0 flips but found: " + strconv.Itoa(result))
    }
}

func TestFlipPancakesCount_Case1(t *testing.T) {

    //Case #1, you only need to execute the maneuver once, flipping the first (and only) pancake.
    result := FlipPancakesCount("-")

    if  result != 1 {
        t.Error("Expected 1 flip but found: " + strconv.Itoa(result))
    }
}


func TestFlipPancakesCount_Case2(t *testing.T) {

    //Case #2, you only need to execute the maneuver once, flipping only the first pancake.
    result := FlipPancakesCount("-+")

    if  result != 1 {
        t.Error("Expected 1 flip but found: " + strconv.Itoa(result))
    }
}


func TestFlipPancakesCount_Case3(t *testing.T) {

    //In Case #3, you must execute the maneuver twice. One optimal solution is to flip only the first pancake,
    //changing the stack to ­­, and then flip both pancakes, changing the stack to ++. Notice that you cannot
    //just flip the bottom pancake individually to get a one­move solution; every time you execute the
    //maneuver, you must select a stack starting from the top.
    result := FlipPancakesCount("+-")

    if  result != 2 {
        t.Error("Expected 2 flips but found: " + strconv.Itoa(result))
    }
}


func TestFlipPancakesCount_Case4(t *testing.T) {

    //Case #4, all of the pancakes are already happy side up, so there is no need to do anything.
    result := FlipPancakesCount("+++")

    if  result != 0 {
        t.Error("Expected 0 flips but found: " + strconv.Itoa(result))
    }
}

func TestFlipPancakesCount_Case5(t *testing.T) {

    //Case #5, one valid solution is to first flip the entire stack of pancakes to get +-++,
    //then flip the top pancake to get --++, then finally flip the top two pancakes to get ++++ .
    result := FlipPancakesCount("--+-")

    if  result != 3 {
        t.Error("Expected 3 flips but found: " + strconv.Itoa(result))
    }
}

func TestIsValidPancake_Invalid(t *testing.T) {

    isValid := IsValidPancake("-+sdas-+")

    if isValid != false {
        t.Error("Expected pancake to be invalid but was valid")
    }

}

func TestIsValidPancake_Valid(t *testing.T) {

    isValid := IsValidPancake("-++--+")

    if isValid != true {
        t.Error("Expected pancake to be valid but was invalid")
    }

}