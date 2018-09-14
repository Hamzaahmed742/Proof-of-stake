package bc

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/binary"
	// "fmt"
	"golang.org/x/crypto/sha3"
	// "hash"
	"crypto/elliptic"
	"math/big"
	"reflect"
)

type Transaction struct {
	Hash [32]byte
	Sig  [64]byte
	Info TxInfo
}

type TxInfo struct {
	Nonce, Amount int64
	From, To      [64]byte
}

func ConstrTx(nonce, amount int64, from, to Account, key *ecdsa.PrivateKey) (tx Transaction, err error) {
	if amount > from.Balance && amount > 0 {
		return
	}
	if reflect.DeepEqual(from, to) {
		return
	}
	if nonce != from.Nonce {
		return
	}

	serialized := encodeTxContent(nonce, amount, from.Id, to.Id)
	tx.Hash = sha3.Sum256(serialized)

	r, s, err := ecdsa.Sign(rand.Reader, key, tx.Hash[:])

	copy(tx.Sig[:32], r.Bytes())
	copy(tx.Sig[32:], s.Bytes())

	tx.Info.From = from.Id
	tx.Info.To = to.Id
	tx.Info.Amount = amount

	return
}

func encodeTxContent(nonce, amount int64, from, to [64]byte) (enc []byte) {
	var buf bytes.Buffer

	hash := TxInfo{nonce, amount, from, to}
	binary.Write(&buf, binary.BigEndian, hash)
	return buf.Bytes()
}

func (tx *Transaction) VerifyTx() bool {
	pub1, pub2 := new(big.Int), new(big.Int)
	r, s := new(big.Int), new(big.Int)

	pub1.SetBytes(tx.Info.From[:32])
	pub2.SetBytes(tx.Info.From[32:])

	pubKey := ecdsa.PublicKey{elliptic.P256(), pub1, pub2}
	r.SetBytes(tx.Sig[:32])
	s.SetBytes(tx.Sig[32:])

	correct := ecdsa.Verify(&pubKey, tx.Hash[:], r, s)

	return correct
}
