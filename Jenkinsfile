pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[credentialsId: '6557e880-4df2-4e5f-a28f-1e13d07529f5', url: 'git@github.com:jan2ary/go-echo.git']]])
            }
        }
        stage ('Build') {
            steps {
                sh 'docker build -f "Dockerfile" -t jan3ary/echo-go:latest .'
            }
        }
    }
}
