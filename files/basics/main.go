package main
 
import (
    "fmt"
    "log"
    "os"
)
 
func main() {
    newFile, err := os.Create("k.txt")
    if err != nil {
        // log the error and exit the program
        log.Fatal(err)
 
    }

    err = os.Truncate("k.txt", 0) //0 means completely empty the file.
    if err != nil {
        log.Fatal(err)
    }
    newFile.Close()
 
    file, err := os.Open("k.txt") // open in read only mode
    if err != nil {
        log.Fatal(err)
    }
    file.Close()
 
    file, err = os.OpenFile("k.txt", os.O_APPEND, 0644)
    // We can Use opening attributes individually or combined
    // using an OR between them
    // e.g. os.O_CREATE|os.O_APPEND
    // or os.O_CREATE|os.O_TRUNC|os.O_WRONLY
    // os.O_RDONLY // Read only
    // os.O_WRONLY // Write only
    // os.O_RDWR // Read and write
    // os.O_APPEND // Append to end of file
    // os.O_CREATE // Create is none exist
    // os.O_TRUNC // Truncate file when opening
    if err != nil {
        log.Fatal(err)
    }
    file.Close()

    var fileInfo os.FileInfo
    fileInfo, err = os.Stat("k.txt")
 
    p := fmt.Println
    p("File Name:", fileInfo.Name())
    p("Size in bytes:", fileInfo.Size())
    p("Last modified:", fileInfo.ModTime())
    p("Is Directory? ", fileInfo.IsDir())
    p("Pemissions:", fileInfo.Mode())
 
    fileInfo, err = os.Stat("k.txt")
    if err != nil {
        if os.IsNotExist(err) {
            log.Fatal("The file does not exist")
        }
    }
 
    oldPath := "k.txt"
    newPath := "kkkk.txt"
    err = os.Rename(oldPath, newPath)
    if err != nil {
        log.Fatal(err)
    }
 
    err = os.Remove("kkkk.txt")
    if err != nil {
        log.Fatal(err)
    }
 
}