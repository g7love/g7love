package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"g7love/controller"
	"g7love/database"
	"g7love/model"
	"g7love/services"
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
	migrate()
	router := gin.Default()

	HomeGroup := router.Group("/home")
	{
		HomeGroup.POST("/loginjudge",apiHandle("HomeLoginjudge"), controller.Homes.Loginjudge)
	}

	registeredGroup := router.Group("/registered")
	{
		registeredGroup.POST("/provinces",apiHandle("RegisteredProvinces"), controller.Registered.Provinces)
		registeredGroup.POST("/registered",apiHandle("RegisteredRegistered"), controller.Registered.Registered)
	}

	dynamicGroup := router.Group("/dynamic")
	{
		dynamicGroup.POST("/getdynamic",apiHandle("DynamicGetdynamic"), controller.Dynamic.Getdynamic)
		dynamicGroup.POST("/posting",apiHandle("DynamicPosting"), controller.Dynamic.Posting)
	}

	EvaluationGroup := router.Group("/evaluation")
	{
		EvaluationGroup.POST("/evaluation",apiHandle("EvaluationEvaluation"), controller.Evaluation.Evaluation)
	}

	http.ListenAndServe(":"+port(), router)
}

type resultdata struct {
	Code int `json:"code"`
	Status int `json:"status"`
	Data interface{} `json:"data"`
}

func result(data interface{},status int,code int) resultdata {
	result := resultdata{}
	result.Code = code
	result.Status = status
	result.Data = data
	return result
}



/**
 * User: lizhengxinag
 * createtime: 17-05-14 14:04
 * functionRole:验证登录
 */
func apiHandle(authority string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//需要忽略验证的模块
		ignoreValidation := map[string] int{
			"HomeLoginjudge" : 1,
			"RegisteredProvinces" : 1,
			"RegisteredRegistered":1,
			"DynamicGetdynamic":1,
		}
		var userData services.User
		Token :=  c.Param("Token")
		Token = "aaaaa"
		if Token != "" {
			userData.Id = "100177"
		} else if ignoreValidation[authority] == 1 {
			userData.Id = ""
		} else {
			//强制登录
			data :=result("",0,0)
			c.JSON(http.StatusOK, data)
			c.Abort()
			return
		}
		c.Set("user", userData)
		c.Next()
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
	db.AutoMigrate(&model.Dynamic{}, &model.Dynamiclog{},&model.Registered{},
		&model.School{},&model.Thumblog{})
}
