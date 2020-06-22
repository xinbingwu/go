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
			
			options {
				timeout(time: 3, unit: 'MINUTES')
	
			}

			steps {
				script {
					def browsers = ['chrome','firefox']
					for (int i=0; i<browsers.size(); ++i){
						echo "Testing the ${browsers[i]} browser"
				         }
	
					 println env.WORKSPACE
					 println env
					 dir("/tmp/jenkins/"){
						pwd
						deleteDir()
					}	
			      }
			}
			
	

		}	
	}

	post {
		always {
			mail to: 'xinbingwu@yeah.net', subject: 'The pipeline is failed!', body: "Something is wrong with ${env.BUILD_URL}"

		}
	}

}




