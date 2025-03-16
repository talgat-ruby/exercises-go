package main

import (
    "strconv"
)

func binary(decimalStr int) string {

    return strconv.FormatInt(int64(decimalStr), 2)

}
