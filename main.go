package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(Prompt("My Blockchain"))
  var bc Blockchain
  _, err := os.Stat("blockchain.json")
  if err != nil && os.IsNotExist(err) {
    bc = *NewBlockchain()
    err := SaveBlockchain(bc)
    if err != nil {
      log.Fatal(err)
    }
  } else {
    bc, err = LoadBlockchain()
    if err != nil {
      log.Fatal(err)
    }
  }

  if os.IsNotExist(err) {
    bc.AddBlock("Block 1 Data")
    err = SaveBlockchain(bc)
    if err != nil {
      log.Fatal(err)
    }

    bc.AddBlock("Block 2 Data")
    err = SaveBlockchain(bc)
    if err != nil {
      log.Fatal(err)
    }

    bc.AddBlock("Block 3 Data")
    err = SaveBlockchain(bc)
    if err != nil {
      log.Fatal(err)
    }
  }

  for _, block := range bc.Chain {
    fmt.Printf("Index: %d\n", block.Index)
    fmt.Printf("Timestamp: %s\n", block.TimeStamp)
    fmt.Printf("Data: %s\n", block.Data)
    fmt.Printf("Hash: %s\n", block.Hash)
    fmt.Printf("Previous Hash: %s\n", block.PrevHash)
    fmt.Println(strings.Repeat("-", 40))
  }
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

func GenerateBlock(prevHash Block, data string) Block {
  var block Block
  block.Index = prevHash.Index + 1
  block.TimeStamp = time.Now().String()
  block.Data = data
  block.PrevHash = prevHash.Hash
  block.Hash = CalculateHash(block)
  return block
}

func (bc *Blockchain) AddBlock(data string) {
  prevBlock := bc.Chain[len(bc.Chain)-1]
  newBlock := GenerateBlock(prevBlock, data)
  bc.Chain = append(bc.Chain, newBlock)
}

func NewBlockchain() *Blockchain {
  genesisBlock := Block{0, time.Now().String(), "Genesis Block", "", ""}
  genesisBlock.Hash = CalculateHash(genesisBlock)
  return &Blockchain{[]Block{genesisBlock}}
}

func SaveBlockchain(bc Blockchain) error {
  data, err := json.Marshal(bc)
  if err != nil {
    return err
  }
  err = os.WriteFile("blockchain.json", data, 0644)
  if err != nil {
    return err
  }
  return nil
}

func LoadBlockchain() (Blockchain, error) {
  data, err := os.ReadFile("blockchain.json")
  if err != nil {
    return Blockchain{}, err
  }
  var bc Blockchain
  err = json.Unmarshal(data, &bc)
  if err != nil {
    return Blockchain{}, err
  }
  return bc, nil
}


