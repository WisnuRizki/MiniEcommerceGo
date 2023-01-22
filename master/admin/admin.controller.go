package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"miniecommerce.wisnu.net/helpers"
)

func (admin *Admin) RegisterAdmin(c *gin.Context){
	err := c.BindJSON(&admin)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Error with Json data",
		})
		return
	}

	// Check user already exist or not
	_,row := admin.CheckDataByUsername(admin.Username)
	if row > 0 {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "User Already Exist",
		})
		return
	}

	// hash password
	hashPassword,err := bcrypt.GenerateFromPassword([]byte(admin.Password),10)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Something Went Wrong",
		})
		return
	}

	admin.Password = string(hashPassword)
	admin.ID = 0

	err = admin.CreateAdmin(admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Internal Server error",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message": "Success",
		"data": admin,
	})
}

func (admin *Admin) LoginAdmin(c *gin.Context){
	err := c.BindJSON(&admin)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Error JSON",
		})
		return
	}

	if admin.Username == "" || admin.Password == "" {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Need Email Or Password",
		})
		return
	}

	res,row := admin.CheckDataByUsername(admin.Username)
	if row == 0 {
		c.JSON(http.StatusNotFound,gin.H{
			"message": "User Not Found",
		})
		return
	}

	jwtString,pl := helpers.GenerateJWT(res.ID,res.Username)
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