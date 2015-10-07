#!/usr/bin/env bats

load  ${BATS_TEST_DIRNAME}/lib.sh

setup() {
    echo '{{a}}-{{b}}-{{c}}' | $GOPATH/bin/script-server -l ':17822' - >/dev/null 2>&1 &
}

@test 'Listen on custom port' {
    run curl -s 'http://localhost:17822/get?a=5&b=ergo&c=344'
    [ "$status" -eq "0" ]
    echo $output
    [ "$output" = "5-ergo-344" ]
}

