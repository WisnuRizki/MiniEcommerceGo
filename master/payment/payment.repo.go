package payment

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	_ "github.com/midtrans/midtrans-go/iris"
	"github.com/midtrans/midtrans-go/snap"
)


func CreateSnap(transNumber int,totalPayment int64) *snap.Response{
	// 1. Set you ServerKey with globally
midtrans.ServerKey = "SB-Mid-server-pGwf5OGA9V8ic8VFlc4mIpu9"
midtrans.Environment = midtrans.Sandbox
id := strconv.Itoa(transNumber)
// 2. Initiate Snap request
req := & snap.Request{
	TransactionDetails: midtrans.TransactionDetails{
		OrderID:  id, 
		GrossAmt: totalPayment,
	}, 
	CreditCard: &snap.CreditCardDetails{
		Secure: true,
	},
}

// 3. Request create Snap transaction to Midtrans
	snapResp, _ := snap.CreateTransaction(req)
	return snapResp
}

func TryCallBackMidtrans(c *gin.Context){
	var notificationPayload map[string]interface{}

	err := c.BindJSON(&notificationPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Something went wrong",
		})
		return
	}

	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		// do something when key `order_id` not found
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "OrderId from midtrans is missing",
		})
		return
	}

	transactionStatusResp, e := coreapi.CheckTransaction(orderId)
	if e != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Something went wrong",
		})
		return
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					// TODO set transaction status on your database to 'success'
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				// TODO set transaction status on your databaase to 'success'
				fmt.Printf("%s",orderId)
				// c.JSON(http.StatusInternalServerError,gin.H{
				// 	"message": "Transaction with id: " + " Success",
				// })
				return
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
			}
		}
	}
}