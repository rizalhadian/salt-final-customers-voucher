package http_request

type Transaction struct {
	Customer_id                 int                           `json:"customer_id,omitempty"`
	TransactionsItems           []TransactionsItem            `json:"items,omitempty"`
	TransactionsVoucherRedeemed []TransactionsVoucherRedeemed `json:"vouchers_redeemed,omitempty"`
}

type TransactionsItem struct {
	Item_id       int64   `json:"item_id,omitempty"`
	Items_type_id int64   `json:"items_type_id,omitempty"`
	Price         float64 `json:"price,omitempty"`
	Qty           int32   `json:"qty,omitempty"`
	Note          string  `json:"note,omitempty"`
}

type TransactionsVoucherRedeemed struct {
	Code string `json:"code,omitempty"`
}
