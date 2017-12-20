pipeline {
  agent any
  stages {
    stage('build') {
      parallel {
        stage('build') {
          steps {
            sh 'echo "===================="'
            echo 'hello'
          }
        }
        stage('') {
          steps {
            echo 'hello'
          }
        }
      }
    }
  }
}