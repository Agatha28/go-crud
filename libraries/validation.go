package libraries

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validation struct {
	validate *validator.Validate
	trans    ut.Translator
}

func NewValidation() *Validation {
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	return &Validation{
		validate: validate,
		trans:    trans,
	}
}
