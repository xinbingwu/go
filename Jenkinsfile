pipeline {
	agent any 
	environment {
		CC = 'clang'
	}


	stages {
		stage('Example'){
			environment {
				DEBUG_FLAGS = '-g'
			}
		}
		stage {
			sh "${CC}${DEBUG_FLAGS}"
			sh 'printenv'
		}

	}

}





