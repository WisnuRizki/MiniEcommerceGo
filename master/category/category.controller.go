package category

import (
	"net/http"
	"strings"
	"strconv"

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

// Deleta Category
func (category *Category) DeleteCategory(c *gin.Context){
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

	params := c.Param("id")
	id,err := strconv.Atoi(params)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Bad Request",
		})
		return
	}

	err = category.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Something went wrong",
		})
	}

	c.JSON(http.StatusInternalServerError,gin.H{
		"message": "Success delete category",
	})

}

// Get All Category

func (category *Category) GetAllCategory(c *gin.Context){
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

	res := category.GetAll()
	if res == nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Category Not Found",
		})
		return
	}

	c.JSON(http.StatusInternalServerError,gin.H{
		"message": "Success",
		"data": res,
	})
}

// Get Category By Id
func (category *Category) GetCategoryById(c *gin.Context){
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

	params := c.Param("id")
	id,err := strconv.Atoi(params)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Bad Request",
		})
		return
	}

	res := category.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Something went wrong",
		})
	}

	c.JSON(http.StatusInternalServerError,gin.H{
		"message": "Success",
		"data": res,
	})

}

// Update

func (category *Category) UpdateCategoryById(c *gin.Context){
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

	params := c.Param("id")
	id,err := strconv.Atoi(params)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Bad Request",
		})
		return
	}

	err = c.BindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Bad Request",
		})
		return
	}
	category.Name = strings.ToUpper(category.Name)

	err = category.Update(id,category)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Something went wrong",
		})
	}

	c.JSON(http.StatusInternalServerError,gin.H{
		"message": "Success Update",
	})

}
