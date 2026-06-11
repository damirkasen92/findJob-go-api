package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/damir/jobfinder/internal/model"
	validatorv10 "github.com/go-playground/validator/v10"
)

var Validate = validatorv10.New()

func ValidateStruct(req interface{}) error {
	err := Validate.Struct(req)
	if err != nil {
		var ve validatorv10.ValidationErrors
		if errors.As(err, &ve) {
			var details []string
			for _, fe := range ve {
				details = append(details,
					fmt.Sprintf("Field '%s' failed on '%s' rule", fe.Field(), fe.Tag()))
			}

			return fmt.Errorf("%w: %s", model.ErrValidation, strings.Join(details, "; "))
		}
		return err
	}
	return nil
}
