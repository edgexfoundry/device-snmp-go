//
// Copyright (c) 2019 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

loadGlobalLibrary()

pipeline {
    agent {
        label 'centos7-docker-2c-1g'
    }
    options {
        timestamps()
    }
    stages {
        stage('Prepare') {
            steps {
                edgeXSetupEnvironment()
                edgeXDockerLogin(settingsFile: env.MVN_SETTINGS)
                edgeXSemver 'init'
                script {
                    def semverVersion = edgeXSemver()
                    env.setProperty('VERSION', semverVersion)
                    sh 'echo $VERSION > VERSION'
                    stash name: 'semver', includes: '.semver/**,VERSION', useDefaultExcludes: false
                }
            }
        }
        stage('Docker Build') {
            parallel {
                stage('build-amd64') {
                    agent {
                        label 'centos7-docker-4c-2g'
                    }
                    steps {
                        doBuild 'amd64'
                    }
                }
                stage('build-arm64') {
                    agent {
                        label 'ubuntu18.04-docker-arm64-4c-2g'
                    }
                    steps {
                        doBuild 'arm64'
                    }
                }
            }
        }
        stage('SemVer Tag') {
            when {
                expression {
                    edgex.isReleaseStream()
                }
            }
            steps {
                unstash 'semver'
                sh 'echo v${VERSION}'
                edgeXSemver('tag')
                edgeXInfraLFToolsSign(command: 'git-tag', version: 'v${VERSION}')
                edgeXSemver('push')
            }
        }
        stage('Docker Push') {
            when {
                expression {
                    edgex.isReleaseStream()
                }
            }
            steps {
                doPush 'amd64'
                doPush 'arm64'
            }
        }
        stage('SemVer Bump') {
            when {
                expression {
                    edgex.isReleaseStream()
                }
            }
            steps {
                unstash 'semver'
                edgeXSemver('bump pre')
                edgeXSemver('push')
            }
        }
    }

    post {
        always {
            edgeXInfraPublish()
        }
    }
}

def doBuild(arch, image='edgexfoundry/device-snmp-go') {
    unstash 'semver'
    dockerBuild(arch, image)
    dockerSave(arch, image)
}

def doPush(arch, image='edgexfoundry/device-snmp-go') {
    unstash 'semver'
    dockerLoad(arch, image)
    dockerPush(arch, image, "https://${env.DOCKER_REGISTRY}:10004")
    // dockerPush(arch, image, "https://${env.DOCKERHUB_REGISTRY}")
}

def dockerBuild(arch, image) {
    sh "docker build --build-arg 'MAKE=build test' --tag ${image}:${arch} ."
}

def dockerSave(arch, image) {
    sh "docker image save --output ${arch}.tar ${image}:${arch}"
    sh "docker system info --format {{.Architecture}} > ${arch}.txt"
    stash name: arch, includes: "${arch}.tar,${arch}.txt"
}

def dockerLoad(arch, image) {
    unstash "${arch}"
    sh "docker image load --input ${arch}.tar"
}

def dockerPush(arch, image, registry) {
    sh 'env | sort'
    def img = docker.image("${image}:${arch}")
    def mach = readFile("${arch}.txt").trim()
    def versions = [ env.VERSION, 'latest' ]
    for (ver in versions) {
        docker.withRegistry(registry, 'device-snmp') {
            img.push("${ver}-${arch}")
            img.push("${ver}-${mach}")
        }
    }
}

def loadGlobalLibrary(branch = '*/master') {
    library(identifier: 'edgex-global-pipelines@master',
        retriever: legacySCM([
            $class: 'GitSCM',
            userRemoteConfigs: [[url: 'https://github.com/edgexfoundry/edgex-global-pipelines.git']],
            branches: [[name: branch]],
            doGenerateSubmoduleConfigurations: false,
            extensions: [[
                $class: 'SubmoduleOption',
                recursiveSubmodules: true,
            ]]]
        )
    ) _
}