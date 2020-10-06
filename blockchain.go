package assignment01IBC

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type Block struct {
	transactions []string
	prevPointer  *Block
	prevHash     string
	currentHash  string
}

func CalculateHash(inputBlock *Block) string {
	//concatenate the entire blck in a string in order to calculate joint hash
	transList := strings.Join(inputBlock.transactions, ", ")
	entireBlock := fmt.Sprintf(transList, inputBlock.prevHash)
	//calculate hash
	ebHash := sha256.Sum256([]byte(entireBlock))

	fmt.Println("Hash calculated and stored")
	return fmt.Sprintf("%x", ebHash)
}

func InsertBlock(transactionsToInsert []string, chainHead *Block) *Block {
	//
	var newBlock *Block = new(Block)
	newBlock.transactions = transactionsToInsert

	if chainHead == nil {
		newBlock.prevHash = ""
		newBlock.prevPointer = nil
	} else {
		newBlock.prevHash = chainHead.currentHash
		newBlock.prevPointer = chainHead
	}

	newBlock.currentHash = CalculateHash(newBlock)

	//update chainhead
	chainHead = newBlock

	fmt.Println("New block inserted")
	return newBlock
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	currBlock := chainHead
	found := 0
	if chainHead != nil {
		for {
			for i, s := range currBlock.transactions {
				if s == oldTrans {
					currBlock.transactions[i] = newTrans
					found = 1
					fmt.Println("Replaced old transaction with new one!")
					break
				}
			}
			if found == 1 {
				break
			} else {
				currBlock = currBlock.prevPointer
			}
		}
	}

}

func ListBlocks(chainHead *Block) {

	currBlock := chainHead

	if chainHead != nil {
		for {
			blockTransactions := strings.Join(currBlock.transactions, ", ")
			fmt.Println(blockTransactions, currBlock.prevPointer, currBlock.prevHash, currBlock.currentHash)

			if currBlock.prevPointer == nil {
				break
			} else {
				currBlock = currBlock.prevPointer
				fmt.Println(">>>>>")
			}
		}
	}
}

func VerifyChain(chainHead *Block) {
	currBlock := chainHead
	if chainHead != nil {
		for {
			verHash := CalculateHash(currBlock.prevPointer)
			if currBlock.prevHash != verHash {
				fmt.Println("Uh-oh, blockchain has been compromised!")
				break
			} else {
				currBlock = currBlock.prevPointer
			}
			println("Blockchain verified successfully!")
		}
	}
}
