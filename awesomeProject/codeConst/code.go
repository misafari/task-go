package codeConst

import "fmt"

const (
	Data            = "data"
	Error           = "error"
	TokenIsRequired = "token is required."
	TokenNotValid   = "token is not valid."
	//TokenExpired = "token is expired."
)

func EntityNotFound(entityName string) string {
	return fmt.Sprintf("%s not found", entityName)
}
