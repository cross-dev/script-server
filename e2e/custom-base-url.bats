#!/usr/bin/env bats

load  ${BATS_TEST_DIRNAME}/lib.sh

setup() {
    echo '{{a}}' | $GOPATH/bin/script-server -b '/a/b/c/' - >/dev/null 2>&1 &
}

@test 'Custom base URL' {
    run curl -s 'http://localhost:41267/a/b/c/get?a=5'
    echo $output
    [ "$status" -eq "0" ]
    [ "$output" = "5" ]
}

