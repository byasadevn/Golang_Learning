package main

import (
	"fmt"
	"reflect"
)

type emp0 struct {
	id   int
	name string
}

func test1(emp emp0) {

	if reflect.ValueOf(emp).Kind() == reflect.Struct {
		v := reflect.ValueOf(emp)
		for i := 0; i < v.NumField(); i++ {

			fmt.Printf("The %d : %T : %v\n", i, v.Field(i), v.Field(i))
		}

	}

	/*t := reflect.TypeOf(emp)
	fmt.Println(t.Kind())
	v := reflect.ValueOf(emp)
	fmt.Println(v.Kind())
	fmt.Println(t)
	fmt.Println(v)*/

}
func createQuery(e interface{}) {
	if reflect.ValueOf(e).Kind() == reflect.Struct {
		var query string
		t := reflect.TypeOf(e).Name()
		query = fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(e)
		for i := 0; i < v.NumField(); i++ {

			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s,%d", query, v.Field(i).Int())
				}

			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s%s", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s,%s", query, v.Field(i).String())
				}
			default:
				fmt.Println("Not supported")
			}
		}
		query = fmt.Sprintf("%s%s", query, ")")
		fmt.Println(query)

	}

}
func reflection1() {

	e := emp0{123, "basu"}
	createQuery(e)
}
