pipeline {
	agent any
	stages {
		stage("build with go-new"){
			tools {
				go 'go-new'
			}
			environment {
				GOPATH = "${env.WORKSPACE}/"	
			}
			steps {
				sh "printenv"
				sh "go build"
			}
		}


		stage("build with go-1.13"){
			tools {
				go 'go-1.13'
			}
			environment {
				GOPATH = "${env.WORKSPACE}/subdir/"
			}
			steps {
				sh "go build"
			}
		}
	}

}



