package blockchain

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type CLI struct {
	Blockchain *BlockChain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  printchain - Print all the blocks of the blockchain")
	fmt.Println("  addblock - with the data flag you must specify the data of the new block that you want to add Ex: addblock -data \"Some data\"")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChain := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		Validate(addBlockCmd.Parse(os.Args[2:]))
	case "printchain":
		Validate(printChain.Parse(os.Args[2:]))
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChain.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) addBlock(data string) {
	cli.Blockchain.AddBlock(data)
	fmt.Println("[+] Success!")
}

func (cli *CLI) printChain() {
	iter := cli.Blockchain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Prev. Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		proofOfWork := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(proofOfWork.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}
