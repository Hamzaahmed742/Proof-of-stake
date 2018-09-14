package bc

import "hash"

type Block struct {
	hash, prevHash   hash.Hash32
	timeStamp        int64
	nrOfTransactions int
	data             []Transaction
}

func (b *Block) addTransact(t Transaction) {

}

func (b *Block) finalizeBlock() {

}
