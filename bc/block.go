package bc

import "hash"

type Block struct {
	Hash, PrevHash   hash.Hash32
	TimeStamp        int64
	NrOfTransactions int
	data             []Transaction
	Version          uint8
	StateCopy        map[[64]byte]int64
}

func (b *Block) AddTx(tx *Transaction) {
	if !tx.VerifyTx() || tx.Info.Amount > b.StateCopy[tx.Info.From] {
		return
	}

	b.StateCopy[tx.Info.From] -= tx.Info.Amount
	b.StateCopy[tx.Info.To] += tx.Info.Amount
	b.NrOfTransactions++
}

func (b *Block) FinalizeBlock() {

}
