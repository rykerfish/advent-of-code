package problems

import(
    "fmt"
    "bufio"
    "github.com/rykerfish/advent-of-code/fileIO"
)

func ProblemThree(){
    file := fileIO.CreateFile("3.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)

    tree_map := make([][]rune, 0)

    // Reads in the map
    for scanner.Scan() {

        line := make([]rune, 0)

        for _, char := range(scanner.Text()){
            if char == '\n'{
                break
            }

            line = append(line, char)
        }

        // avoids reading in the newline at the bottom of the file
        if len(line) != 0{
            tree_map = append(tree_map, line)
        }

    }

    tree_product := make([]int, 5)

    down_list := [5]int{1, 1, 1, 1, 2}
    right_list := [5]int{1, 3, 5, 7, 1}

    row_length := len(tree_map[0])

    // Stores the number of trees each slope would encounter in different indices of tree_product
    i := 0
    for i < 5{

        tree_counter := 0

        y := down_list[i]
        x := right_list[i]
        for y < len(tree_map){

            if tree_map[y][x % row_length] == '#'{
                tree_counter++
            }

            y += down_list[i]
            x += right_list[i]

            tree_product[i] = tree_counter
        }
        i++
    }

    product := tree_product[0]
    i = 1
    for i < 5 {
        product *= tree_product[i]
        i++
    }

    fmt.Println("------------Part 1------------")
    fmt.Println("You'd hit", tree_product[1], "trees")
    fmt.Println("------------Part 2------------")
    fmt.Println("The product is", product)



}
