library 'LEAD'
pipeline {
  agent {
    label "lead-toolchain-skaffold"
  }
  stages {
    /// [build]
    stage('Build') {
      steps {
        notifyPipelineStart()
        notifyStageStart()
        container('skaffold') {
          sh "skaffold build --quiet > image.json"
        }
      }
      post {
        success {
          notifyStageEnd()
        }
        failure {
          notifyStageEnd([result: "Build Failed"])
        }
      }
    }
    /// [build]

    /// [stage]
    stage("Deploy to Staging") {
      when {
          branch 'master'
      }
      environment {
        TILLER_NAMESPACE = "${env.stagingNamespace}"
        ISTIO_DOMAIN   = "${env.stagingDomain}"
      }
      steps {
        notifyStageStart()
        container('skaffold') {
          sh "skaffold deploy -a image.json -n ${TILLER_NAMESPACE}"
        }
      }
      post {
        success {
          notifyStageEnd([status: "Successfully deployed to staging:\nspringtrader.${env.stagingDomain}/spring-nanotrader-web/"])
        }
        failure {
          notifyStageEnd([result: "Deploy failed"])
        }
      }
    }
    /// [stage]

    /// [gate]
    stage ('Manual Ready Check') {
      options {
        timeout(time: 30, unit: 'MINUTES')
      }
      when {
        branch 'master'
      }
      options {
        timeout(time: 30, unit: 'MINUTES')
      }
      input {
        message 'Deploy to Production?'
      }
      steps {
        echo "Deploying"
      }
    }
    /// [gate]

    /// [prod]
    stage("Deploy to Production") {
      when {
          branch 'master'
      }
      environment {
        TILLER_NAMESPACE = "${env.productionNamespace}"
        ISTIO_DOMAIN   = "${env.productionDomain}"
      }
      steps {
        notifyStageStart()
        container('skaffold') {
          sh "skaffold deploy -a image.json -n ${TILLER_NAMESPACE}"
        }
      }
      post {
        success {
          notifyStageEnd([status: "Successfully deployed to production:\nspringtrader.${env.productionNamespace}/spring-nanotrader-web/"])
        }
        failure {
          notifyStageEnd([result: "fail"])
        }
      }
    }
    /// [prod]
  }
  post {
    success {
      echo "Pipeline Success"
      notifyPipelineEnd()
    }
    failure {
      echo "Pipeline Fail"
      notifyPipelineEnd([result: "fail"])
    }
  }
}
