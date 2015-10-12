#!/usr/bin/env bats

load  ${BATS_TEST_DIRNAME}/lib.sh

setup() {
    echo '{{a}}-{{b}}-{{c}}' | script-server - >/dev/null 2>&1 &
}

@test 'Extraneous parameters' {
    run curl -s 'http://localhost:41267/get?a=5&b=ergo&c=344&d=xxx'
    [ "$status" -eq "0" ]
    echo $output
    [ "$output" = "5-ergo-344" ]
}

