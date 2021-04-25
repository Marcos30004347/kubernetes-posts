package foo

import (
	"context"
	"fmt"

	"github.com/Marcos30004347/kubernetes-posts/pkg/apis/baz"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
)

// NewStrategy creates and returns a fooStrategy instance
func NewStrategy(typer runtime.ObjectTyper) fooStrategy {
	return fooStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, the presence of Initializers if any
// and error in case the given runtime.Object is not a Foo
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*baz.Foo)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Foo")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), nil
}

// MatchFoo is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchFoo(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *baz.Foo) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type fooStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (fooStrategy) NamespaceScoped() bool {
	return true
}

func (fooStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (fooStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (fooStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	// if you define a ValidateFoo in the validation package of you api

	// foo := obj.(*baz.Foo)
	// return validation.ValidateFoo(foo)

	return field.ErrorList{}
}

func (fooStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (fooStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (fooStrategy) Canonicalize(obj runtime.Object) {
}

func (fooStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
