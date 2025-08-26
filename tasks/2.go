package main

import (
	"fmt"
)

type MyErr struct{}

func (me MyErr) Error() string {
	return "my err string"
}

func main() {
	fmt.Println(returnError() == nil)
	fmt.Println(returnErrorPtr() == nil)

	fmt.Println(returnCustomError() == nil)
	fmt.Println(returnCustomErrorPtr() == nil)

	fmt.Println(returnMyError() == nil)
}

func returnError() error {
	var err error
	return err
}

func returnErrorPtr() *error {
	var err *error
	return err
}

func returnCustomError() error {
	var customErr MyErr
	return customErr
}

func returnCustomErrorPtr() error {
	var customErr *MyErr
	return customErr
}

func returnMyError() *MyErr {
	return nil
}
