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

		stage('retry'){
			steps{
				retry(5){
					script  {
						sh(script: 'curl http://example', returnStatus:true)
						for(i=0; i<5; ++i){
							echo "这是第${i}个数"
						}
					}
				}
			}


		}
	}

}


