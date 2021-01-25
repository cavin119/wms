package response

import (
	"7youo-wms/global"
	"7youo-wms/utils"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

//处理验证失败错误信息中的英文
//错误字段:{英文:中文} 大小写敏感
var requestFieldTrans = map[string]map[string]string {
	"RePassword": {"Password": "密码"},
}

//获取中文验证器
func getCnValidation() (*validator.Validate, ut.Translator) {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		label := fld.Tag.Get("label")
		return label
	})
	//注册验证器中文翻译
	transErr := zh_trans.RegisterDefaultTranslations(validate, trans)
	if transErr != nil {
		global.G_LOG.Error("request validate translation error:", zap.Any("err", transErr))
	}
	return validate, trans
}

//@param data 请求数据
//@return []interface{}, bool,验证成功 bool = true, 验证失败bool = false,并返回验证失败信息
func RequestValidate(data interface{}) ([]interface{}, bool) {
	validate, trans := getCnValidation()
	//验证数据
	err := validate.Struct(data)
	if err != nil {
		//错误array
		errs := err.(validator.ValidationErrors)
		errsMap := make([]interface{}, 0)//用于格式化post参数错误，key=>error
		//格式化错误信息输出
		for _, validationErr := range errs {
			//request错误字段
			filed := validationErr.StructField()
			//中文化错误
			msg := validationErr.Translate(trans)
			//如果存在错误中没法翻译的字段
			if _, ok := requestFieldTrans[filed]; ok {
				for oldStr, newStr := range requestFieldTrans[filed] {
					msg = strings.Replace(msg, oldStr, newStr, -1)
				}
			}
			//把字段错误信息加到map里面
			//大写转成下划线，统一输出标准
			errsMap = append(errsMap, map[string]string{
				utils.UnderScoreName(filed): msg,
			})
		}
		return errsMap, false
	}
	return nil, true
}
