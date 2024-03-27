package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// implementare citire din fisier mai lunga
	file, err := os.Open("goFiles\\test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)

	// implemenatre citire din fisier mai scurta
	bs2, err := ioutil.ReadFile("goFiles\\test.txt")
	if err != nil {
		return
	}
	for ind, cuv := range bs2 {
		fmt.Printf("cuvantul pe pozitia i:%v este:%c\n", ind, cuv)
	}

	str2 := string(bs2)
	cuv := strings.Fields(str2)
	fmt.Printf("%v\n", cuv)

	//creating a file
	file2, err := os.Create("goFiles\\test2.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file2.Close()
	file2.WriteString("I AM A COCONUT")

	//sa deschidem un dir si sa afisam ce e in el
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
