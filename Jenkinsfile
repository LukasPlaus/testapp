pipeline {
    agent any
    tools {
        go 'go-1.11'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Compile') {
            steps {
                sh 'go build'
            }
        }
        stage('Test') {
            environment {
                CODECOV_TOKEN = credentials('codecov_token')
            }
            steps {
                sh 'go test ./... -coverprofile=coverage.txt'
                sh "curl -s https://codecov.io/bash | bash -s -"
            }
        }
        stage('Code Analysis') {
            steps {
                sh 'curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $GOPATH/bin v1.12.5'
                sh 'golangci-lint run'
            }
        }
        stage('Build Docker Image') {
            when {
                branch 'master'
            }
            steps {
                echo '=== Building Docker Image ==='
                script {
                    app = docker.build("lukasplaus/go-app-test")
                }
            }
        }
                stage('Push Docker Image') {
            when {
                branch 'master'
            }
            steps {
                echo '=== Pushing Docker Image ==='
                script {
                    GIT_COMMIT_HASH = sh (script: "git log -n 1 --pretty=format:'%H'", returnStdout: true)
                    SHORT_COMMIT = "${GIT_COMMIT_HASH[0..7]}"
                    docker.withRegistry('https://registry.hub.docker.com', 'dockerHubCredentials') {
                        app.push("$SHORT_COMMIT")
                        app.push("latest")
                        
                    }
                }
            }
        }
        stage('Remove local images') {
            steps {
                echo '=== Delete the local docker images ==='
                sh("docker rmi -f lukasplaus/go-app-test:latest || :")
                sh("docker rmi -f lukasplaus/go-app-test:$SHORT_COMMIT || :")
            }
        }
    }
}
