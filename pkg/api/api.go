package api

import (
	"fmt"

	constants "github.com/dk1242/temp_mod/constants"
)

func CallAPI() {
	fmt.Println("API called")
	fmt.Println("Log level ", constants.Log_level)
}
