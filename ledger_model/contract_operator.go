package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/1 下午7:23
 */

type ContractOperator interface {

	// 部署合约
	Contracts() *ContractCodeDeployOperationBuilder

	// 合约更新操作
	Contract(address []byte) *ContractOperationBuilder

	// contract events
	ContractEvents() *ContractEventSendOperationBuilder
}
