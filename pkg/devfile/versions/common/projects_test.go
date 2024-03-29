package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValidateProjects(t *testing.T) {

	t.Run("No projects present", func(t *testing.T) {

		// Empty projects
		projects := []DevfileProject{}

		got := ValidateProjects(projects)
		want := fmt.Errorf(ErrorNoProjects)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: '%v', got: '%v'", want, got)
		}
	})

	t.Run("Valid project type", func(t *testing.T) {

		// Valid project type
		projects := []DevfileProject{
			{Source: DevfileProjectSource{Type: DevfileProjectTypeGit}},
		}

		got := ValidateProjects(projects)

		if got != nil {
			t.Errorf("Error '%v' not expected", got)
		}
	})

	t.Run("Invalid project type", func(t *testing.T) {

		// Invalid project type
		projects := []DevfileProject{
			{Source: DevfileProjectSource{Type: DevfileProjectType("invalidType")}},
		}

		got := ValidateProjects(projects)
		want := fmt.Errorf(fmt.Sprintf(ErrorInvalidProjectType, projects[0].Source.Type))

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected error, didn't get one")
		}
	})
}
