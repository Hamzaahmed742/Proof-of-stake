package bc

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/binary"
	// "fmt"
	"golang.org/x/crypto/sha3"
	// "hash"
	"math/big"
	"reflect"
)

type Transaction struct {
	Hash     [32]byte
	R, S     *big.Int
	From, To Account
	Amount   int64
}

type HashBuidlingBlocks struct {
	Nonce, Amount int64
	From, To      [64]byte
}

func ConstrTransaction(nonce, amount int64, from, to Account, key *ecdsa.PrivateKey) (tx Transaction, err error) {
	if amount > from.Balance {
		return
	}
	if reflect.DeepEqual(from, to) {
		return
	}
	if nonce != from.Nonce {
		return
	}

	serialized := encodeTransactContent(nonce, amount, from.Id, to.Id)
	tx.R, tx.S, err = ecdsa.Sign(rand.Reader, key, serialized)

	tx.From = from
	tx.To = to
	tx.Amount = amount
	tx.Hash = sha3.Sum256(serialized)

	return
}

func encodeTransactContent(nonce, amount int64, from, to [64]byte) (enc []byte) {
	var buff bytes.Buffer

	hash := HashBuidlingBlocks{nonce, amount, from, to}
	binary.Write(&buff, binary.LittleEndian, hash)
	return buff.Bytes()
}
