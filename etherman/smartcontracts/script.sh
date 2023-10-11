#!/bin/sh

set -e

gen() {
    local package=$1

    abigen --bin bin/${package}.bin --abi abi/${package}.abi --pkg=${package} --out=${package}/${package}.go
}

gen xagonzkevm
gen xagonzkevmbridge
gen matic
gen xagonzkevmglobalexitroot
gen mockverifier