package operations

import (
	"github.com/niktrix/go-tezos/account"
	"github.com/niktrix/go-tezos/delegate"
)

type TezosOperationsService interface {
	CreateBatchPayment(payments []delegate.Payment, wallet account.Wallet, paymentFee int, gaslimit int, batchSize int) ([]string, error)
	InjectOperation(op string) ([]byte, error)
	GetBlockOperationHashes(id interface{}) ([]string, error)
	ActivateAccount(payments []delegate.Payment, wallet account.Wallet) ([]string, error)
	AccountReveal(payments []delegate.Payment, wallet account.Wallet) ([]string, error)
}
