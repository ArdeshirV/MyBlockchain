package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println(Prompt("My Blockchain"))
}

type Block struct {
  Index int
  TimeStamp string
  Data string
  Hash string
  PrevHash string
}

type Blockchain struct {
  Chain []Block
}

func CalculateHash(block Block) string {
  record := fmt.Sprintf("%d%s%s%s", block.Index, block.TimeStamp,
    block.Data, block.PrevHash)
  h := sha256.New()
  h.Write([]byte(record))
  hashed := h.Sum(nil)
  return hex.EncodeToString(hashed)
}
