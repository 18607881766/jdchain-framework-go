package ledger_model

import (
	"time"

	binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"
	"github.com/blockchain-jd-com/framework-go/crypto"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
)

/*
 * Author: imuge
 * Date: 2020/5/29 上午9:17
 */

var _ TransactionBuilder = (*TxBuilder)(nil)

type TxBuilder struct {
	ledgerHash    *framework.HashDigest
	hashAlgorithm framework.CryptoAlgorithm
	opFactory     BlockchainOperationFactory
}

func (t *TxBuilder) Settings() *SettingsOperationBuilder {
	return t.opFactory.Settings()
}

func (t *TxBuilder) Consensus() *ConsensusSettingsUpdateOperationBuilder {
	return t.opFactory.Consensus()
}

func (t *TxBuilder) MetaInfo() *MetaInfoUpdateOperationBuilder {
	return t.opFactory.MetaInfo()
}

func (t *TxBuilder) User(address []byte) *UserUpdateOperationBuilder {
	return t.opFactory.User(address)
}

func (t *TxBuilder) Contract(address []byte) *ContractOperationBuilder {
	return t.opFactory.Contract(address)
}

func (t *TxBuilder) ContractEvents() *ContractEventSendOperationBuilder {
	return t.opFactory.ContractEvents()
}

func (t *TxBuilder) EventAccounts() *EventAccountRegisterOperationBuilder {
	return t.opFactory.EventAccounts()
}

func (t *TxBuilder) EventAccount(accountAddress []byte) *EventAccountOperationBuilder {
	return t.opFactory.EventAccount(accountAddress)
}

func (t *TxBuilder) Participants() *ParticipantRegisterOperationBuilder {
	return t.opFactory.Participants()
}

func (t *TxBuilder) States() *ParticipantStateUpdateOperationBuilder {
	return t.opFactory.States()
}

func NewTxBuilder(ledgerHash *framework.HashDigest, hashAlgorithm framework.CryptoAlgorithm) *TxBuilder {
	return &TxBuilder{
		ledgerHash:    ledgerHash,
		hashAlgorithm: hashAlgorithm,
		opFactory:     BlockchainOperationFactory{},
	}
}

func (t *TxBuilder) Contracts() *ContractCodeDeployOperationBuilder {
	return t.opFactory.Contracts()
}

func (t *TxBuilder) Security() *SecurityOperationBuilder {
	return t.opFactory.Security()
}

func (t *TxBuilder) GetLedgerHash() *framework.HashDigest {
	return t.ledgerHash
}

func (t *TxBuilder) Users() *UserRegisterOperationBuilder {
	return t.opFactory.Users()
}

func (t *TxBuilder) DataAccounts() *DataAccountRegisterOperationBuilder {
	return t.opFactory.DataAccounts()
}

func (t *TxBuilder) DataAccount(accountAddress []byte) *DataAccountOperationBuilder {
	return t.opFactory.DataAccount(accountAddress)
}

func (t *TxBuilder) PrepareRequestNow() (TransactionRequestBuilder, error) {
	return t.PrepareRequest(time.Now().UnixNano() / 1e6)
}

func (t *TxBuilder) PrepareRequest(time int64) (TransactionRequestBuilder, error) {
	txContent := t.PrepareContent(time)
	transactionHash, err := ComputeTxContentHash(t.hashAlgorithm, txContent)
	if err != nil {
		return nil, err
	}
	return NewTxRequestBuilder(transactionHash, txContent), nil
}

func (t *TxBuilder) PrepareContentNow() *TransactionContent {
	return t.PrepareContent(time.Now().UnixNano() / 1e6)
}

func (t *TxBuilder) PrepareContent(time int64) *TransactionContent {
	return NewTransactionContent(t.ledgerHash, t.opFactory.operationList, time)
}

func ComputeTxContentHash(algorithm framework.CryptoAlgorithm, txContent *TransactionContent) (*framework.HashDigest, error) {
	contentBodyBytes, err := binary_proto.NewCodec().Encode(txContent)
	if err != nil {
		return nil, err
	}
	contentHash := crypto.GetHashFunction(algorithm).Hash(contentBodyBytes)
	return contentHash, nil
}

func VerifyTxContentHash(algorithm framework.CryptoAlgorithm, txContent *TransactionContent, verifiedHash *framework.HashDigest) bool {
	hash, err := ComputeTxContentHash(algorithm, txContent)
	if err != nil {
		return false
	}
	return hash.Equals(verifiedHash)
}
