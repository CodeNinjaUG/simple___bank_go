package api

import (
	"github.com/codeninjaug/simple_bank/db/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		//check if currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
