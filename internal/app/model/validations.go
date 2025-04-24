package model

import validation "github.com/go-ozzo/ozzo-validation"

func requiredIf(conf bool) validation.RuleFunc {
	return func(value interface{}) error {
		if conf {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}
