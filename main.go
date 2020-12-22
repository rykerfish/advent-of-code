package main

import(
    "os"
    // "fmt"
    problems "github.com/rykerfish/advent-of-code/problems"
)

func call( f func() ){
    f()
}

func main(){

    m := map[string]func() {
        "1": problems.ProblemOne,
        "2": problems.ProblemTwo,
        "3": problems.ProblemThree,
        "4": problems.ProblemFour,
    }

    problemNum := os.Args[1]

    call(m[problemNum])
}
