//Package apperrors handes application errors
package apperrors

import (
	"fmt"
	"time"
)

//Error struct holds the error information
type Error struct {
	Message string
	Date    time.Time
	Code    string
}

//implements the Error interface so we can return an error
func (e *Error) Error() string {
	return fmt.Sprintf("%v: %s", e.Date.Format("2006-01-02 15:04:05"), e.Message)
}

//New is a constructor that creates a new error
func New(message string, code string) error {

	err := &Error{
		Message: message,
		Date:    time.Now(),
		Code:    code,
	}

	return err

}
