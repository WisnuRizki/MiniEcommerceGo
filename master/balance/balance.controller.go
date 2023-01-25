package balance

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"miniecommerce.wisnu.net/master/admin"
)

func (balance *Balance) AddBalanceUser(c *gin.Context){
	admin := admin.Admin{}
	idMidleware := c.MustGet("id").(float64)

	// isAdmin Exist
	isExist := admin.CheckDataById(uint(idMidleware))
	if isExist != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Admin Not Found",
		})
		return
	}

	err := c.BindJSON(&balance)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Error Json",
		})
		return
	}
	
	res,row := balance.CheckBalanceByUserId(balance.UserId)
	if row == 0 {
		c.JSON(http.StatusNotFound,gin.H{
			"message": "Balance Not Found",
		})
		return
	}

	err = balance.AddBalance(res.UserId,balance.Amount,res.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Balance Not Found",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message": "Success",
		"data": balance,
	})
}

// Get ALl Balance

func (balance *Balance) GetAllBalance(c *gin.Context){
	res := balance.GetAll()
	if res == nil {
		c.JSON(http.StatusOK,gin.H{
			"message": "Failed To Get All Data Balance",
		})

		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message": "Success",
		"data":  res,
	})
}