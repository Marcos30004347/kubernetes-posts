package initializer

import (
	informers "github.com/Marcos30004347/kubernetes-posts/pkg/generated/informers/externalversions"
	"k8s.io/apiserver/pkg/admission"
)

type bazInformerPluginInitializer struct {
	informers informers.SharedInformerFactory
}

var _ admission.PluginInitializer = bazInformerPluginInitializer{}

// New creates an instance of custom admission plugins initializer.
func New(informers informers.SharedInformerFactory) bazInformerPluginInitializer {
	return bazInformerPluginInitializer{
		informers: informers,
	}
}

// Initialize checks the initialization interfaces implemented by a plugin
// and provide the appropriate initialization data
func (i bazInformerPluginInitializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsBazInformerFactory); ok {
		wants.SetBazInformerFactory(i.informers)
	}
}
