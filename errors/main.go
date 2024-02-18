package main

import "fmt"

//var _ error = AppError{}

type appError struct {
	Err    error
	Custom string
	Field  int
}

func (e *appError) Error() string {
	//err := fmt.Errorf("new error %ws", e.Custom)\
	//fmt.Println(e.Custom)
	return e.Err.Error()
}

func (e *appError) Unwrap() error {
	return e.Err
}

func main() {
	err := method1()
	if err != nil {
		fmt.Println(err.Custom)
		fmt.Println(err.Unwrap())
		return
	}
}

func method1() *appError {
	return method2()
}

func method2() *appError {
	return &appError{
		Err:    fmt.Errorf("method2 failed"),
		Custom: "Make sure... =D",
		Field:  34,
	}
}
