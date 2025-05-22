package main

import (
	"github.com/ezzycreative1/toko-kue-system/cmd"
	"github.com/ezzycreative1/toko-kue-system/deps"
)

func main() {
	dependency := deps.SetupDependencies()
	cmd.ExecuteHTTP(dependency)
}
