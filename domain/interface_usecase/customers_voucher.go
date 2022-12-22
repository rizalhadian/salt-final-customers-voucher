package interface_usecase

import (
	"context"
	"salt-final-voucher/domain/entity"
)

type InterfaceRepoCustomersVoucher interface {
	Redeem(ctx context.Context, dto_transaction *entity.DTOTransaction, dto_transactions_items []*entity.DTOTransactionsItem, dto_transactions_vouchers_redeem []*entity.DTOCustomersVoucher) ([]*entity.CustomersVoucher, error)
	Generate(ctx context.Context, customer_id int64) ([]*entity.CustomersVoucher, error)
	// Update(ctx context.Context, customers_voucher *entity.CustomersVoucher) error
	// Delete(ctx context.Context, code string) error
}
