package apis

import (
	"github.com/epmd-edp/go-go-operator-sdk/pkg/apis/v1/v1alpha1"
)

func init() {
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
}
