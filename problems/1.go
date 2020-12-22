package problems

import(
    "fmt"
    "bufio"
    "strconv"
    "github.com/rykerfish/advent-of-code/fileIO"
)

func ProblemOne(){

    file := fileIO.CreateFile("1.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var slice []int

    for scanner.Scan(){
        num, _ := strconv.Atoi(scanner.Text())
        slice = append(slice, num)
    }

    var product int

    for i := range slice {
        for j := range slice {
            if i == j {continue}
            if slice[i] + slice[j] == 2020 {
                product = slice[i] * slice[j]
                break
            }
        }
    }

    fmt.Println("------------Part 1------------")
    fmt.Println(product)

    for i := range slice {
        for j := range slice {
            if i == j {continue}
            for k := range slice {
                if i == k || j == k {continue}

                if slice[i] + slice[j] + slice[k] == 2020 {
                    product = slice[i] * slice[j] * slice[k]
                }
            }
        }
    }

    fmt.Println("------------Part 2------------")
    fmt.Println(product)
}
