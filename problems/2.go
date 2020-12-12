package main

import(
    "fmt"
    "bufio"
    "strconv"
    "github.com/rykerfish/advent-of-code/fileIO"
)

func main(){

    file := fileIO.CreateFile("2.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)

    validPasswords := 0
    partTwoValidPasswords := 0
    for scanner.Scan(){

        text := scanner.Text()

        // extracts the lower bound
        var lowerNum int
        var dashIndex int
        if(text[1] == '-'){
            lowerNum, _ = strconv.Atoi(string(scanner.Text()[0]))
            dashIndex = 1
        }else{
            lowerNum, _ = strconv.Atoi(string(scanner.Text()[0:2]))
            dashIndex = 2
        }

        // extracts the upper bound
        var upperNum int
        var spaceIndex int
        if(lowerNum >= 10){ // both lowerNum and upperNum are 2 digits long
            upperNum, _ = strconv.Atoi(string(scanner.Text()[3:5]))
            spaceIndex = 5
        }else{ // lowerNum is 1 digit long but upperNum could be 1 or 2 digits long
            if(string(text[dashIndex + 2]) == " "){ // upperNum is 1 digit long
                upperNum, _ = strconv.Atoi(string(scanner.Text()[dashIndex + 1]))
                spaceIndex = dashIndex + 2
            }else{ // upperNum is 2 digits long
                upperNum, _ = strconv.Atoi(string(scanner.Text()[(dashIndex+1):(dashIndex+3)]))
                spaceIndex = dashIndex + 3
            }
        }

        wantedChar := rune(text[spaceIndex + 1])
        text = string(text[(spaceIndex+4):])

        // Counts up the number of wantedChars in the given string
        wantedCharCount := 0
        for _, char := range(text){
            if(wantedChar == char){
                wantedCharCount++
            }
        }

        // deals with Part 1
        if(wantedCharCount >= lowerNum && wantedCharCount <= upperNum){
            validPasswords++
        }

        // deals with Part 2
        posOneHasChar := (wantedChar == rune(text[lowerNum - 1]))
        posTwoHasChar := (wantedChar == rune(text[upperNum - 1]))
        if(posOneHasChar != posTwoHasChar){
            partTwoValidPasswords++
        }

    }

    fmt.Printf("Part one valid passwords: %d\n", validPasswords)
    fmt.Printf("Part two valid passwords: %d\n", partTwoValidPasswords)

}
