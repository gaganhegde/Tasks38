package tasks

import (
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha2"
	corev1 "k8s.io/api/core/v1"
)

// GenerateDeployFromSourceTask will return a github-status-task
func GenerateDeployFromSourceTask() v1alpha1.Task {
	task := v1alpha1.Task{
		TypeMeta:   createTaskTypeMeta(),
		ObjectMeta: createTaskObjectMeta("deploy-from-source-task"),
		Spec: v1alpha1.TaskSpec{
			Inputs: createInputsForDeployFromSourceTask(),
			TaskSpec: v1alpha2.TaskSpec{
				Steps: createStepsForDeployFromSourceTask(),
			},
		},
	}
	return task
}

func createStepsForDeployFromSourceTask() []v1alpha1.Step {
	return []v1alpha1.Step{
		v1alpha1.Step{
			Container: corev1.Container{
				Name:       "run-kubectl",
				Image:      "quay.io/kmcdermo/k8s-kubectl:latest",
				WorkingDir: "/workspace/source",
				Command:    []string{"kubectl"},
				Args:       argsForRunKubectlStep(),
			},
		},
	}
}

func argsForRunKubectlStep() []string {
	return []string{
		"apply",
		"--dry-run=$(inputs.params.DRYRUN)",
		"-n",
		"$(inputs.params.NAMESPACE)",
		"-k",
		"$(inputs.params.PATHTODEPLOYMENT)",
	}
}

func createInputsForDeployFromSourceTask() *v1alpha1.Inputs {
	return &v1alpha1.Inputs{
		Resources: []v1alpha1.TaskResource{
			createTaskResource("source", "git"),
		},
		Params: []v1alpha1.ParamSpec{
			v1alpha1.ParamSpec{
				Name:        "PATHTODEPLOYMENT",
				Description: "Path to the manifest to apply",
				Type:        v1alpha1.ParamTypeString,
				Default: &v1alpha1.ArrayOrString{
					StringVal: "deploy",
				},
			},
			v1alpha1.ParamSpec{
				Name:        "NAMESPACE",
				Type:        v1alpha1.ParamTypeString,
				Description: "Namespace to deploy into",
			},
			v1alpha1.ParamSpec{
				Name:        "DRYRUN",
				Type:        v1alpha1.ParamTypeString,
				Description: "If true run a server-side dryrun.",
				Default: &v1alpha1.ArrayOrString{
					StringVal: "false",
				},
			},
		},
	}
}
