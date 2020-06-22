pipeline {
	agent none
	
	stages {
		stage('stash-保存文件'){
			agent { label "master" }
			steps {
				writeFile(file: "a.txt",text: "$BUILD_NUMBER")
				stash(name: "abc", includes: "a.txt")
			}
		}

		stage('unstash-提取文件'){
			agent { label "node2" }
			steps{
				script {
					unstash("abc")
					def content = readFile("a.txt")
					echo "${content}"
				}
			}
		}
