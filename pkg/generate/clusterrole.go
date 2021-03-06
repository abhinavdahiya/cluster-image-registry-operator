package generate

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	authapi "github.com/openshift/api/authorization/v1"

	"github.com/openshift/cluster-image-registry-operator/pkg/apis/dockerregistry/v1alpha1"
	"github.com/openshift/cluster-image-registry-operator/pkg/strategy"
)

func ClusterRole(cr *v1alpha1.OpenShiftDockerRegistry) Template {
	role := &authapi.ClusterRole{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "ClusterRole",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "system:registry",
		},
		Rules: []authapi.PolicyRule{
			{
				Verbs:     []string{"list"},
				APIGroups: []string{""},
				Resources: []string{
					"limitranges",
					"resourcequotas",
				},
			},
			{
				Verbs:     []string{"get"},
				APIGroups: []string{ /* "", */ "image.openshift.io"},
				Resources: []string{
					"imagestreamimages",
					/* "imagestreams/layers", */
					"imagestreams/secrets",
				},
			},
			{
				Verbs:     []string{ /* "list", */ "get", "update"},
				APIGroups: []string{ /* "", */ "image.openshift.io"},
				Resources: []string{
					"imagestreams",
				},
			},
			{
				Verbs:     []string{ /* "get", */ "delete"},
				APIGroups: []string{ /* "", */ "image.openshift.io"},
				Resources: []string{
					"imagestreamtags",
				},
			},
			{
				Verbs:     []string{"get", "update" /*, "delete" */},
				APIGroups: []string{ /* "", */ "image.openshift.io"},
				Resources: []string{
					"images",
				},
			},
			{
				Verbs:     []string{"create"},
				APIGroups: []string{ /* "", */ "image.openshift.io"},
				Resources: []string{
					"imagestreammappings",
				},
			},
		},
	}
	addOwnerRefToObject(role, asOwner(cr))
	return Template{
		Object:   role,
		Strategy: strategy.Override{},
	}
}
