package mapper_gorm_mysql

import (
	"salt-final-voucher/domain/entity"
	model_gorm_mysql "salt-final-voucher/internal/repository/gorm/models"
)

func CustomersVoucherEntityToModel(entity *entity.CustomersVoucher) *model_gorm_mysql.ModelCustomersVoucher {
	model := model_gorm_mysql.ModelCustomersVoucher{}
	if entity.GetId() != 0 {
		model.Id = entity.GetId()
	}
	model.Customer_id = entity.GetCustomerId()
	model.Voucher_id = entity.GetVoucherId()

	return &model
}

func CustomersVoucherModelToEntity(model *model_gorm_mysql.ModelCustomersVoucher) *entity.CustomersVoucher {
	entity := &entity.CustomersVoucher{}
	entity.SetId(model.Id)
	entity.SetCustomerId(model.Customer_id)
	entity.SetVoucherId(model.Voucher_id)
	entity.SetVoucherName(model.Voucher_name)
	entity.SetCode(model.Code)
	entity.SetExpiredAt(model.Expired_at)
	entity.SetTransactionId(model.Transaction_id)
	entity.SetTotalAmount(model.Total_amount)
	entity.SetTotalDiscountAmount(model.Total_discount_amount)
	entity.SetFinalTotalAmount(model.Final_total_amount)
	entity.SetStatus(model.Status)
	entity.SetCreatedAt(model.CreatedAt)
	entity.SetUpdatedAt(model.Updated_at)
	entity.SetDeletedAt(model.Deleted_at)
	return entity
}
