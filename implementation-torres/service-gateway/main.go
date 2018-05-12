package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println("hello")
	err := errors.New("Error~!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println(err)
}

