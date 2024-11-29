package main

import (
	"fmt"
	"reflect"
)

type MyError struct {
	message string
}

func (e MyError) Error() string {
	return e.message
}

type MyInterface interface {
	GetSomething() int64
}

type MyStruct struct{}

func (l MyStruct) GetSomething() int64 { return 0 }

func main() {
	interfaceAssignment()
	nilError()
}

func interfaceAssignment() {
	var interf MyInterface
	if interf == nil {
		fmt.Println("nil interface")
	}

	var struc *MyStruct
	if struc == nil {
		fmt.Println("nil struct")
	}

	interf = struc
	if interf == nil || reflect.ValueOf(interf).IsNil() { // <- тут добавляем reflect
		fmt.Println("nil assignment")
	}
}

func nilError() {
	err := getErrorOrNil(0)
	if err != nil {
		fmt.Println("0 oops")
	} else {
		fmt.Println("0 ok")
	}

	err = getErrorOrNil(10)
	if err != nil {
		fmt.Println("10 oops")
	} else {
		fmt.Println("10 ok")
	}
}

func getErrorOrNil(i int) error {
	var err *MyError
	if i == 0 {
		err = &MyError{message: "error i=0"}
		return err
	}

	return err // <------------- тут надо вернуть nil
}
