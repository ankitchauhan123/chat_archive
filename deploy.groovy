job('go-chat'){
    description ("Chat build and deployment job ")
    scm{
        github('ankitchauhan123/chat','main')
    }
    triggers{
        scm('* * * * *')
    }
    steps{
        shell('go build -o ./build/chat_build ./cmd')
    }
    publishers {
            archiveArtifacts('build/*')
    }

}
