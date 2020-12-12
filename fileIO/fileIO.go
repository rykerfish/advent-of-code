package fileIO

import(
    "log"
    "os"
)

func CreateFile(fileName string) *os.File {

    prefix := "/home/rfish/Desktop/go_workspaces/AOC2020/src/files/"

    fullPath := prefix + fileName

    file, err := os.Open(fullPath)

    if err != nil {
        log.Fatal("Failed to open file")
    }

    return file
}
