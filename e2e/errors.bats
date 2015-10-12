#!/usr/bin/env bats

@test 'Non-existent file passed' {
    local input=$(mktemp $BATS_TMPDIR/XXXXXXX)
    rm $input
    run script-server $input
    [ "$status" -ne "0" ]
}

@test 'Unreadable file passed' {
    local input=$(mktemp $BATS_TMPDIR/XXXXXXX)
    chmod u-r $input
    run script-server $input
    [ "$status" -ne "0" ]
}

@test 'No file passed' {
    run script-server
    [ "$status" -ne "0" ]
}

@test 'Relative base URL' {
    run script-server -b a/b/c <(echo '{{a}}')
    [ "$status" -ne "0" ]
}

@test 'Protected port passed' {
    run script-server -l ':84' <(echo '{{a}}')
    [ "$status" -ne "0" ]
}
