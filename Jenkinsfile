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

		stage('retry'){i
			steps{
				retry(5){
					script  {
						sh(script: 'curl http://example', returnStatus:true)
					}
				}
			}


		}
	}

}


