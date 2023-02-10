package helper

import (
	"Gurumu/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

// Initiate coreapi client
func MidtransCoreAPIClient() coreapi.Client {
	c := coreapi.Client{}
	c.New(config.MIDTRANSSERVERKEY, midtrans.Sandbox)
	return c
}

// Initiate charge request
func CreateReservasiTransaction(kodePembayaran string, totalTarif int, metodePembayaran string) *coreapi.ChargeResponse {
	c := MidtransCoreAPIClient()

	//Initiate charge request by payment method
	switch {
	//Bank BCA
	case metodePembayaran == "transfer_va_bca":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  kodePembayaran,
				GrossAmt: int64(totalTarif),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBca,
			},
			CustomExpiry: &coreapi.CustomExpiry{
				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		coreApiRes, _ := c.ChargeTransaction(chargeReq)
		return coreApiRes
	//Bank BRI
	case metodePembayaran == "transfer_va_bri":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  kodePembayaran,
				GrossAmt: int64(totalTarif),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBri,
			},
			CustomExpiry: &coreapi.CustomExpiry{
				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		coreApiRes, _ := c.ChargeTransaction(chargeReq)
		return coreApiRes
	//Bank BNI
	case metodePembayaran == "transfer_va_bni":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  kodePembayaran,
				GrossAmt: int64(totalTarif),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBni,
			},
			CustomExpiry: &coreapi.CustomExpiry{
				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		coreApiRes, _ := c.ChargeTransaction(chargeReq)
		return coreApiRes
	//Bank Permata
	case metodePembayaran == "transfer_va_permata":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  kodePembayaran,
				GrossAmt: int64(totalTarif),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankPermata,
			},
			CustomExpiry: &coreapi.CustomExpiry{
				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		coreApiRes, _ := c.ChargeTransaction(chargeReq)
		return coreApiRes

	//QRIS (Gopay)
	case metodePembayaran == "qris":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "qris",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  kodePembayaran,
				GrossAmt: int64(totalTarif),
			},
			CustomExpiry: &coreapi.CustomExpiry{
				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		coreApiRes, _ := c.ChargeTransaction(chargeReq)
		return coreApiRes
	}
	return &coreapi.ChargeResponse{}
}

// get status of transaction that already recorded on midtrans (already `charge`-ed)
func CheckStatusPayment(kodePembayaran string) (*coreapi.TransactionStatusResponse, error) {
	c := MidtransCoreAPIClient()

	res, err := c.CheckTransaction(kodePembayaran)
	if err != nil {
		return nil, err
	}

	return res, nil
}
