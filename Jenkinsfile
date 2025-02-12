pipeline {
    agent { kubernetes { label 'go' } }

    // 定义环境变量
    environment {
        // 例如设置项目相关的变量
        PROJECT_NAME = "OceanWang"
        CONTAINER_NAME = "webhook_demo"
        HARBOR_URL = "oceanwang.harbor.domain"
        HARBOR_PROJECT = "library"
        GITHUB_REPO_URL = "https://github.com/Zy1bREAd/k8s-webhook-demo.git"
        // 采用argoCD
        // DEVELOP_SERVER_IP = "10.0.20.5"
        // DEVELOP_SERVER_USER = "ubuntu"
        // DEVELOP_SERVER_CRED_ID = "ssh-for-password-10.0.20.5"
    }

    // 构建步骤
    stages {
        stage('Checkout GitHub Branch and Pull Code') {
            steps {
                script {
                    if (env.GIT_BRANCH){
                        DOCKER_IMAGE_NAME = "webhook-easy"
                        DOCKER_IMAGE_TAG = "${env.GIT_BRANCH}"
                    }else if (env.GIT_TAG){
                        DOCKER_IMAGE_NAME = "webhook-easy"
                        DOCKER_IMAGE_TAG = "${env.GIT_TAG}"
                    }else {
                        echo "无法识别当前分支或标签"
                    }
                }
            }
        }
        // 登录Harbor镜像仓库
        stage('Login Image Registry') {
            steps {
                container('dind') {
                    withCredentials([usernamePassword(credentialsId: 'harbor_ci_robot', passwordVariable: 'harbor_robot_token', usernameVariable: 'harbor_robot_account')]) {
                        sh "sudo docker login ${HARBOR_URL} -u ${harbor_robot_account} -p ${harbor_robot_token}"
                    }
                }

            }
        }
        // 构建镜像在dev环境
        stage('Build On Image For Develop') {
            when {
                anyOf {
                    expression { "${env.GIT_BRANCH}" =~ /(main)/ }
                    expression { "${env.GIT_TAG}" =~ /(^v.*)/ }
                }
            }
            steps {
                container('dind') {
                    sh "sudo docker build -t ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG} -f Dockerfile --no-cache ."
                }
            }
        }

        // 构建镜像在prod环境
        stage('Build On Image For Production') {
            when {
                anyOf {
                    expression { "${env.GIT_BRANCH}" =~ /(^release-v.*)/ }
                    expression { "${env.GIT_TAG}" =~ /(^release-v.*)/ }
                }
            }
            steps {
                sh "echo 'Enter Prod Pipeline for Build image.'"
            }
        }
        stage('Push Image') {
            // 推送镜像到Harbor
            steps {
                container('dind') {
                    sh "sudo docker tag ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG} ${HARBOR_URL}/${HARBOR_PROJECT}/${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}"
                    sh "sudo docker push ${HARBOR_URL}/${HARBOR_PROJECT}/${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}"
                }
            }
        }
        // stage('Deploy To Develop Env') {
        //     when {
        //         anyOf {
        //             expression { "${env.GIT_BRANCH}" =~ /(main)/ }
        //             expression { "${env.GIT_TAG}" =~ /(^v.*)/ }
        //         }
        //     }
        //     steps {
        //         script {
        //             def remote = [:]
        //             remote.name = 'develop-server-01'
        //             remote.host = "${DEVELOP_SERVER_IP}"
        //             remote.allowAnyHosts = true
        //             withCredentials([usernamePassword(credentialsId: 'harbor_robot_account', passwordVariable: 'harbor_robot_token', usernameVariable: 'harbor_robot_account'), usernamePassword(credentialsId: 'ssh-for-password-10.0.20.5', passwordVariable: 'dev_server_pwd', usernameVariable: 'dev_server_user')]) {
        //                 // 设置ssh server的login info
        //                 remote.user = "${dev_server_user}"
        //                 remote.password = "${dev_server_pwd}"
        //                 // 登录Harbor
        //                 sshCommand remote: remote, command: "sudo docker login ${HARBOR_URL} -u ${harbor_robot_account} -p ${harbor_robot_token}"
        //                 // 停止并删除之前的容器
        //                 sshCommand remote: remote, command: "if [ -n \"\$(sudo docker ps -a -q --filter name=${CONTAINER_NAME})\" ];then sudo docker stop ${CONTAINER_NAME} && sudo docker rm ${CONTAINER_NAME};else echo 'Container is not exist';fi"
        //                 sshCommand remote: remote, command: "sudo docker pull ${HARBOR_URL}/${HARBOR_PROJECT}/${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}"
        //                 sshCommand remote: remote, command: "sudo docker run -itd -p 17443:17443 --name=${CONTAINER_NAME} ${HARBOR_URL}/${HARBOR_PROJECT}/${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}"
        //             }
        //         }
        //     }
        // }
    }
}