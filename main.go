package main

import (
	"github.com/Guilherme415/cep-api/cmd"
	"github.com/Guilherme415/cep-api/internal"
)

func main() {
	internal.LoadDependencies()
	cmd.StartApi()
}
