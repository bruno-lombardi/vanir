package helpers

import (
	"errors"

	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
)

var (
	uni *ut.UniversalTranslator
)

type CustomValidator struct {
	Validator  *validator.Validate
	Translator ut.Translator
}

func (cv *CustomValidator) Validate(i interface{}) (err error) {
	if err = cv.Validator.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			return errors.New(e.Translate(cv.Translator))
		}
	}
	return nil
}

func NewCustomValidator() *CustomValidator {
	pt := pt.New()
	uni = ut.New(pt, pt)
	translator, _ := uni.GetTranslator("pt")

	validate := validator.New()
	// Required translation
	validate.RegisterTranslation("required", translator, func(ut ut.Translator) error {
		return ut.Add("required", "O campo {0} é obrigatório", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
	// Greater than or equals
	validate.RegisterTranslation("gte", translator, func(ut ut.Translator) error {
		return ut.Add("gte", "O campo {0} deve ser maior ou igual a {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("gte", fe.Field(), fe.Param())
		return t
	})

	// Less than or equals
	validate.RegisterTranslation("lte", translator, func(ut ut.Translator) error {
		return ut.Add("lte", "O campo {0} deve ser menor ou igual a {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("lte", fe.Field(), fe.Param())
		return t
	})

	// String max length
	validate.RegisterTranslation("max", translator, func(ut ut.Translator) error {
		return ut.Add("max", "O campo {0} deve possuir até {1} caracteres", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field(), fe.Param())
		return t
	})

	// String min length
	validate.RegisterTranslation("min", translator, func(ut ut.Translator) error {
		return ut.Add("min", "O campo {0} deve possuir no mínimo {1} caracteres", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field(), fe.Param())
		return t
	})

	// Email
	validate.RegisterTranslation("email", translator, func(ut ut.Translator) error {
		return ut.Add("email", "O valor {0} é um email inválido", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		value, ok := fe.Value().(string)
		var t string

		if !ok {
			t, _ = ut.T("email", fe.Field())
		}
		t, _ = ut.T("email", value)
		return t
	})

	return &CustomValidator{Validator: validate, Translator: translator}
}
