package main

import (
	"fmt"
)

type MyErr struct{}

func (me MyErr) Error() string {
	return "my err string"
}

func main() {
	fmt.Println(returnError() == nil, returnError())
	fmt.Println(returnErrorPtr() == nil, returnErrorPtr())

	fmt.Println(returnCustomError() == nil, returnCustomError())
	fmt.Println(returnCustomErrorPtr() == nil, returnCustomErrorPtr())

	fmt.Println(returnMyError() == nil, returnMyError())
}

func returnError() error {
	var err error
	return err
} // true

func returnErrorPtr() *error {
	var err *error
	return err
} // true

func returnCustomError() error {
	var customErr MyErr
	return customErr
} // false

func returnCustomErrorPtr() error {
	var customErr *MyErr
	return customErr
} // false

func returnMyError() *MyErr {
	return nil
} // true

// не понимаю почему 4ое false если оно nil
