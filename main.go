package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/nervosnetwork/ckb-sdk-go/rpc"
)

type Script struct {
	CODE_HASH string `json:"CODE_HASH"`
	HASH_TYPE string `json:"HASH_TYPE"`
	TX_HASH   string `json:"TX_HASH"`
	INDEX     string `json:"INDEX"`
	DEP_TYPE  string `json:"DEP_TYPE"`
	SHORT_ID  *int   `json:"SHORT_ID,omitempty"`
}

type Scripts struct {
	SECP256K1_BLAKE160          Script `json:"SECP256K1_BLAKE160"`
	SECP256K1_BLAKE160_MULTISIG Script `json:"SECP256K1_BLAKE160_MULTISIG"`
	DAO                         Script `json:"DAO"`
}

type Config struct {
	PREFIX  string  `json:"PREFIX"`
	SCRIPTS Scripts `json:"SCRIPTS"`
}

func generate(url string, path string) {
	secpCodeHash := "0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"
	multisigCodeHash := "0x5c5069eb0857efc65e1bca0c07df34c31663b3622fd3876c876320fc9634e2a8"
	daoCodeHash := "0x82d76d1b75fe2fd9a27dfbaa65a039221a380d76c926f378d3f81cf3e7e13f2e"

	mainnetGenesisBlockHash := "0x92b197aa1fba0f63633922c61c92375c9c074a93e85963554f5499fe1450d0e5"

	client, err := rpc.Dial(url)
	if err != nil {
		log.Fatalf("Create rpc client error: %v", err)
	}

	block, err := client.GetBlockByNumber(context.Background(), 0)
	if err != nil {
		log.Fatalf("Can't get genesis block: %v", err)
	}

	// secp
	codeTransaction := block.Transactions[0]
	depGroupTransaction := block.Transactions[1]

	secpShortId := 0
	secpScript := Script{
		CODE_HASH: secpCodeHash,
		HASH_TYPE: "type",
		TX_HASH:   depGroupTransaction.Hash.String(),
		INDEX:     "0x0",
		DEP_TYPE:  "dep_group",
		SHORT_ID:  &secpShortId,
	}

	// multisig
	multisigShortId := 1
	multisigScript := Script{
		CODE_HASH: multisigCodeHash,
		HASH_TYPE: "type",
		TX_HASH:   depGroupTransaction.Hash.String(),
		INDEX:     "0x1",
		DEP_TYPE:  "dep_group",
		SHORT_ID:  &multisigShortId,
	}

	// dao
	daoScript := Script{
		CODE_HASH: daoCodeHash,
		HASH_TYPE: "type",
		TX_HASH:   codeTransaction.Hash.String(),
		INDEX:     "0x2",
		DEP_TYPE:  "code",
		SHORT_ID:  nil,
	}

	var prefix string
	if block.Header.Hash.String() == mainnetGenesisBlockHash {
		prefix = "ckb"
	} else {
		prefix = "ckt"
	}
	config := Config{
		PREFIX: prefix,
		SCRIPTS: Scripts{
			SECP256K1_BLAKE160:          secpScript,
			SECP256K1_BLAKE160_MULTISIG: multisigScript,
			DAO:                         daoScript,
		},
	}

	j, _ := json.MarshalIndent(config, "", "  ")
	ioutil.WriteFile(path, j, os.ModePerm)
}

func main() {
	path := "config.json"
	url := "http://127.0.0.1:8114"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	if len(os.Args) > 2 {
		url = os.Args[2]
	}

	generate(url, path)
}
