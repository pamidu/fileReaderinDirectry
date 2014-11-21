package main

import (
    "fmt"
    //"io/ioutil"
    "path/filepath"

)

func main() {
    /*files, _ := ioutil.ReadDir("./")
    for _, f := range files {
            fmt.Println(f.Name())
    }

*/
	files1, _ := filepath.Glob("*.config")
    //fmt.Println(files1) // contains a li
    for files:=range files1{
    	fmt.Println(files1[files])
    }

}