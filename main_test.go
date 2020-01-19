package main

import (
	"testing"

	"github.com/akarasso/Golang_OAuth/container"
	"github.com/akarasso/Golang_OAuth/model"
)

type dependencies struct {
	s2sPixel *model.MongoCollection `db:"mongo" collection:"s2s_pixel"`
}

func TestNewContainer(t *testing.T) {
	c := dependencies{}
	container.Build(&c)
}
