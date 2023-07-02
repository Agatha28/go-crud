package libraries

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type validation struct {
	validate *validator.Validate
	trans    ut.Translator
}
