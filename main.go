package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	_ "strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var wallets []common.Address
var values []*big.Int
var tokens *big.Int

func main() {

	conf := GetConfig()

	contractAddress := common.HexToAddress(conf.ContractAddress)

	privateKey := getPrivateKeyContent(conf.PrivateKeyPath)
	privateKeyPassphrase := conf.PrivateKeyPassphrase
	node := conf.NodeAddress

	perTxAddress := conf.BatchSize

	client, err := ethclient.Dial(node)

	if err != nil {
		log.Printf("Failed to connect ethereum client [%s], Error [%v]", node, err)
	}

	log.Printf("Connection established %v", client)

	auth, err := bind.NewTransactor(strings.NewReader(privateKey), privateKeyPassphrase)

	if err != nil {
		log.Fatalf("Invalid private key password! Failed to create authorized transactor: %v ", err)
	}

	log.Printf("Auth: %x", auth.From)

	contract, err := bindAirdropContract(contractAddress, client, client, nil)

	if err != nil {
		log.Fatalf("Failed to instantiate Freso Airdrop contract: %v", err)
	}

	log.Println("Fresco Airdrop contract is instanticated!")

	tOpts := bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: conf.MaxGasLimit,
	}

	tts := &AirdropContractTransactorSession{
		Contract: &AirdropContractTransactor{
			contract: contract,
		},
		TransactOpts: tOpts,
	}
	// log.Println(tts)
	recipents := readCSV(conf.RecipientsCSVPath)

	r := csv.NewReader(strings.NewReader(recipents))

	records, err := r.ReadAll()

	if err != nil {
		log.Fatalf("Error occured while parsing the csv file: %v", err)
	}

	recordsLen := len(records)

	if recordsLen < 1 {
		log.Fatalf("CSV file is empty!")
	}

	tokens, ok := new(big.Int).SetString(strings.TrimSpace(conf.NumberOfTokens), 10)
	if !ok {
		log.Fatal("Number of tokens is invalid!")
	}

	counter := 0
	for i := 0; i < recordsLen; i++ {

		counter++

		record := records[i]
		record_elements := len(record)
		address := common.HexToAddress(record[0])
		wallets = append(wallets, address)

		tkns := tokens
		if record_elements > 1 {
			v, ok := new(big.Int).SetString(strings.TrimSpace(record[1]), 10)
			if ok {
				tkns = v
			}
		}
		values = append(values, tkns)
		// log.Printf("%d Address %x will receive %d Tokens", i, address, tokens)
		// log.Printf("Wallets count %d values %d ", len(wallets), len(values))
		if (i > (perTxAddress-1) && i%perTxAddress == 0) || (i+1 == recordsLen) {
			// log.Println(tts)
			log.Printf("Record %d to %d Wallets count %d values %d", i-perTxAddress, i, len(wallets), len(values))

			// log.Printf("Transfering [%s] NKT to address [%x]", tokens, address)
			tryToTransfer := true
			for tryToTransfer {
				transaction, err := tts.Send(wallets, values)

				if err != nil {
					log.Printf("Failed to transfer token to address [%x] ... Retrying in 30 seconds, %v", address, err)

					select {
					case <-time.After(time.Second * 30):
						continue
					}
				}
				tryToTransfer = false

				log.Printf("Trnsaction accepted by network. Tx %x", transaction.Hash())
				wallets = nil
				values = nil
				select {
				case <-time.After(time.Second * 45):
					continue
				}

			}
		}
		// else {
		// 	log.Printf("Error parsing input from csv [%s]", record[0])
		// }
	}

}

func readCSV(path string) string {

	if _, err := os.Stat(path); err != nil {
		log.Fatalf("CSV file [%s] is not readable or doesn't exists.", path)
	}

	log.Printf("Reading the csv file [%s]", path)

	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("Error occured while reading the content of csv file %v", err)
	}

	return string(bytes)
}

func getPrivateKeyContent(path string) string {

	if _, err := os.Stat(path); err != nil {
		log.Fatalf("Private key [%s] is not readable or doesn't exists.", path)
	}

	log.Printf("Reading the private key [%s]", path)

	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("Error occured while reading the content of private key %v", err)
	}

	return string(bytes)

}
