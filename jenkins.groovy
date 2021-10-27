job('go-chat'){
    description ("Chat build and deployment job ")
    scm{
        git("https://github.com/ankitchauhan123/chat.git",main)
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
