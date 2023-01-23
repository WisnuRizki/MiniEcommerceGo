package product

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"miniecommerce.wisnu.net/master/seller"
	
)

func (product Product) Create(c *gin.Context){
	products := []Product{}
	seller := seller.Seller{}

	idMidleware := c.MustGet("id").(float64)

	_,row := seller.CheckDataSellerExistById(uint(idMidleware))
	if row == 0 {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "User Not Found",
		})
		return 
	} 


	err := c.BindJSON(&products)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Error with Json data",
		})

		fmt.Printf("%s",err)
		return
	}

	for i := 0; i < len(products); i++ {
		products[i].SellerId = int(idMidleware)
	}
	
	err = product.CreateProduct(products)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "Internal Server Error",
		})

		fmt.Printf("%s",err)
		return
	}


	c.JSON(http.StatusOK,gin.H{
		"message": "Success",
		"data": products,
	})

}

func (product Product) GetAllProductBySeller(c *gin.Context){
	idMidleware := c.MustGet("id").(float64)

	res,err := product.GetAllProductSeller(int(idMidleware))
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "Product Not Found",
		})
		return 
	}

	c.JSON(http.StatusBadRequest,gin.H{
		"message": "Success Get All Product",
		"data": res,
	})
}
