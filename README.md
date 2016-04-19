# go-sql-test-harness
Confirm connections are properly closed during testing

Requires Go >= 1.4 because of `TestMain`

Set credentials for a postgres DB in `harness_test.setup`. Compare the output of:

    go test -run=TestGood

to:

    go test -run=TestBad


Happy hacking!

aodin, 2016
