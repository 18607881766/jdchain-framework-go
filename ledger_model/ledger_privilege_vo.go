package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/28 下午2:24
 */

type LedgerPrivilegeVO struct {
	Privilege       []LedgerPermission
	PermissionCount int32
}
