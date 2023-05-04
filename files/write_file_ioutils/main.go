package main
 
import (
    "io/ioutil"
    "log"
    "os"
)
 
func main() {
    file, err := os.OpenFile(
        "j.txt",
        os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
        0644,
    )
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

 
    byteSlice := []byte("I learn Golang!")   // converting a string to a bytes slice
    bytesWritten, err := file.Write(byteSlice) // writing bytes to file.
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Bytes written: %d\n", bytesWritten)

    // ioutil.WriteFile() handles creating, opening, writing a slice of bytes and closing the file.
    // if the file doesn't exist WriteFile() creates it
    // and if it already exists the function will truncate it before writing to file.
 
    bs := []byte("Go Programming is cool!")
    err = ioutil.WriteFile("d.txt", bs, 0644)
    // error handling
    if err != nil {
        log.Fatal(err)
    }
}