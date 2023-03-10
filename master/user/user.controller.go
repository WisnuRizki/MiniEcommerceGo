package user

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"miniecommerce.wisnu.net/helpers"
	"miniecommerce.wisnu.net/master/balance"
	"miniecommerce.wisnu.net/master/history"
	"miniecommerce.wisnu.net/master/payment"
	"miniecommerce.wisnu.net/master/product"
	"miniecommerce.wisnu.net/master/transaction"
)

func (user *User) Register(c *gin.Context){
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Error with Json data",
		})
		return
	}

	// Check user already exist or not
	_,row := user.CheckDataByEmail(user.Email)
	if row > 0 {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "User Already Exist",
		})
		return
	}

	// hash password
	hashPassword,err := bcrypt.GenerateFromPassword([]byte(user.Password),10)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Something Went Wrong",
		})
		return
	}

	user.Password = string(hashPassword)
	user.ID = 0

	err = user.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Internal Server error",
		})
		return
	}


	// Membuat balance di akun baru
	balance := balance.Balance{}
	err = balance.CreateBalance(int(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Failed to create balance",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"data": user,
	})
}

func (user *User) Login(c *gin.Context){
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Error JSON",
		})
		return
	}

	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Need Email Or Password",
		})
		return
	}

	res,row := user.CheckDataByEmail(user.Email)
	if row == 0 {
		c.JSON(http.StatusNotFound,gin.H{
			"message": "User Not Found",
		})
		return
	}

	jwtString,pl := helpers.GenerateJWT(res.ID,user.Email)
	if pl != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Internal Server error",
		})
		return
	}

	pol := transaction.GetTrans()

	c.JSON(http.StatusOK,gin.H{
		"message": "Login Success",
		"token": jwtString,
		"pol": pol,
	})

}

// Buy Product

func (user *User) BuyProduct(c *gin.Context){
	history := []history.History{}
	balance := balance.Balance{}
	product := product.Product{}
	transaction := transaction.Transaction{}
	grandTotal := 0
	totalProduct := 0
	idMidleware := c.MustGet("id").(float64)
	
	err := c.BindJSON(&history)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Bad request json",
		})
		return 
	}

	// Get Amount Balance User
	amount,row := balance.CheckBalanceByUserId(int(idMidleware))
	if row == 0 {
		c.JSON(http.StatusNotFound,gin.H{
			"message": "Balance Not Found",
		})
		return 
	}

	// Check stock ,reduce stock & sum total_price
	for i,data := range history {
		// Check Stock
		stock,err := product.GetProductById(uint(data.ProductId))
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"message": "Product Not found",
			})
			return 
		}

		if stock.Quantity < data.Quantity {
			c.JSON(http.StatusBadRequest,gin.H{
				"message": "not enough stock",
			})
			return 
		}

		err = stock.UpdateProductStock(uint(data.ProductId),data.Quantity,stock.Quantity,"reduce")
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"message": "Failed To reduce stock",
			})
			return 
		}
		history[i].UserId = int(idMidleware)
		history[i].Status = "Pending"
		// total price
		grandTotal = grandTotal + int(data.TotalPrice)
		
		// total product
		totalProduct = totalProduct + data.Quantity
	}

	// Check grand total with balance amount
	if amount.Amount < int64(grandTotal){
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Not Enough Balance",
		})
		return 
	}

	// Insert Transaction
	transaction.UserId = int(idMidleware)
	transaction.Status = "Pending"
	transaction.TotalPayment = int64(grandTotal)
	transaction.TransNumber = fmt.Sprintf("OrderNumber+ %d", rand.Int())
	transactionId,err := transaction.Create(&transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Something went wrong with transaction",
		})
		return 
	}

	// Insert to history
	for i,_ := range history {
		history[i].TransId = transactionId
	}


	err = history[0].CreateHistory(history)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Something went wrong",
		})
		return 
	}
	// Reduce Balance User
	err = balance.ReduceBalance(history[0].UserId,int64(grandTotal),amount.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Something went wrong with history",
		})
		return 
	}

	snapMidtrans := payment.CreateSnap(transactionId,int64(grandTotal))

	c.JSON(http.StatusOK,gin.H{
		"messaga": "Success buy Product",
		"data": transaction,
		"paymentLink": snapMidtrans,
	})
}

// Get User History
func (user *User) CheckHistory(c *gin.Context){
	history := history.History{}
	idMidleware := c.MustGet("id").(float64)


	res := history.GetHistoryByUserId(int(idMidleware))
	if res == nil {
		c.JSON(http.StatusNotFound,gin.H{
			"message": "History Not Found",
			"id": idMidleware,
		})
		return
	}

	c.JSON(http.StatusBadRequest,gin.H{
		"message": "Success",
		"data": res,
	})
	
}

// Check Balance
func (user *User) CheckBalance(c *gin.Context){
	balance := balance.Balance{}
	idMidleware := c.MustGet("id").(float64)

	res,row := balance.CheckBalanceByUserId(int(idMidleware))
	if row == 0 {
		c.JSON(http.StatusNotFound,gin.H{
			"message": "Balance Not Found",
		})
		return
	}

	c.JSON(http.StatusNotFound,gin.H{
		"message": "Success Get Balance",
		"data": res,
	})
	
}

// Delete User
func (user *User) DeleteUser(c *gin.Context){
	history := history.History{}
	balance := balance.Balance{}

	params := c.Param("id")
	id,err := strconv.Atoi(params)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "ID Not found",
		})
		return
	}
	
	// Delete History
	err = history.DeleteByUserId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Cant Delete History",
		})
		return
	}

	// Delete Balance
	err = balance.DeleteByUserId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Cant Delete Balance",
		})
		return
	}

	// Delete User
	err = user.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Cant Delete User",
		})
		return
	}

	// Succes Response

	c.JSON(http.StatusInternalServerError,gin.H{
		"message": "Success Delete User",
	})
}


