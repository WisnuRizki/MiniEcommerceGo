package seller


import (
	"net/http"
	"miniecommerce.wisnu.net/helpers"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

)

func (seller *Seller) RegisterSeller(c *gin.Context){
	err := c.BindJSON(&seller)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Bad Request",
		})
		return
	}
	
	_,row := seller.CheckDataSellerExists(seller.Email)
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

func (seller *Seller) LoginSeller(c *gin.Context){
	err := c.BindJSON(&seller)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Error JSON",
		})
		return
	}

	if seller.Email == "" || seller.Password == "" {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Need Email Or Password",
		})
		return
	}

	res,row := seller.CheckDataSellerExists(seller.Email)
	if row == 0 {
		c.JSON(http.StatusNotFound,gin.H{
			"message": "Seller Not Found",
		})
		return
	}

	jwtString,pl := helpers.GenerateJWT(res.ID,res.Email)
	if pl != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Internal Server error",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message": "Login Success",
		"token": jwtString,
	})

}

// Get All Produk Seller