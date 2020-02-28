package main

import (
	"fmt"
	"log"
	"reflect"
)

type User struct {
	Name  string
	Email string
}

func (u *User) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(u)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("name      :", reflectType.Field(i).Name)
		fmt.Println("type      :", reflectType.Field(i).Type)
		fmt.Println("value     :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func main() {
	varInt := 123
	log.Println("variable type", reflect.TypeOf(varInt)) // output int

	varArr := []string{"one", "two", "three"}
	log.Println("variable type", reflect.TypeOf(varArr)) // output []string

	varUser := &User{"John Doe", "john@yahoo.com"}
	log.Println("variable type", reflect.TypeOf(varUser)) // output *main.User

	varUser.getPropertyInfo()
	//output
	// 	name      : Name
	// type      : string
	// value     : John Doe

	// name      : Email
	// type      : string
	// value     : john@yahoo.com

}
