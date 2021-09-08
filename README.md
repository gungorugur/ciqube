#Important!

Sonarqube now supports quality gate checking.

https://docs.sonarqube.org/latest/analysis/ci-integration-overview/

This repository archived.

#Ciqube

A simple sonarqube quality gate result checking tool. Put your sonarqube quality gate results to ci/cd pipelines, break pipelines according to quality gate result if you want to.

## Installation and Usage

Ciqube designed for working in ci/cd pipelines. You can simply use it in CircleCi, Gitlab Ci, Azure DevOps etc because they all already support docker agents or dockerfiles. Checkout dockerhub for related images.

[Ciqube Dockerhub](https://hub.docker.com/r/ugurgungor/ciqube)

### Example Jenkins Pipeline
```
pipeline {
    agent none
    stages {
        stage('Quality Gate Check') {
            
            agent {
                docker {
                    image "ugurgungor/ciqube"
                }
            }
           
        steps {       
            sh "ciqube --host http://sonarqube.mycompany.com --token MY_TOKEN --projectkey  MY_PROJECT_KEY --fail"
          }
        }
    }
}
```

You need to pass host, projectkey and sonarqube token. If you pass --fail parameter pipeline will break according to quality gate result. If you want to wait for background task completing then pass to --wait-progress parameter.

```
  -fail
        Pipeline fails, if quality gateway fails (Optional)
  -host string
        Sonarqube Host URL (Required)
  -projectkey string
        Sonarqube Project Key (Required)
  -timeout int
        Timeout for wait progress in seconds (Optional) (default 300)
  -token string
        Sonarqube Token (Required)
  -wait-progress
        Wait for background tasks if exists (Optional)

```

#### Example succeeded pipeline without fail parameter.
<img src="https://raw.githubusercontent.com/gungorugur/ciqube/master/assets/pipelinedemo1.jpeg" width="600">

#### Example succeeded pipeline with fail parameter.
<img src="https://raw.githubusercontent.com/gungorugur/ciqube/master/assets/pipelinedemo2.jpeg" width="600">

#### Example failed pipeline with fail parameter.
<img src="https://raw.githubusercontent.com/gungorugur/ciqube/master/assets/pipelinedemo3.jpeg" width="600">

#### Example succeeded pipeline with wait-progress parameter.
<img src="https://raw.githubusercontent.com/gungorugur/ciqube/master/assets/pipelinedemo4.jpeg" width="600">


### TODO

- Add more tests
