package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/trstringer/k8s-controller-custom-resource/pkg/apis/myresource"
)

// GroupVersion is the identifier for the API which includes
// the name of the group and the version of the API
var SchemeGroupVersion = schema.GroupVersion{
	Group:   myresource.GroupName,
	Version: "v1",
}

// create a SchemeBuilder which uses functions to add types to
// the scheme
// more comments
// even some more
var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// addKnownTypes adds our types to the API scheme by registering
// MyResource and MyResourceList
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&MyResource{},
		&MyResourceList{},
	)

	// register the type in the scheme
	meta_v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
