pipeline {
    agent {
        kubernetes {
            yamlFile 'podTemplate.yaml'
        }
    } 
    // 定义环境变量
    environment {
        // 例如设置项目相关的变量
        PROJECT_NAME = "OceanWang"
        CONTAINER_NAME = "webhook_demo"
        HARBOR_URL = "oceanwang.harbor.domain"
        HARBOR_PROJECT = "library"
        GITHUB_REPO_URL = "https://github.com/Zy1bREAd/k8s-webhook-demo.git"
        // 采用argoCD
    }

    stages {
        stage('select-git'){
            steps {
                container('jnlp'){
                    script {
                        if (env.GIT_BRANCH) {
                            env.DOCKER_IMAGE_NAME = "webhook-easy"
                            env.DOCKER_IMAGE_TAG = "${env.GIT_BRANCH}"
                        } else if (env.GIT_TAG) {
                            env.DOCKER_IMAGE_NAME = "webhook-easy"
                            env.DOCKER_IMAGE_TAG = "${env.GIT_TAG}"
                        } else {
                            error "无法识别当前分支或标签"  // 终止流程
                        }
                    }
                }
            }
            
        }

        stage('docker-login') {
            steps {
                container('dind') {
                    withCredentials([usernamePassword(credentialsId: 'harbor_ci_robot', passwordVariable: 'harbor_robot_token', usernameVariable: 'harbor_robot_account')]) {
                        sh "echo ${harbor_robot_token}"
                        sh "echo ${harbor_robot_account}"
                        sh "docker login -u ${harbor_robot_account} -p ${harbor_robot_token} ${HARBOR_URL}"
                    }
                }
            }
        }

        stage('build-image') {
            when {
                anyOf {
                    expression { "${env.GIT_BRANCH}" =~ /(main)/ }
                    expression { "${env.GIT_TAG}" =~ /(^v.*)/ }
                }
            }
            steps {
                container('dind') {
                    sh "docker build -t ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG} -f Dockerfile --no-cache ."
                }
            }
        }

        stage('push-registry') {
            steps {
                container('dind') {
                    sh "docker tag ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG} ${HARBOR_URL}/${HARBOR_PROJECT}/${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}"
                    sh "docker push ${HARBOR_URL}/${HARBOR_PROJECT}/${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}"
                }
            }
        }
    }
}