package v1alpha1

import (
	"github.com/Marcos30004347/kubernetes-posts/pkg/apis/baz"
	"k8s.io/apimachinery/pkg/conversion"
)

func Convert_v1alpha1_FooSpec_To_baz_FooSpec(in *FooSpec, out *baz.FooSpec, s conversion.Scope) error {
	for _, bar := range in.Bar {
		out.Bar = append(out.Bar, baz.FooBar{
			Name: bar,
		})
	}

	return nil
}

func Convert_baz_FooSpec_To_v1alpha1_FooSpec(in *baz.FooSpec, out *FooSpec, s conversion.Scope) error {
	for i := range in.Bar {
		out.Bar = append(out.Bar, in.Bar[i].Name)
	}

	return nil
}
