package main

import (
	"fmt"

	"github.com/dk1242/temp_mod/pkg/api"
	"rsc.io/quote/v4"
)

func main() {
	fmt.Println(quote.Go())
	api.CallAPI()
}
