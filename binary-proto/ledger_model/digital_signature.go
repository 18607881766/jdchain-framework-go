package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/26 下午1:45
 */

var _ binary_proto.DataContract = (*DigitalSignature)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(DigitalSignature{})
}

type DigitalSignature struct {
	PubKey []byte `primitiveType:"BYTES"`
	Digest []byte `primitiveType:"BYTES"`
}

func (d DigitalSignature) Code() int32 {
	return binary_proto.DIGITALSIGNATURE
}

func (d DigitalSignature) Name() string {
	return "DigitalSignature"
}

func (d DigitalSignature) Description() string {
	return ""
}