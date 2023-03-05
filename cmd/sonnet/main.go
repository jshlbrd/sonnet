package main

import (
	"syscall/js"

	"github.com/google/go-jsonnet"
)

var vm = jsonnet.MakeVM()

func main() {
	js.Global().Get("window").Set("sonnet", js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			sonnet := args[0].String()

			jsonStr, err := vm.EvaluateAnonymousSnippet("", sonnet)
			if err != nil {
				return js.ValueOf(err.Error())
			}

			return js.ValueOf(jsonStr)
		},
	))

	select {}
}
