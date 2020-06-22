pipeline {
	agent any 
	environment {
		CC = 'clang'
	}


	stages {
		stage('Example'){
			environment {
				DEBUG_FLAGS = '-g'
			}  //这里是在stage中定义变量
	
			steps {
				sh "${CC}${DEBUG_FLAGS}"
				sh 'printenv'
			}

		}

	}

}





