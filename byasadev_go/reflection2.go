package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	id   int
	name string
}

func createQuery1(emp interface{}) {
	if reflect.TypeOf(emp).Kind() == reflect.Struct {
		var query string
		t := reflect.TypeOf(emp).Name()
		query = fmt.Sprintf("insert into %s values (", t)
		v := reflect.ValueOf(emp)
		for i := 0; i < v.NumField(); i++ {
			if i != 0 {
				query = fmt.Sprintf("%s,%v", query, v.Field(i))
			} else {
				query = fmt.Sprintf("%s%v", query, v.Field(i))
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Printf(query)
	}

}

func reflection2() {

	emp := employee{123, "Mcd"}
	createQuery1(emp)
}
