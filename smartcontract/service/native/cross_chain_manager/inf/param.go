package inf

import (
	"github.com/ontio/multi-chain/common"
	"github.com/ontio/multi-chain/smartcontract/service/native"
	"github.com/ontio/multi-chain/smartcontract/service/native/utils"
)

type ChainHandler interface {
	Verify(service *native.NativeService) (*EntranceParam, error)
	MakeTransaction(service *native.NativeService, param *EntranceParam) error
}

type EntranceParam struct {
	SourceChainID  uint64
	TxData         string
	Height         uint32
	Proof          string
	RelayerAddress string
	TargetChainID  uint64
}

func (this *EntranceParam) Deserialization(source *common.ZeroCopySource) error {
	sourcechainid, err := utils.DecodeVarUint(source)
	if err != nil {
		return err
	}
	txdata, err := utils.DecodeString(source)
	if err != nil {
		return err
	}
	height, err := utils.DecodeVarUint(source)
	if err != nil {
		return err
	}
	proof, err := utils.DecodeString(source)
	if err != nil {
		return err
	}
	relayerAddr, err := utils.DecodeString(source)
	if err != nil {
		return err
	}
	targetchainid, err := utils.DecodeVarUint(source)
	if err != nil {
		return err
	}

	this.SourceChainID = sourcechainid
	this.TxData = txdata
	this.Height = uint32(height)
	this.Proof = proof
	this.RelayerAddress = relayerAddr
	this.TargetChainID = targetchainid

	return nil
}

func (this *EntranceParam) Serialization(sink *common.ZeroCopySink) {
	utils.EncodeVarUint(sink, this.SourceChainID)
	utils.EncodeString(sink, this.TxData)
	utils.EncodeVarUint(sink, uint64(this.Height))
	utils.EncodeString(sink, this.Proof)
	utils.EncodeString(sink, this.RelayerAddress)
	utils.EncodeVarUint(sink, this.TargetChainID)
}
