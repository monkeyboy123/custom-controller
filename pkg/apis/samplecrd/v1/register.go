package v1
import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
func addKnownTypes(scheme *runtime.Scheme) error {
scheme.AddKnownTypes(
SchemeGroupVersion,
&Network{},
&NetworkList{},
)
metav1.AddGroupVersion(Scheme,SchemeGroupVersion)
return nil

}
