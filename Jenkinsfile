node('mslave') {

  stage('prepare'){

    checkout([
      $class: 'GitSCM',
      branches: [[name: '*/master']],
      doGenerateSubmoduleConfigurations: false,
      extensions: [[$class: 'RelativeTargetDirectory', relativeTargetDir: 'hostvars-docs']],
      submoduleCfg: [],
      userRemoteConfigs: [[credentialsId: 'githubuser', url: 'https://github.com/tjtjtjtj/hostvars-docs.git']]
    ])

    // host-varsの取得
    sh "rsync -r --delete 192168.20.41:/home/jenkins/ansible/host_vars/ ./ansible-host_vars
    sh "rsync -r --delete 192168.20.41:/home/jenkins/serverspec/host_vars/ ./serverspec-host_vars

    withEnv(["URLPATH=\$(curl -s https://api.github.com/repos/tjtjtjtj/host-docs/releases | jq -r '.[0].assets[] | select(.name | test(\"host-docs\")) | .browser_download_url')"]) {
      sh "curl -sLO $URLPATH"
      sh "chmod 755 host-docs"
    }
  }

  stage('build'){
    // host-varsによるmarkdownの一覧更新
    sh "./host-docs --ansibledir ./ansible-host_vars/ --serverspecdir ./serverspec-host_vars/ --outputdir ./"
  }

  stage('update'){
    def BRANCH_NAME="update-${BUILD_NUMBER}"

    // 更新されたmarkdownの一覧をコピー
    sh "cp production.md hostvars-docs/production.md"
    sh "cp staging.md hostvars-docs/staging.md"
    sh "cp stress.md hostvars-docs/stress.md"

    // gitのリポジトリへpush
    dir('hostvars-docs') {
            
    // 各種設定
    sh """
      git config --local user.name tjtjtjtj
      git config --local user.email 'test@dummy.com'
      git config --local push.default simple
      git checkout -b ${BRANCH_NAME}
      git add .
      git commit -m 'Update docs'
    """

    withCredentials([usernamePassword(credentialsId: 'githubuser', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]) {

      // Remoteにpush
      sh "git push --set-upstream https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/tjtjtjtj/hostvars-docs.git ${BRANCH_NAME}"
           
      // PR作成
      sh """
        sudo docker run --rm \
        -v `pwd`:/opt/hub \
        -e GITHUB_HOST=github.com \
        -e GITHUB_USER=${GIT_USERNAME} \
        -e GITHUB_PASSWORD=${GIT_PASSWORD} \
        tianon/github-hub:latest \
        hub pull-request -m 'Update docs'
      """
    }
  }
}
