package seller


import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	_"miniecommerce.wisnu.net/helpers"

)

func (seller *Seller) RegisterSeller(c *gin.Context){
	err := c.BindJSON(&seller)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Bad Request",
		})
		return
	}
	
	_,row := seller.CheckDataSellerExists(seller.Email,seller.NoHp)
	if row > 0 {
		c.JSON(http.StatusFound,gin.H{
			"message": "Seller already exist",
		})
		return
	}

	hashPassword,err := bcrypt.GenerateFromPassword([]byte(seller.Password),10)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Something Went Wrong",
		})
		return
	}


	seller.Password = string(hashPassword)
	err = seller.CreateSeller(seller)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message": "Success",
		"data": seller,
	})


}