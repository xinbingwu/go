pipeline {
	agent any
	
	stages {
		stage('docker') {
			steps {
				script { 
					def t = tool(name:'docker')
					echo "${t}"
				}
			}

		}

	}
}


