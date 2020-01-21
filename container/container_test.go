package container

import (
	"testing"

	"github.com/akarasso/Golang_OAuth/libcontainer"
)

func TestNewContainer(t *testing.T) {
	libcontainer.Build(&Dependencies)
}
