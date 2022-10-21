package main

import (
	"github.com/dankChain/dank-faucet/cmd"
)

//go:generate npm run build
func main() {
	cmd.Execute()
}
