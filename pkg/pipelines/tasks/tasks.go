package tasks

import (
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GenerateTasks will return a slice of tasks
func GenerateTasks() []v1alpha1.Task {
	tasks := []v1alpha1.Task{}
	tasks = append(tasks,
		GenerateGithubStatusTask(),
		GenerateDeployFromSourceTask())
	return tasks
}

func createTaskTypeMeta() v1.TypeMeta {
	return v1.TypeMeta{
		Kind:       "Task",
		APIVersion: "tekton.dev/v1alpha1",
	}
}

func createTaskObjectMeta(name string) v1.ObjectMeta {
	return v1.ObjectMeta{
		Name: name,
	}
}

func createTaskResource(name string, resourceType string) v1alpha1.TaskResource {
	return v1alpha1.TaskResource{
		ResourceDeclaration: v1alpha1.ResourceDeclaration{
			Name: name,
			Type: resourceType,
		},
	}
}

func createEnvFromSecret(name string, localObjectReference string, key string) corev1.EnvVar {
	return corev1.EnvVar{
		Name: name,
		ValueFrom: &corev1.EnvVarSource{
			SecretKeyRef: &corev1.SecretKeySelector{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: localObjectReference,
				},
				Key: key,
			},
		},
	}
}
