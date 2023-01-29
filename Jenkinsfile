pipeline{
    agent{
         node {
            label 'master'
            customWorkspace "workspace/${env.BRANCH_NAME}/src/github.com/aditya37/geospatial-tracking"
        }
    }
    environment {
        SERVICE  = "geospatial-tracking"
        NOTIFDEPLOY = -522638644
    }
    options {
        buildDiscarder(logRotator(daysToKeepStr: env.BRANCH_NAME == 'main' ? '90' : '30'))
    }
    stages{
        stage("Checkout"){
            when {
                anyOf { branch 'main'; branch 'develop'; branch 'staging'}
            }
            // Do clone
            steps {
                echo 'Checking out from git'
                checkout scm
                script {
                    env.GIT_COMMIT_MSG = sh (script: 'git log -1 --pretty=%B ${GIT_COMMIT}', returnStdout:true).trim()
                }
            }
        }
        stage('Build and deploy') {
            environment {
                GOPATH = "${env.JENKINS_HOME}/workspace/${env.BRANCH_NAME}"
                PATH = "${env.GOPATH}/bin:${env.PATH}"
            }
            stages {
                // build to dev
                stage('Deploy to env development') {
                    when {
                        branch 'develop'
                    }
                    environment {
                        NAMESPACE = 'geospatial-development'
			            TAG= '0.0.0'
                    }
                    steps {
                        // get credential file
                        withCredentials(
						[file(credentialsId: '8a0b10f8-0ccf-4d2d-9a01-c5881c57237f', variable: 'config')],
						[file(credentialsId: '43e011d2-1a2b-4182-8579-89bdf35c4270',variable:'firebase-sa')]
					) {
                            echo 'Build image'
                            sh "cp $config .env.geospatial.tracking"
                            sh "chmod 644 .env.geospatial.tracking"
					   sh "cp $firebase-sa sa.fbs.device.service.json"
					   sh 'chmod 644'
			             sh 'chmod +x build.sh'
			             sh './build.sh'
                            sh 'chmod +x deploy.sh'
                            sh './deploy.sh'
                            sh 'rm .env.geospatial.tracking'
					   sh 'sa.fbs.device.service.json'
                        }
                    }
                }
            }
        }
    }
    post{
        success{
            telegramSend(message:"Application $SERVICE has been [deployed] With Commit Message $GIT_COMMIT_MSG",chatId:"$NOTIFDEPLOY")
        }
        failure{
            telegramSend(message:"Application $SERVICE has been [Failed] With Commit Message $GIT_COMMIT_MSG",chatId:"$NOTIFDEPLOY")
        }
    }
}
