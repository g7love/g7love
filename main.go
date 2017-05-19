package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"g7love/controller"
	"g7love/database"
	"g7love/model"
	"gopkg.in/bluesuncorp/validator.v5"
	"io"
	"net/http"
	"os"
	"unicode"
)

const defaultPort = "8080"

var (
	msgInvalidJSON     = "Invalid JSON format"
	msgInvalidJSONType = func(e *json.UnmarshalTypeError) string {
		return "Expected " + e.Value + " but given type is " + e.Type.String() + " in JSON"
	}
	msgValidationFailed = func(e *validator.FieldError) string {
		switch e.Tag {
		case "required":
			return toSnakeCase(e.Field) + ": required"
		case "max":
			return toSnakeCase(e.Field) + ": too_long"
		default:
			return e.Error()
		}
	}
)

func main() {
	//migrate()
	router := gin.Default()

	HomeGroup := router.Group("/home")
	{
		HomeGroup.POST("/loginjudge", controller.Homes.Loginjudge)
	}

	registeredGroup := router.Group("/registered")
	{
		registeredGroup.POST("/provinces", controller.Registered.Provinces)
	}

	dynamicGroup := router.Group("/dynamic")
	{
		dynamicGroup.POST("/getdynamic", controller.Dynamic.Getdynamic)
	}


	http.ListenAndServe(":"+port(), router)
}

func apiHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		errs := make([]string, 0, len(c.Errors))
		for _, e := range c.Errors {
			switch e.Err {
			case io.EOF:
				errs = append(errs, msgInvalidJSON)
			default:
				switch e.Err.(type) {
				case *json.SyntaxError:
					errs = append(errs, msgInvalidJSON)
				case *json.UnmarshalTypeError:
					errs = append(errs, msgInvalidJSONType(e.Err.(*json.UnmarshalTypeError)))
				case *validator.StructErrors:
					for _, fieldErr := range e.Err.(*validator.StructErrors).Flatten() {
						errs = append(errs, msgValidationFailed(fieldErr))
					}
				default:
					errs = append(errs, e.Error())
				}
			}
		}

		if len(c.Errors) > 0 {
			c.JSON(-1, gin.H{"errors": errs}) // -1 == not override the current error code
		}
	}
}

// https://gist.github.com/elwinar/14e1e897fdbe4d3432e1
func toSnakeCase(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	return port
}

func migrate() {
	db := database.GetDB()

	db.AutoMigrate(&model.Registered{}, &model.Todo{})
}
