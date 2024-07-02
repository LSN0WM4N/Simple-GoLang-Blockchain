package main

import (
	BlockChain "main/BlockChain"
)

func main() {
	bc := BlockChain.NewBlockChain()
	defer bc.Db.Close()

	cli := &BlockChain.CLI{Blockchain: bc}
	cli.Run()
}
