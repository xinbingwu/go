pipeline {
	agent any
	triggers {
		gitlab(triggerOnPush:true,
			triggerOnMergeRequest:true,
			branchFilterType:'All',
			secretToken:"abcde")
	}

	stages {
		stage('build'){
			steps{
				echo "hello world from gitlab trigger"
			}
		}
	}

}



