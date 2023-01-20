package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"miniecommerce.wisnu.net/helpers"
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

	c.JSON(http.StatusOK,gin.H{
		"message": "Login Success",
		"token": jwtString,
	})

}