#!/usr/bin/env bats

load  ${BATS_TEST_DIRNAME}/lib.sh

setup() {
    echo '{{a}}-{{b}}-{{c}}' | $GOPATH/bin/script-server - >/dev/null 2>&1 &
}

@test 'Missing parameters' {
    run curl -s 'http://localhost:41267/get?a=5&b=ergo'
    [ "$status" -eq "0" ]
    echo $output
    [ "$output" != "" ]
    run curl --fail 'http://localhost:41267/get?a=5&b=ergo'
    [ "$status" -ne "0" ]
}

