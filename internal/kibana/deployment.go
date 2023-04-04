package kibana

import (
	"github.com/openshift/elasticsearch-operator/internal/constants"
	"github.com/openshift/elasticsearch-operator/internal/manifests/deployment"
	"github.com/openshift/elasticsearch-operator/internal/utils"
	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewDeployment stubs an instance of a Deployment
func NewDeployment(deploymentName string, namespace string, loggingComponent string, component string, replicas int32, podSpec core.PodSpec) *apps.Deployment {
	labels := utils.CommonLabels("Kibana", component, loggingComponent)

	kibanaDeployment := deployment.New(constants.Kibana, namespace, labels, replicas).
		WithSelector(metav1.LabelSelector{
			MatchLabels: labels,
		}).
		WithStrategy(apps.RollingUpdateDeploymentStrategyType).
		WithTemplate(core.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Name:   deploymentName,
				Labels: labels,
			},
			Spec: podSpec,
		}).
		Build()

	return kibanaDeployment
}
