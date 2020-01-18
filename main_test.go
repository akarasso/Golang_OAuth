package main

import (
	"testing"

	"github.com/akarasso/oauth_server/container"
)

type dependencies struct {
	s2s_pixel container.MongoCollection `db:"mongo"`
}

func TestNewContainer(t *testing.T) {
	c := dependencies{}
	container.Build(&c)
}
