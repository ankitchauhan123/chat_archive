pipeline {
    agent any 
    tools {
        go '1.17.2'
    }
    
    stages {
         stage('Build') {
            steps {
                echo 'Compiling and Building...'
                sh 'go build -o ./build/chat_build ./cmd'
            }
        }
        
    }
    post {
        success {
            archiveArtifacts artifacts: 'build/*'

        }
    }
    
}
