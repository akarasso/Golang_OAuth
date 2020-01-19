package container

import (
	"log"
	"reflect"

	"github.com/akarasso/Golang_OAuth/model"
	"github.com/alexsasharegan/dotenv"
)

func isMongoDbField(field reflect.StructField) bool {
	return field.Type == reflect.ValueOf(&model.MongoCollection{}).Type()
}

// Build :: initialiser container
func Build(i interface{}) {
	err := dotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	containerValue := reflect.ValueOf(i)
	container := reflect.Indirect(containerValue)
	for i := 0; i < container.Type().NumField(); i++ {
		field := container.Type().Field(i)
		if isMongoDbField(container.Type().Field(i)) == true {
			log.Printf(`Dependencies: Loading "%s" collection from "%s" database`,
				field.Tag.Get("collection"),
				field.Tag.Get("db"))
			collection := model.NewMongoCollection(field.Tag.Get("db"), field.Tag.Get("collection"))
			containerValue.Elem().Field(i).Set(reflect.ValueOf(collection))
		}
	}
}
