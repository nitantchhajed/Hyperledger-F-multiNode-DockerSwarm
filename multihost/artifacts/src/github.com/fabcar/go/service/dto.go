package service

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/Subskribo-BV/dnn-fabric-chaincode/common/contract"
)

const (
	ContractStateActive   = "active"
	ContractStateVoided   = "voided"
	ContractStateExpired  = "expired"
	ContractStateReleased = "released"
)

type NewAssetReq struct {
	ImmutableContract     contract.ImmutableContract `json:"immutable_contract"`
	ImmutableContractHash string                     `json:"immutable_contract_hash"`
	NotaryOU              string                     `json:"notary_ou"`
}

type VoidAssetReq struct {
	ImmutableContract     contract.ImmutableContract `json:"immutable_contract"`
	ImmutableContractHash string                     `json:"immutable_contract_hash"`

	ContractId  int64  `json:"contract_id"`
	PackageId   int64  `json:"packageId"`
	PackageHash string `json:"packageHash"`
	// todo: add void specific fields
}

type ExpireAssetReq struct {
	ImmutableContract     contract.ImmutableContract `json:"immutable_contract"`
	ImmutableContractHash string                     `json:"immutable_contract_hash"`
	NotaryOU              string                     `json:"notary_ou"`

	ContractId  int64  `json:"contract_id"`
	PackageId   int64  `json:"packageId"`
	PackageHash string `json:"packageHash"`
	// todo: add expire specific fields
}

type ReleaseAssetReq struct {
	ImmutableContract     contract.ImmutableContract `json:"immutable_contract"`
	ImmutableContractHash string                     `json:"immutable_contract_hash"`
	NotaryOU              string                     `json:"notary_ou"`

	ContractId  int64  `json:"contract_id"`
	PackageId   int64  `json:"packageId"`
	PackageHash string `json:"packageHash"`
	// todo: add release specific fields
}

type Contract struct {
	ContractId   int64    `json:"contract_id"`
	Version      int64    `json:"version"`
	ContractHash string   `json:"contractHash"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
	State        string   `json:"state"`
	Changes      []Change `json:"changes"`
}

type Change struct {
	PackageID   int64  `json:"package_id"`
	PackageHash string `json:"package_hash"`
	PackageDate string `json:"package_date"`
	CallerSdn   string `json:"caller_sdn"`
	Action      string `json:"action"`
	NewState    string `json:"new_state"`
}

func (e *Contract) Checksum() string {
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(e.ContractHash))))
}
