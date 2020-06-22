pipeline {
	agent any 

        options {
		retry(4)
	}

	tools {
		maven  'mvn-3.5.4'
	}	
	stages {
		stage("输出文字"){
			steps {
				sh "mvn clean package spring-boot:repackage"
				sh "printenv"
			
			}

			options {
				timeout(time: 3, unit: 'MINUTES')
	
			}
	

		}	
	}

	post {
		failure {
			mail to: 'xinbingwu@yeah.net', subject: 'The pipeline is failed!'
		}
	}

}




