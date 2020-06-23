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

	post {
                failure {
                        updateGitlabCommitStatus name:'build', state:'faild'
                }
                success {
                        updateGitlabCommitStatus name:'build', state:'success'
                }
        }

        options {
                gitLabConnection('gitlab')
        }



}



