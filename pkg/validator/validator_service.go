package validator

import universalTranslator "github.com/go-playground/universal-translator"

type Service interface {
	ValidateStruct(targetValidatorStruct interface{}) error
	ValidateVar(targetValidatorStruct interface{}, validatorTags string) error
	ParseValidationError(validationError error)
	GetEngTranslator() universalTranslator.Translator
}
