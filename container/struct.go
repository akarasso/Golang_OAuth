package container

import (
	"log"
	"reflect"
)

type MongoCollection struct {
}

func Build(i interface{}) {

	container := reflect.Indirect(reflect.ValueOf(i))
	for i := 0; i < container.Type().NumField(); i++ {
		log.Println(container.Type().Field(i).Type == reflect.ValueOf(MongoCollection{}).Type())
	}
}
