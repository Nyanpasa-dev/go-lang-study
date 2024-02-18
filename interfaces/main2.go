package main

import "fmt"

var _ User = &superUser{}

type user struct {
	FIO, address, phone string
	isBlocked           bool
}

type superUser struct {
	name      string
	age       int
	isBlocked bool
}

func (s *superUser) block() {
	s.isBlocked = true
}

func (s *superUser) changeFIO(newFio string) string {
	//TODO implement me
	panic("implement me")
}

func (s *superUser) changeAddress(newAddress string) string {
	//TODO implement me
	panic("implement me")
}

func (u *user) changeFIO(newFio string) string {
	u.FIO = newFio
	return u.FIO
}

func (u *user) changeAddress(newAddress string) string {
	u.address = newAddress
	return u.address
}

func (u *user) block() {
	u.isBlocked = true
}

type User interface {
	changeFIO(newFio string) string
	changeAddress(newAddress string) string
	block()
}

func newUser(fio, address, phone string) User {
	u := user{
		FIO:     fio,
		address: address,
		phone:   phone,
	}

	return &u
}

func main() {
	userStruct := newUser("nyanpasa", "pumba 12", "12112")

	a := userStruct.changeAddress("Valhalla 13")
	b := userStruct.changeFIO("Nyanpasa Verflucht")
	userStruct.block()

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(userStruct)

}
