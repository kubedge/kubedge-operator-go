package controller

import (
	"github.com/kubedge/kubedge-operator-arpscan/pkg/controller/arpscan"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, arpscan.AddArpscanController)
}
