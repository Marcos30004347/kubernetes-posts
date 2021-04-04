package initializer

import (
	"k8s.io/apiserver/pkg/admission"

	informers "github.com/Marcos30004347/kubernetes-posts/pkg/generated/informers/externalversions"
)

// WantsBazInformerFactory defines a function which sets InformerFactory for admission plugins that need it
type WantsBazInformerFactory interface {
	SetBazInformerFactory(informers.SharedInformerFactory)
	admission.InitializationValidator
}
