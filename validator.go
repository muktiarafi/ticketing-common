package common

import (
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
)

type ValidationError struct {
	Errors []string
}

func (e *ValidationError) Error() string {
	var msg strings.Builder

	for _, v := range e.Errors {
		fmt.Fprintf(&msg, "%s. ", v)
	}

	return msg.String()
}

type ErrValidationTranslator struct {
	EnTranslator ut.Translator
	IdTranslator ut.Translator
}

func NewDefaultTranslator(v *validator.Validate) *ErrValidationTranslator {
	english := en.New()
	uni := ut.New(english, english)
	enTrans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(v, enTrans)

	id := id.New()
	uni = ut.New(id, id)
	idTrans, _ := uni.GetTranslator("id")
	_ = idTranslations.RegisterDefaultTranslations(v, idTrans)

	return &ErrValidationTranslator{
		EnTranslator: enTrans,
		IdTranslator: idTrans,
	}
}

func translateError(err error, trans ut.Translator) (errs []string) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Sprint(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}

type CustomValidator struct {
	Validator  *validator.Validate
	Translator *ErrValidationTranslator
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {

		validationError := &ValidationError{
			// default is English translate and can be changed to Indonesia Translate
			Errors: translateError(err, cv.Translator.EnTranslator),
		}

		return validationError
	}
	return nil
}
