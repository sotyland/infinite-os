package dto

import (
	"github.com/goinfinite/os/src/domain/valueObject"
)

type DeleteAccount struct {
	AccountId         valueObject.AccountId `json:"accountId"`
	OperatorAccountId valueObject.AccountId `json:"-"`
	OperatorIpAddress valueObject.IpAddress `json:"-"`
}

func NewDeleteAccount(
	accountId, operatorAccountId valueObject.AccountId,
	operatorIpAddress valueObject.IpAddress,
) DeleteAccount {
	return DeleteAccount{
		AccountId:         accountId,
		OperatorAccountId: operatorAccountId,
		OperatorIpAddress: operatorIpAddress,
	}
}
