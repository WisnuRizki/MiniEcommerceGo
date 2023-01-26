package payment

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"miniecommerce.wisnu.net/master/history"
	"miniecommerce.wisnu.net/master/transaction"

	"github.com/midtrans/midtrans-go/coreapi"
)

func CallBackMidtrans(c *gin.Context){
	var notificationPayload map[string]interface{}
	history := history.History{}
	transaction := transaction.Transaction{}

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
	id,_ := strconv.Atoi(orderId)

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
				
				// Set status transaction ke Success
				err = transaction.UpdateStatus("Success",id)
				if err != nil {
					fmt.Print("error transaction")
				}
				// Set status history ke Success
				err = history.UpdateStatus("Success",id)
				if err != nil {
					fmt.Print("error history")
				}
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