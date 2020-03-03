package printer

import (
	"os"

	"github.com/gungorugur/ciqube/internal/sonarqube"
	"github.com/jedib0t/go-pretty/table"
)

//ProjectStatusPrinter printer implementation
type projectStatusPrinter struct {
	projectStatus *sonarqube.ProjectStatus
}

//Print prints sonarqube result table to stdout
func (p projectStatusPrinter) Print() {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Metric", "Comparator", "Error Threshold", "Actual Value", "Status"})

	for _, element := range p.projectStatus.Conditions {
		t.AppendRow([]interface{}{
			sonarqube.GetMetricDefinition(element.MetricKey),
			element.Comparator,
			element.ErrorThreshold,
			element.ActualValue,
			element.Status})
	}

	var gateResult string

	if p.projectStatus.IsPassed() {
		gateResult = "PASSED"
	} else {
		gateResult = "FAILED"
	}

	t.AppendFooter(table.Row{"", "", "", "Quality Gate", gateResult})
	t.Render()
}
