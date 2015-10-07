#!/usr/bin/env bats

load  ${BATS_TEST_DIRNAME}/lib.sh

setup() {
    echo '{{a}}-{{b}}-{{c}}' | $GOPATH/bin/script-server - >/dev/null 2>&1 &
}

@test 'Port is taken' {
    run $GOPATH/bin/script-server <(echo '{{a}}-{{b}}-{{c}}')
    [ "$status" -ne "0" ]
}

