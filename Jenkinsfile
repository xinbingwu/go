pipeline {
	agent any 

	tools {
		maven  'mvn-3.5.4'
	}	
	stages {
		stage("输出文字"){
			steps {
				sh "mvn clean package spring-boot:repackage"
				sh "printenv"
			
			}
	
		}
	

	}

}
