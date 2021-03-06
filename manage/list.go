package manage

import (
	"fmt"
	"strings"

	"github.com/dtchanpura/deployment-agent/config"
)

// ListProjects for listing the projects
func ListProjects(projects []config.Project) string {
	b := &strings.Builder{}
	if len(projects) > 0 {
		fmt.Fprintln(b, "Following are the projects with their UUIDs")
		for index, project := range projects {
			fmt.Fprintf(b, "%2d. \"%s\", UUID: %s\n", index+1, project.Name, project.UUID)
		}
	} else {
		fmt.Fprintln(b, "No projects found. Add new one using add command.")
	}
	return b.String()
}
