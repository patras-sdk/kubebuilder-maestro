package renderer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"

	"github.com/kudobuilder/kudo/pkg/engine"
	"github.com/kudobuilder/kudo/pkg/test/utils"
	"github.com/kudobuilder/kudo/pkg/util/kudo"
)

func TestEnhancerApply_embeddedMetadata(t *testing.T) {

	tpls := map[string]string{
		"deployment": resourceAsString(statefulSet("sfs1", "default")),
	}

	owner := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "core/v1",
		},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       corev1.PodSpec{},
		Status:     corev1.PodStatus{},
	}
	meta := Metadata{
		Metadata: engine.Metadata{
			InstanceName:        "instance",
			InstanceNamespace:   "namespace",
			OperatorName:        "operator",
			OperatorVersionName: "versionname",
			OperatorVersion:     "1.0.0",
			AppVersion:          "2.0.0",
			ResourcesOwner:      owner,
		},
		PlanName:  "deploy",
		PlanUID:   "uid",
		PhaseName: "phase",
		StepName:  "step",
		TaskName:  "task",
	}
	e := &DefaultEnhancer{
		Scheme: utils.Scheme(),
	}

	objs, err := e.Apply(tpls, meta)
	if err != nil {
		t.Errorf("failed to apply template %s", err)
	}

	for _, o := range objs {
		unstructMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(o)
		if err != nil {
			t.Errorf("failed to parse object to unstructured: %s", err)
		}
		sfs := &appsv1.StatefulSet{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(unstructMap, sfs); err != nil {
			t.Errorf("failed to parse unstructured to StatefulSet: %s", err)
		}

		assert.Equal(t, meta.InstanceNamespace, sfs.GetNamespace())
		assert.Equal(t, string(meta.PlanUID), sfs.Annotations[kudo.PlanUIDAnnotation])
		assert.Equal(t, "kudo", sfs.Labels[kudo.HeritageLabel])

		assert.Equal(t, string(meta.PlanUID), sfs.Spec.Template.Annotations[kudo.PlanUIDAnnotation])
		assert.Equal(t, "kudo", sfs.Spec.Template.Labels[kudo.HeritageLabel])
	}

}

func resourceAsString(resource metav1.Object) string {
	bytes, _ := yaml.Marshal(resource)
	return string(bytes)
}

func statefulSet(name string, namespace string) *appsv1.StatefulSet {
	statefulSet := &appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: name + "Service",
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{},
			},
		},
	}
	return statefulSet
}
