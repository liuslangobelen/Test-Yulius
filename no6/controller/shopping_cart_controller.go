package controller

import (
	"net/http"

	"tog.test/no6/config"
	"tog.test/no6/dto"
	"tog.test/no6/model"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

func AddCart(c *gin.Context) {

	userId := 1

	var reqBody dto.ReqAddCart
	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	//connect db
	db := config.Connect()
	defer config.Close(db)
	//end connect db

	var shoppingCart model.ShoppingCart

	db.Model(&model.ShoppingCart{}).
		Where("user_id = ?", userId).
		Where("code_product = ?", reqBody.CodeProduct).
		First(&shoppingCart)

	if shoppingCart.ID != 0 {
		shoppingCart.Quantity += reqBody.Quantity
	} else {
		shoppingCart.NameProduct = reqBody.NameProduct
		shoppingCart.CodeProduct = reqBody.CodeProduct
		shoppingCart.Quantity = reqBody.Quantity
		shoppingCart.UserId = uint(userId)
	}
	db.Save(&shoppingCart)

	c.JSON(http.StatusOK, "success add cart")
}

func ShowCart(c *gin.Context) {

	userId := 1
	nameProduct := c.DefaultQuery("namaProduct", "")
	quantity := c.DefaultQuery("quantity", "")

	var response []dto.ResShoppingCart
	var shoppingCart []model.ShoppingCart

	//connect db
	db := config.Connect()
	defer config.Close(db)
	//end connect db

	querry := db.Model(&model.ShoppingCart{}).
		Where("user_id = ?", userId)

	if nameProduct != "" {
		querry.Where("name_product LIKE ?", "%"+nameProduct+"%")
	}
	if quantity != "" {
		// qts ,_ :=strconv.Atoi(quantity)
		querry.Where("quantity = ?", quantity)
	}

	querry.
		Find(&shoppingCart)

	for i := 0; i < len(shoppingCart); i++ {
		response = append(response, dto.ResShoppingCart{
			NameProduct: shoppingCart[i].NameProduct,
			CodeProduct: shoppingCart[i].CodeProduct,
			Quantity:    shoppingCart[i].Quantity,
			ID:          shoppingCart[i].ID,
		})
		deepcopier.Copy(&shoppingCart).To(response)

	}

	c.JSON(http.StatusOK, response)
}

func DeleteCart(c *gin.Context) {

	userId := 1

	code := c.Param("code")

	//connect db
	db := config.Connect()
	defer config.Close(db)
	//end connect db

	var shoppingCart model.ShoppingCart

	db.Model(&model.ShoppingCart{}).
		Where("user_id = ?", userId).
		Where("code_product = ?", code).
		Delete(&shoppingCart)

	if shoppingCart.ID != 0 {
		db.Delete(&shoppingCart)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "data not found",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, "success delete cart")
}
