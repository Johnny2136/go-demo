pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'go build -o my-api'
            }
        }
        stage('Docker Build') {
            steps {
                sh 'docker build -t my-api-image .'
            }
        }
        stage('Docker Run') {
            steps {
                sh 'docker run -d --name my-api-container -p 8080:8080 my-api-image'
            }
        }
    }
}
