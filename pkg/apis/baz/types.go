package baz

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Foo struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec FooSpec
}

type FooSpec struct {
	// +k8s:conversion-gen=false
	Bar []FooBar
}

type FooBar struct {
	Name string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FooList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Foo
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Bar struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec BarSpec
}

type BarSpec struct {
	// cost is the cost of one instance of this topping.
	Description string
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BarList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Bar
}
