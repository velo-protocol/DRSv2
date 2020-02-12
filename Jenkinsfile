pipeline {
    agent { label "slave" }
    parameters {
        choice(name: 'branches', choices: ['none', 'feature/pipeline-fix', 'develop', 'master'], description: 'Passing the Branches')
        string(name: 'scc_host',
                defaultValue: 'none',
                description: 'Host for SCC without http://'
        )
        string(name: 'scc_pk',
                defaultValue: 'none',
                description: 'Private Key for SCC'
        )
    }
    environment {
        APPENV = "${params.branches}"
        APPHOST = "${params.scc_host}"
        APPPK = "${params.scc_pk}"
    }

    stages {
        stage('Pipeline CI') {
            when {
                expression { params.branches ==~ /(none)/ }
                anyOf {
                    branch 'master'
                    branch 'develop'
                    branch 'feature/pipeline-fix'
                }
            }
            stages {
                stage('CI-Cleanup') {
                    steps {
                        dir('directoryToDelete') {
                            deleteDir()
                        }
                    }
                }
                stage('CI-Git Clone') {
                    steps {
                        git(
                                url: 'git@github.com:velo-protocol/DRSv2.git',
                                credentialsId: '9d628039-dc04-498a-b6d8-3daa15c87068',
                                branch: "${env.GIT_BRANCH}"
                        )
                    }
                }
                stage('CI-Install Truffle') {
                    steps {
                        sh '''
                                echo "install truffle v_5.1.4"
                                yarn global add truffle@5.1.4
                        '''
                    }
                }
                stage('CI-Install Packages') {
                    steps {
                        sh '''
                                echo "install all dependencies with yarn"
                                yarn install
                        '''
                    }
                }
                stage('CI-Run Test') {
                    steps {
                        sh '''
                                echo "run test with yarn"
                                yarn run test:cov
                        '''
                    }
                }
            }
        }

        stage('Pipeline CD') {
            when {
                expression { params.branches ==~ /(feature\/pipeline-fix|develop|master)/ }
            }
            stages {
                stage('CD-Cleanup') {
                    steps {
                        dir('directoryToDelete') {
                            deleteDir()
                        }
                    }
                }
                stage('CD-Git Clone') {
                    steps {
                        git(
                                url: 'git@github.com:velo-protocol/DRSv2.git',
                                credentialsId: '9d628039-dc04-498a-b6d8-3daa15c87068',
                                branch: "${APPENV}"
                        )
                    }
                }
                stage('CD-Rename Env file') {
                    steps {
                        sh """
                                echo "Rename .env.example file"
                                mv .env.example .env
                                cat .env
                        """
                    }
                }
                stage('CD-Create Env File') {
                    steps {
                        sh """
                                echo "Replace .env file with variable"
                                sed -i "s/#APPHOST#/http:\\/\\/${APPHOST}/g" .env
                                sed -i "s/#APPPK#/${APPPK}/g" .env
                                cat .env
                        """
                    }
                }
                stage('CD-Install Truffle') {
                    steps {
                        sh '''
                                echo "install truffle v_5.1.4"
                                yarn global add truffle@5.1.4
                        '''
                    }
                }
                stage('CD-Install Packages') {
                    steps {
                        sh '''
                                echo "install all dependencies with yarn"
                                yarn install
                        '''

                    }
                }
                stage('CD-Run Test') {
                    steps {
                        sh '''
                                echo "run test with yarn"
                                yarn run reset:dev
                        '''
                    }
                }
            }
        }
    }

    post {
        cleanup {
            deleteDir()
        }
    }
}