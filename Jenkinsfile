pipeline {
	agent any
	environment {
		GOPATH =  "${env.WORKSPACE}/"
	}

	tools {
		go 'go-new'
	}

	stages {
		stage('build'){
			steps{
				sh "go build"
				sh "printenv"
			}
		}
	}

}
