#!/usr/bin/env bats

load  ${BATS_TEST_DIRNAME}/lib.sh

setup() {
    local input=$(mktemp $BATS_TMPDIR/XXXXXXXXX)
    echo '{{a}}-{{b}}-{{c | printf "%q"}}' >$input
    $GOPATH/bin/script-server $input >/dev/null 2>&1 &
}

@test 'Template from file' {
    run curl -s 'http://localhost:41267/get?a=5&b=ergo&c=up'
    [ "$status" -eq "0" ]
    [ "$output" = '5-ergo-"up"' ]
}

