package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fs, err := ioutil.ReadDir("/")
	if err != nil {
		panic(err)
	}

	for _, f := range fs {
		fmt.Printf("%s\n", f.Name())
	}
}
