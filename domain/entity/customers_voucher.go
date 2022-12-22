package entity

import (
	"database/sql"
	"errors"
	"time"
)

type CustomersVoucherStatus struct {
	code        int16
	section     string
	name        string
	description string
}

var customersVoucherStatuses = []CustomersVoucherStatus{
	{
		code:        120,
		section:     "customers_voucher",
		name:        "redeemable",
		description: "Voucher is available",
	},
	{
		code:        121, // Status ini untuk mengantisipasi ketika gagal update voucher setelah transaksi berhasil dilakukan
		section:     "customers_voucher",
		name:        "redeeming",
		description: "Redeeming voucher, and waiting response of the Transaction is succeed or not",
	},
	{
		code:        112,
		section:     "customers_voucher",
		name:        "redeemed",
		description: "Voucher is redeemed and Transaction is succeed",
	},
}

// Perlu ada penambahan transactions items apa saja yang menggunakan voucher
type CustomersVoucher struct {
	id                    int64
	customer_id           int64
	voucher_id            int64
	voucher_name          string
	code                  string
	expired_at            time.Time
	transaction_id        int64
	total_amount          float64
	total_discount_amount float64
	final_total_amount    float64
	status                int16
	voucher               Voucher
	created_at            time.Time
	updated_at            sql.NullTime
	deleted_at            sql.NullTime
}

type DTOCustomersVoucher struct {
	Id                    int64
	Customer_id           int64
	Voucher_id            int64
	Voucher_name          string
	Code                  string
	Expired_at            time.Time
	Transaction_id        int64
	Total_amount          float64
	Total_discount_amount float64
	Final_total_amount    float64
	Status                int16
	Voucher               Voucher
	Created_at            time.Time
	Updated_at            sql.NullTime
	Deleted_at            sql.NullTime
}

func NewCustomersVoucher(dto *DTOCustomersVoucher) (*CustomersVoucher, error) {

	if dto.Voucher_id == 0 {
		return nil, errors.New("voucher_id is required")
	}

	if dto.Voucher_name == "" {
		return nil, errors.New("voucher_name is required")
	}

	if dto.Code == "" {
		return nil, errors.New("voucher's code is required")
	}

	if dto.Status == 0 {
		return nil, errors.New("voucher's status is required")
	}

	entity_customers_voucher := &CustomersVoucher{
		id:                    dto.Id,
		customer_id:           dto.Customer_id,
		voucher_id:            dto.Voucher_id,
		voucher_name:          dto.Voucher_name,
		code:                  dto.Code,
		expired_at:            dto.Expired_at,
		transaction_id:        dto.Transaction_id,
		total_amount:          dto.Total_amount,
		total_discount_amount: dto.Total_discount_amount,
		final_total_amount:    dto.Final_total_amount,
		status:                dto.Status,
		created_at:            dto.Created_at,
		updated_at:            dto.Updated_at,
		deleted_at:            dto.Deleted_at,
	}

	return entity_customers_voucher, nil
}

// Getter
func (cv *CustomersVoucher) GetId() int64 {
	return cv.id
}

func (cv *CustomersVoucher) GetCustomerId() int64 {
	return cv.customer_id
}

func (cv *CustomersVoucher) GetVoucherId() int64 {
	return cv.voucher_id
}

func (cv *CustomersVoucher) GetVoucherName() string {
	return cv.voucher_name
}

func (cv *CustomersVoucher) GetCode() string {
	return cv.code
}

func (cv *CustomersVoucher) GetExpiredAt() time.Time {
	return cv.expired_at
}

func (cv *CustomersVoucher) GetTransactionId() int64 {
	return cv.transaction_id
}

func (cv *CustomersVoucher) GetTotalAmount() float64 {
	return cv.total_amount
}

func (cv *CustomersVoucher) GetTotalDiscountAmount() float64 {
	return cv.total_discount_amount
}

func (cv *CustomersVoucher) GetFinalTotalAmount() float64 {
	return cv.final_total_amount
}

func (cv *CustomersVoucher) GetStatus() int16 {
	return cv.status
}

func (cv *CustomersVoucher) GetVoucher() Voucher {
	return cv.voucher
}

func (cv *CustomersVoucher) GetCreatedAt() time.Time {
	return cv.created_at
}

func (cv *CustomersVoucher) GetUpdatedAt() sql.NullTime {
	return cv.updated_at
}

func (cv *CustomersVoucher) GetDeletedAt() sql.NullTime {
	return cv.deleted_at
}

// Setter

func (cv *CustomersVoucher) SetId(value int64) {
	cv.id = value
}

func (cv *CustomersVoucher) SetCustomerId(value int64) {
	cv.customer_id = value
}

func (cv *CustomersVoucher) SetVoucherId(value int64) {
	cv.voucher_id = value
}

func (cv *CustomersVoucher) SetVoucherName(value string) {
	cv.voucher_name = value
}

func (cv *CustomersVoucher) SetCode(value string) {
	cv.code = value
}

func (cv *CustomersVoucher) SetExpiredAt(value time.Time) {
	cv.expired_at = value
}

func (cv *CustomersVoucher) SetTransactionId(value int64) {
	cv.transaction_id = value
}

func (cv *CustomersVoucher) SetTotalAmount(value float64) {
	cv.total_amount = value
}

func (cv *CustomersVoucher) SetTotalDiscountAmount(value float64) {
	cv.total_discount_amount = value
}

func (cv *CustomersVoucher) SetFinalTotalAmount(value float64) {
	cv.final_total_amount = value
}

func (cv *CustomersVoucher) SetStatus(value int16) {
	cv.status = value
}

func (cv *CustomersVoucher) SetVoucher(value Voucher) {
	cv.voucher = value
}

func (cv *CustomersVoucher) SetCreatedAt(value time.Time) {
	cv.created_at = value
}

func (cv *CustomersVoucher) SetUpdatedAt(value sql.NullTime) {
	cv.updated_at = value
}

func (cv *CustomersVoucher) SetDeletedAt(value sql.NullTime) {
	cv.deleted_at = value
}
