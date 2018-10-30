node('docker') {
    stage('Get Latest Code') {
        cleanWs()
        checkout scm
    }

    stage('Multi-Stage Build Go Code') {
        docker.build("modern-jenkins/go-sample-multi-stage:${env.BUILD_NUMBER}", "-f Dockerfile.multi-stage-build --build-arg GIT_COMMIT=${env.GIT_COMMIT} .")
    }

    stage('Jenkins Based Build Go Code') {
        docker.image('golang:1.11-alpine').inside {
            sh 'mkdir -p /usr/src/app'
            sh 'cp $WORKSPACE/* /usr/src/app/'
            sh 'cd /usr/src/app && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $WORKSPACE/go-sample .'
        }

        docker.build("modern-jenkins/go-sample-scratch:${env.BUILD_NUMBER}", "--build-arg GIT_COMMIT=${env.GIT_COMMIT} .")
        
    }
}