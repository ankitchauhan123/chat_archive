job('go-chat'){
    description ("Chat build and deployment job ")
    wrappers {
        golang('Go 1.17.2')
    }
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
