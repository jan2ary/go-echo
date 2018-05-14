node {
    stage('Checkout the codebase') {
        git url: 'git@github.com:jan2ary/go-echo.git'
    }
    stage ('Build') {
        def customImage = docker.build("jan3ary/rot13:${env.BUILD_ID}")
    }
    stage ('Test') {
        docker.image("jan3ary/rot13:${env.BUILD_ID}").withRun('-p 9000') { c ->
            def (host, port) = c.port(9000).split(':')
            def test = sh (returnStdout: true, script: "echo '0123456789XYZABC!@#$%^&*()_+' | nc -N ${host} ${port}")
            echo "Test has returned a value: $test"
            if (test != '=>?@ABCDEFklmnop.M012k3756l8') {
                currentBuild.result = 'FAILURE'
            }
        }
    }
}
