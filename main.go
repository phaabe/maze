package main

import (
    "fmt"
    "os"
    "strconv"
)

func contains(list []string, item string) bool {
    for _, element := range list {
        if item == element{
            return true
        }
    }
    return false
}

func main() {
    // intro
    if len(os.Args) <= 1 { // there is alaways one argument
        fmt.Println("Follow the maze and find the bigger picture!")
        fmt.Println("Have fun! If impatient use 'hint'!")
        os.Exit(1)
    }

    // hint section
    if os.Args[1] == "hint" {
        fmt.Println("Use the arguments N,S,E,W to move in the maze!")
        fmt.Println("Use 'hint2' for another hint.")
        os.Exit(1)
    }
    if os.Args[1] == "hint2" {
        fmt.Println("Provide your position as second argument.")
        os.Exit(1)
    }

    // init vars
    var position int
    var move string

    // init directions and ... the maze!
    directions := []string{"N", "S", "E", "W"}
    the_maze := []string{
        "N", "W", "N", "W", "E", "N", "S" , "E" ,
        "S" , "E" , "N" , "N" , "S" , "E",
    }

    // navigation
    move = os.Args[1]

    // valid direction?
    if contains(directions, move){

        // init and validate position
        if len(os.Args) <= 2 {
            position = 0
        } else {
            var err error
            position, err = strconv.Atoi(os.Args[2])
            if err != nil {
                fmt.Println("oh, another argument, nice! but no")
                os.Exit(1)
            }
        }

        if position > len(the_maze) - 1{
            fmt.Println("Impossible")
            os.Exit(1)
        }

        if position == len(the_maze) - 1{
            fmt.Println("Congratulations, that's the end of the maze! But do you remember your initial task?")
            os.Exit(1)
        }

        // move or wall?
        if the_maze[position] == move {
            fmt.Println(
                "You moved from position", position, 
                "to position", position + 1,
            )
        } else {
            fmt.Println(
                "*bounce* that's a wall",
            )
        }
    }
}

