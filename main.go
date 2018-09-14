package main

import (
	"blockchain/bc"
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/gob"
	"fmt"
)

func main() {
	state := make(map[[64]byte]int64)
	privA, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	privB, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		return
	}

	accA := bc.Account{Nonce: 0, Balance: 15}
	copy(accA.Id[0:32], privA.PublicKey.X.Bytes())
	copy(accA.Id[32:64], privA.PublicKey.Y.Bytes())

	accB := bc.Account{Nonce: 0, Balance: 12}
	copy(accB.Id[0:32], privB.PublicKey.X.Bytes())
	copy(accB.Id[32:64], privB.PublicKey.Y.Bytes())

	state[accA.Id] = accA.Balance
	state[accB.Id] = accB.Balance

	tx, err := bc.ConstrTx(0, 2, accA, accB, privA)

	// b := bc.Block{StateCopy: state}

	var buf bytes.Buffer
	var rcvTx bc.Transaction
	enc := gob.NewEncoder(&buf)
	enc.Encode(tx)
	fmt.Printf("%x\n", buf.Bytes())
	dec := gob.NewDecoder(&buf)
	dec.Decode(&rcvTx)

	fmt.Printf("%t\n", tx.VerifyTx())

}
