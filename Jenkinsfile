pipeline {
	agent any 
	triggers {
		cron("H/1 * * * *")
	}

	stages {
		stage('Nighty build'){
			steps {
				echo '这是一个耗时的构建，每天凌晨执行。'
			}
		}
	}

}




