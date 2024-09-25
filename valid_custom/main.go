package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Booking1 contains binded and validated data.
type Booking1 struct {
	CheckIn  time.Time `form:"check_in" validate:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" validate:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

var bookableDate1 validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}
var validate *validator.Validate

// func main() {
// 	route := gin.Default()
// 	validate = validator.New()
// 	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
// 		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
// 		if name == "-" {
// 			return ""
// 		}
// 		return name
// 	})
// 	//err := validate.Struct(user)
// 	//v, ok := binding.Validator.Engine().(*validator.Validate)
// 	//if ok {
// 	//	v.RegisterValidation("bookabledate", bookableDate1)
// 	//}

// 	route.GET("/bookable", getBookable1)
// 	route.Run(":8085")
// }

func getBookable1(c *gin.Context) {
	var b Booking1
	if err := c.ShouldBind(&b); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking1 dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//$ curl "localhost:8085/bookable?check_in=2118-04-16&check_out=2118-04-17"
//{"message":"Booking1 dates are valid!"}
//
//$ curl "localhost:8085/bookable?check_in=2118-03-10&check_out=2118-03-09"
//{"error":"Key: 'Booking1.CheckOut' Error:Field validation for 'CheckOut' failed on the 'gtfield' tag"}
