package bar

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

// NewStrategy creates and returns a barStrategy instance
func NewStrategy(typer runtime.ObjectTyper) barStrategy {
	return barStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, the presence of Initializers if any
// and error in case the given runtime.Object is not a Bar
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*baz.Bar)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Bar")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), nil
}

// MatchBar is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchBar(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *baz.Bar) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type barStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (barStrategy) NamespaceScoped() bool {
	return true
}

func (barStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (barStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

// Here is where we actually use the Validate Function defined in the api
func (barStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}

}

func (barStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (barStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (barStrategy) Canonicalize(obj runtime.Object) {
}

func (barStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
