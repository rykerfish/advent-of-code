package problems

import(
    "fmt"
    "bufio"
    "github.com/rykerfish/advent-of-code/fileIO"
    "strings"
    "strconv"
    "regexp"
)

type passport struct{
    fields map[string]string
}

func ProblemFour(){

    file := fileIO.CreateFile("4.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)

    //passports := make([]passport, 0)


    validated_passports := 0
    valid_passports := 0
    for scanner.Scan(){

        // Adds all the key:value pair from one passport to a passport struct
        var str string
        str = scanner.Text()
        var pass passport
        pass.fields = make(map[string]string)
        for str != ""{
            text := scanner.Text()
            text = strings.TrimSpace(text)

            line := strings.Split(text, " ")

            // Splits key:value pairs and adds them to the passport map
            for _, entry := range(line){
                pair := strings.Split(entry, ":")
                pass.fields[pair[0]] = pair[1]
            }

            scanner.Scan()
            str = scanner.Text()
        }

        cid_flag := false
        for key := range(pass.fields){
            if key == "cid"{
                cid_flag = true
            }
        }

        correct_fields := false
        if len(pass.fields) == 8 || (len(pass.fields) == 7 && !cid_flag){
            valid_passports++
            correct_fields = true
        }

        if(correct_fields){ // only try to validate the passport fields if all the correct fields are present
            if(validate_passport(pass)){
                validated_passports++
            }
        }


    }

    fmt.Println("------------Part 1------------")
    fmt.Println(valid_passports)
    fmt.Println("------------Part 2------------")
    fmt.Println(validated_passports)
}


func validate_passport(pass passport) bool{

    m := pass.fields

    byr, _ := strconv.Atoi(m["byr"])
    if(byr < 1920 || byr > 2002 || len(m["byr"]) != 4){
        return false
    }

    iyr, _ := strconv.Atoi(m["iyr"])
    if(iyr < 2010 || iyr > 2020 || len(m["iyr"]) != 4){
        return false
    }

    eyr, _ := strconv.Atoi(m["eyr"])
    if(eyr < 2020 || eyr > 2030 || len(m["eyr"]) != 4){
        return false
    }

    ht_len := len(m["hgt"])
    unit := m["hgt"][ht_len-2:ht_len]
    hgt, _ := strconv.Atoi(m["hgt"][0:ht_len-2])
    if(unit == "in"){
        if(hgt < 59 || hgt > 76){
            return false
        }
    }else if(unit == "cm"){
        if(hgt < 150 || hgt > 193){
            return false
        }
    }else{
        return false
    }

    hcl := m["hcl"]
    print(hcl)
    if(hcl[0] == '#' && len(hcl) == 7){
        for _, char := range(hcl[1:]){
            matched, _ := regexp.Match("[0-9a-f]", []byte(string(char)))
            if(!matched){
                return false
            }
        }
    }else{
        return false
    }

    ecl := m["ecl"]
    switch(ecl){
    case "amb":
    case "blu":
    case "brn":
    case "gry":
    case "grn":
    case "hzl":
    case "oth":
        break
    default:
        return false
    }

    pid := m["pid"]
    matched, _ := regexp.Match("[0-9]{9}?$", []byte(pid))
    if(!matched || len(pid) != 9){
        return false
    }

    return true
}
