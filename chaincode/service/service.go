package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	return nil
}

func JsonHashS256(data any) (string, error) {
	if data == nil {
		return "", errors.New("data is nil")
	}

	bytesData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	var bb bytes.Buffer
	if err := json.Compact(&bb, bytesData); err != nil {
		return "", err
	}

	h := sha256.New()

	h.Write(bb.Bytes())

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func ParseRequest(data string, obj any) error {
	if data == "" {
		return errors.New("data is empty")
	}

	dataBytes, err := base64.RawStdEncoding.DecodeString(data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(dataBytes, obj); err != nil {
		return err
	}

	return nil
}
