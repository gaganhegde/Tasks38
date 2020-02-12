package tasks

import (
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha2"
	corev1 "k8s.io/api/core/v1"
)

// GenerateGithubStatusTask will return a github-status-task
func GenerateGithubStatusTask() v1alpha1.Task {
	task := v1alpha1.Task{
		TypeMeta:   createTaskTypeMeta(),
		ObjectMeta: createTaskObjectMeta("create-github-status-task"),
		Spec: v1alpha1.TaskSpec{
			Inputs: createInputsForGithubStatusTask(),
			TaskSpec: v1alpha2.TaskSpec{
				Steps: createStepsForGithubStatusTask(),
			},
		},
	}
	return task
}

func argsForStartStatusStep() []string {
	return []string{
		"create-status",
		"--repo",
		"$(inputs.params.REPO)",
		"--sha",
		"$(inputs.params.COMMIT_SHA)",
		"--state",
		"$(inputs.params.STATE)",
		"--target-url",
		"$(inputs.params.TARGET_URL)",
		"--description",
		"$(inputs.params.DESCRIPTION)",
		"--context",
		"$(inputs.params.CONTEXT)",
	}
}

func createStepsForGithubStatusTask() []v1alpha1.Step {
	return []v1alpha1.Step{
		v1alpha1.Step{
			Container: corev1.Container{
				Name:       "start-status",
				Image:      "quay.io/kmcdermo/github-tool:latest",
				WorkingDir: "/workspace/source",
				Env: []corev1.EnvVar{
					createEnvFromSecret("GITHUB_TOKEN", "github-auth", "token"),
				},
				Command: []string{"github-tools"},
				Args:    argsForStartStatusStep(),
			},
		},
	}
}

func createInputsForGithubStatusTask() *v1alpha1.Inputs {
	inputs := v1alpha1.Inputs{
		Params: []v1alpha1.ParamSpec{
			v1alpha1.ParamSpec{
				Name:        "REPO",
				Type:        v1alpha1.ParamTypeString,
				Description: "The repo to publish the status update for e.g. tektoncd/triggers",
			},
			v1alpha1.ParamSpec{
				Name:        "COMMIT_SHA",
				Type:        v1alpha1.ParamTypeString,
				Description: "The specific commit to report a status for.",
			},
			v1alpha1.ParamSpec{
				Name:        "STATE",
				Type:        v1alpha1.ParamTypeString,
				Description: "The state to report error, failure, pending, or success.",
			},
			v1alpha1.ParamSpec{
				Name:        "TARGET_URL",
				Type:        v1alpha1.ParamTypeString,
				Description: "The target URL to associate with this status.",
				Default: &v1alpha1.ArrayOrString{
					StringVal: "",
				},
			},
			v1alpha1.ParamSpec{
				Name:        "DESCRIPTION",
				Type:        v1alpha1.ParamTypeString,
				Description: "A short description of the status.",
			},
			v1alpha1.ParamSpec{
				Name:        "CONTEXT",
				Type:        v1alpha1.ParamTypeString,
				Description: "A string label to differentiate this status from the status of other systems.",
			},
		},
	}
	return &inputs
}
