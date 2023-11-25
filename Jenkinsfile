pipeline {
    agent any
    environment {
        repo = 'alhazmy20/go-webapi'
        repo_ver = 'latest'
    }

    stages {
        stage('Respositry Cloning') {
            steps {
                checkout scm
            }
        }

        stage('Testing API') {
            steps {
                sh 'go mod download'
                sh 'go mod verify'
                sh 'go mod tidy'
                sh 'go test ./...'
            }
        }

        stage('Building Image') {
            steps {
                script {
                    docker.build("${repo}", '--build-arg APP_ENV=production DB_CONTAINER_NAME=dockerDB')
                }
            }
        }

        stage('Pushing Image') {
            steps {
                script {
                    docker.withRegistry('', 'docker-credentials') {
                        docker.image("${repo}:${repo_ver}").push()
                    }
                }
            }
        }

        stage('Cleaning') {
            steps {
                script {
                    sh "docker rmi ${repo}"
                }
            }
        }
    }
}
