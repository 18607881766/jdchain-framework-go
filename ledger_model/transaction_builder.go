package ledger_model

import "framework-go/crypto/framework"

/*
 * Author: imuge
 * Date: 2020/5/29 上午9:11
 */

// 区块链交易模板
type TransactionBuilder interface {
	ClientOperator

	GetLedgerHash() framework.HashDigest

	// 基于当前的系统时间完成交易定义，并生成就绪的交易数据
	// 调用此方法后，不能再向当前对象加入更多的操作
	PrepareRequestNow() TransactionRequestBuilder

	// 生成交易内容
	PrepareContentNow() TransactionContent

	// 基于当前的系统时间完成交易定义，并生成就绪的交易数据
	PrepareRequest(time int64) TransactionRequestBuilder

	// 生成交易内容
	PrepareContent(time int64) TransactionContent
}
