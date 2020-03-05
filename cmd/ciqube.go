package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gungorugur/ciqube/internal/printer"
	"github.com/gungorugur/ciqube/internal/sonarqube"
	"github.com/gungorugur/ciqube/pkg/httpclient"
)

func main() {
	host := flag.String("host", "", "Sonarqube Host URL (Required)")
	token := flag.String("token", "", "Sonarqube Token (Required)")
	projectKey := flag.String("projectkey", "", "Sonarqube Project Key (Required)")
	fail := flag.Bool("fail", false, "Pipeline fails, if quality gateway fails (Optional)")
	waitProgress := flag.Bool("wait-progress", false, "Wait for background tasks if exists (Optional)")
	waitProgressTimeout := flag.Int("timeout", 300, "Timeout for wait progress in seconds (Optional)")
	flag.Parse()

	if *waitProgress {
		err := sonarqube.CheckBackgroundTask(*host, *token, *projectKey, *waitProgressTimeout, httpclient.New(http.DefaultClient))
		exitOnError(err)
	}

	projectStatus, err := sonarqube.GetQualityGateProjectStatus(*host, *token, *projectKey, httpclient.New(http.DefaultClient))

	exitOnError(err)

	printer.NewProjectStatusPrinter(projectStatus).Print()

	exit(*fail, projectStatus.IsPassed())
}

func exit(fail bool, isPassed bool) {

	if fail && isPassed {
		exitWithSuccess()
	} else if fail {
		exitWithError()
	} else {
		exitWithSuccess()
	}
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		exitWithError()
	}
}

func exitWithSuccess() {
	os.Exit(0)
}

func exitWithError() {
	os.Exit(-1)
}
