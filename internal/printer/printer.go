package printer

import "github.com/gungorugur/ciqube/internal/sonarqube"

//Printer interface
type Printer interface {
	Print()
}

//NewProjectStatusPrinter returns new project status printer
func NewProjectStatusPrinter(projectStatus *sonarqube.ProjectStatus) Printer {

	printer := projectStatusPrinter{projectStatus: projectStatus}

	return printer
}
