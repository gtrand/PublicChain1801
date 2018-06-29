package blc

import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
)

type ProofOfWork struct {
	block *Block
	target *big.Int
}

const targetbit = 4

func NewProofOfWork(block *Block) *ProofOfWork {
	newdata := big.NewInt(1)
	newdata = newdata.Lsh(newdata,256-targetbit)
	return &ProofOfWork{block,newdata}
}

func (pow *ProofOfWork)run() ([]byte,int64){
	nonce := 0
	var hash [32]byte
	var hashInt big.Int
	for {
		data := pow.prepareData(nonce)
		fmt.Println(data)
		// 生成hash
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x",hash)

		hashInt.SetBytes(hash[:])
		if pow.target.Cmp(&hashInt) == 1{
			break
		}
		nonce++


	}
	return hash[:],int64(nonce)
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{pow.block.Hash,
								pow.block.Data,
								pow.block.PrevBlockHash,
								Int2Byte(pow.block.Height),
								Int2Byte(pow.block.Timestamp),
								Int2Byte(pow.block.Nonce)},[]byte{})
	return data
}