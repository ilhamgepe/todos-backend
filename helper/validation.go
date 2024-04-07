package helper

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)


func GenerateMessage(err error)map[string]string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve){
		messages := make(map[string]string, len(ve))
		for _, v := range ve {
			switch v.Tag(){
			case "email":
				messages[v.Field()] = fmt.Sprintf("%s is not valid email", v.Value())
			case "required":
				messages[v.Field()] = fmt.Sprintf("%s is required", v.Field())
			case "min":
				messages[v.Field()] = fmt.Sprintf("%s must be at least %s characters", v.Field(), v.Param())
			case "max":
				messages[v.Field()] = fmt.Sprintf("%s must be at most %s characters", v.Field(), v.Param())
			}
		}
		return messages
	}
	return map[string]string{"error": err.Error()}
}