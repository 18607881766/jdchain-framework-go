package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/6/9 下午4:57
 */

var _ binary_proto.DataContract = (*ParticipantStateUpdateOperation)(nil)

func init() {
	binary_proto.RegisterContract(ParticipantRegisterOperation{})
}

type ParticipantStateUpdateOperation struct {
	ParticipantID *BlockchainIdentity  `refContract:"144"`
	State         ParticipantNodeState `refEnum:"2852"`
}

func (p ParticipantStateUpdateOperation) ContractCode() int32 {
	return binary_proto.TX_OP_PARTICIPANT_STATE_UPDATE
}

func (p ParticipantStateUpdateOperation) ContractName() string {
	return OperationTypeParticipantStateUpdateOperation
}

func (p ParticipantStateUpdateOperation) Description() string {
	return ""
}
