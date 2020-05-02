package main

import (
	"log"
	"net/http"

	"github.com/daliyo/holidayapp/holiday"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/:date", func(ctx *gin.Context) {
		date := ctx.Param("date")
		data, err := holiday.Query(date)
		if err != nil {
			ctx.String(http.StatusBadRequest, "%s", err.Error())
		} else {
			ctx.JSON(http.StatusOK, data)
		}
	})
	log.Fatal(r.Run(":3000"))
}
