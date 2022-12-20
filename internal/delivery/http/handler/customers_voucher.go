package http_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"salt-final-voucher/domain/interface_usecase"
	http_response "salt-final-voucher/internal/delivery/http/response"
	"strconv"

	"github.com/gorilla/mux"
)

type HandlerCustomersVoucher struct {
	usecase_customers_voucher interface_usecase.InterfaceRepoCustomersVoucher
}

func NewHandlerCustomersVoucher(router *mux.Router, usecase_customers_voucher_val interface_usecase.InterfaceRepoCustomersVoucher) {
	HandlerCustVoucher := &HandlerCustomersVoucher{
		usecase_customers_voucher: usecase_customers_voucher_val,
	}

	// router.HandleFunc("/api/customer/{customer_id}/transaction", HandlerTrans.GetListByCustomer).Methods(http.MethodGet)
	// router.HandleFunc("/api/customer/{customer_id}/transaction", HandlerTrans.Store).Methods(http.MethodPost)
	// router.HandleFunc("/api/customer/{customer_id}/transaction/{id}", HandlerTrans.FindById).Methods(http.MethodGet)
	// router.HandleFunc("/api/customer/{customer_id}/transaction/{id}", HandlerTrans.Update).Methods(http.MethodPut)
	// router.HandleFunc("/api/customer/{customer_id}/transaction/{id}", HandlerTrans.Delete).Methods(http.MethodDelete)

	// router.HandleFunc("/api/transaction", HandlerTrans.GetList).Methods(http.MethodGet)

	router.HandleFunc("/api/voucher/generate/{customer_id}", HandlerCustVoucher.Generate).Methods(http.MethodGet)
	router.HandleFunc("/api/voucher/redeem", HandlerCustVoucher.Redeem).Methods(http.MethodPost)
	router.HandleFunc("/api/voucher/redeem", HandlerCustVoucher.RedeemUpdate).Methods(http.MethodPut)
}

func (hcv *HandlerCustomersVoucher) Generate(w http.ResponseWriter, r *http.Request) {

	customer_id_string := mux.Vars(r)["customer_id"]
	customer_id, customer_id_conv_err := strconv.Atoi(customer_id_string)

	if customer_id_conv_err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		customers_vouchers, customers_vouchers_err := hcv.usecase_customers_voucher.Generate(r.Context(), int64(customer_id))
		if customers_vouchers_err != nil {
			if customers_vouchers_err.Error() == "404" {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			if customers_vouchers_err.Error() == "500" {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if customers_vouchers_err.Error() == "Cannot Use customer_id 0" {
				resp_skeleton_err := http_response.SkeletonError{
					Success: false,
					Message: "Cannot Generate Voucher. Customer_Id 0 Is Non Member Transaction.",
				}
				resp, resp_json_err := json.Marshal(resp_skeleton_err)
				if resp_json_err != nil {
					fmt.Println(resp_json_err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				w.Write(resp)
				return
			}
		}

		var customers_vouchers_response []http_response.CustomersVoucher

		for _, customers_voucher := range customers_vouchers {
			customers_voucher_response := http_response.CustomersVoucher{
				Id:           customers_voucher.GetId(),
				Customer_id:  customers_voucher.GetCustomerId(),
				Voucher_id:   customers_voucher.GetVoucherId(),
				Voucher_name: customers_voucher.GetVoucherName(),
				Code:         customers_voucher.GetCode(),
				Expired_at:   customers_voucher.GetExpiredAt(),
				Status:       customers_voucher.GetStatus(),
				Created_at:   customers_voucher.GetCreatedAt(),
			}
			customers_vouchers_response = append(customers_vouchers_response, customers_voucher_response)
		}

		resp_skeleton := http_response.SkeletonSingleResponse{
			Success: true,
			Message: "",
			Data:    customers_vouchers_response,
		}

		resp, resp_json_err := json.Marshal(resp_skeleton)
		if resp_json_err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
	return
}

func (hcv *HandlerCustomersVoucher) Redeem(w http.ResponseWriter, r *http.Request) {

}

func (hcv *HandlerCustomersVoucher) RedeemUpdate(w http.ResponseWriter, r *http.Request) {

}
