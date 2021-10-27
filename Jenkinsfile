pipeline {
    agent any 
    tools {
        go '1.17.2'
    }
    
    stages {
        stage('Clone Repo') {
            steps {
                git branch: 'main',
                url: 'https://github.com/ankitchauhan123/chat.git'
            }
        }
         stage('Build') {
            steps {
                echo 'Compiling and Building'
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
