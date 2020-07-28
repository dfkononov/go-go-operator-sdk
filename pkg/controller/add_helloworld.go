package controller

import (
	"github.com/epmd-edp/go-go-operator-sdk/pkg/controller/helloworld"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, helloworld.Add)
}
