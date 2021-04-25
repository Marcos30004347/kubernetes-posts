package bar

import (
	"github.com/Marcos30004347/kubernetes-posts/pkg/apis/baz"
	"github.com/Marcos30004347/kubernetes-posts/pkg/registry"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

// NewREST returns a RESTStorage object that will work against API services.
func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
	strategy := NewStrategy(scheme)

	store := &genericregistry.Store{
		// Here is where you set that the bars objets are of kind Bar
		NewFunc:     func() runtime.Object { return &baz.Bar{} },
		NewListFunc: func() runtime.Object { return &baz.BarList{} },

		PredicateFunc:            MatchBar,
		DefaultQualifiedResource: baz.Resource("bars"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		TableConvertor: rest.NewDefaultTableConvertor(baz.Resource("bars")),
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &registry.REST{store}, nil
}
