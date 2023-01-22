package category

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"miniecommerce.wisnu.net/master/admin"
)

func (category *Category) CreateCategory(c *gin.Context){
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

	
	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Error Json",
		})
		return
	}

	category.Name = strings.ToUpper(category.Name)

	err = category.IsCategoryExist(category.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Category Exist",
		})
		return
	}

	err = category.Create(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusBadRequest,gin.H{
		"message": "Success Create Category",
	})

}