package validator

//TODO:实现统一参数验证

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"regexp"
	ErrHandle "src/middleware/errHandle"
	"strconv"
)

// 验证手机号
func mobileNumber(fl validator.FieldLevel) bool {
	pattern := `^[1]([3-9])[0-9]{9}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(fl.Field().String())
}

//注册自定义验证器
func RegisterCustomValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("c_mobile", mobileNumber)
	}
}

// 绑定模型获取验证错误的方法
func GetError(error interface{}) []string {
	var retErrs []string
	switch errs := error.(type) {
	case validator.ValidationErrors:
		for _, err := range errs {
			if err.Tag() == "required" {
				retErrs = append(retErrs, fmt.Sprint(err.Field(), "不能为空"))
			} else if err.Tag() == "email" {
				retErrs = append(retErrs, fmt.Sprint(err.Field(), "地址不合法"))
			} else if err.Tag() == "c_mobile" {
				retErrs = append(retErrs, fmt.Sprint(err.Field(), "手机号不合法"))
			}
		}
	case *strconv.NumError:
		retErrs = append(retErrs, "参数解析错误，非法字符")
	default:
		fmt.Println(errs, reflect.TypeOf(errs))
	}
	return retErrs
}

//author: zfz
//createAt: 2020/2/24
//description: uri, query, form 三者是指针,用于验证uri,query,form的参数，不存在可设置为nil
func ValidateAll(ctx *gin.Context, uri, query, form interface{}) []string {
	errors := make([]string, 0)
	if uri != nil {
		if err := ctx.ShouldBindUri(uri); err != nil {
			errors = append(errors, GetError(err)...)
		}
	}
	if query != nil {
		if err := ctx.ShouldBindQuery(query); err != nil {
			errors = append(errors, GetError(err)...)
		}
	}
	if form != nil {
		if err := ctx.ShouldBind(form); err != nil {
			errors = append(errors, GetError(err)...)
		}
	}
	return errors
}

//用于在controller中统一调用
func Validate(ctx *gin.Context, uri, query, form interface{}) {
	if errors := ValidateAll(ctx, uri, query, form); len(errors) != 0 {
		panic(ErrHandle.NewError(200, 0, errors))
	}
}
