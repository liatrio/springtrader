folder('liatrio') {
}
multibranchPipelineJob('liatrio/springtrader') {
  factory {
    workflowBranchProjectFactory {
      scriptPath('.cache/Jenkinsfile')
    }
  }
  triggers {
    periodic(10)
  }
  branchSources {
    git {
      id('http://github.com/liatrio/springtrader.git')
      remote('https://github.com/liatrio/springtrader.git')
      includes('solution1')
    }
  }
  orphanedItemStrategy {
    discardOldItems {
      numToKeep(20)
    }
  }
}
