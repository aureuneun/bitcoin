package main

import (
	"github.com/aureuneun/bitcoin/cli"
	"github.com/aureuneun/bitcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
