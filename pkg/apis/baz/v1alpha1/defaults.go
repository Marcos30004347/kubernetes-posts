package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

func SetDefaults_FooSpec(obj *FooSpec) {
	if len(obj.Bar) == 0 {
		obj.Bar = []string{"foo0", "foo1", "foo2"}
	}
}
