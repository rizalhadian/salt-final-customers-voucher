package interface_repo

import (
	"context"
	"salt-final-voucher/domain/entity"
)

type InterfaceRepoCustomersVoucher interface {
	GetByCode(ctx context.Context, code string) (*entity.CustomersVoucher, error)
	GetByVoucherId(ctx context.Context, voucher_id int32) (*entity.CustomersVoucher, error)
	Store(ctx context.Context, customers_voucher *entity.CustomersVoucher) error
	Update(ctx context.Context, customers_voucher *entity.CustomersVoucher) error
	Delete(ctx context.Context, code string) error
}
