package govalidator

import (
	"fmt"
	"regexp"
	"time"
)

type Validator interface {
	Validate(v interface{}) bool
}

type Validation struct{}

func (v *Validation) Validate(validator Validator, data interface{}) bool {
	return validator.Validate(data)
}

// Min Validation
// //@Data.Validate.Min(1)
type Min struct {
	Min int
}

func (d Min) Validate(v interface{}) bool {
	if num, ok := v.(int); ok {
		return num >= d.Min
	}
	return false
}

// Max validation
// //@Data.Validate.Max(20)
// maxAge int
type Max struct {
	Max int
}

func (d Max) Validate(v interface{}) bool {
	if num, ok := v.(int); ok {
		return num <= d.Max
	}
	return false
}

// Len Validation for string
// //@Data.Validate.Len(20)
// max20CharStr string
type Len struct {
	Len int
}

func (d Len) Validate(v interface{}) bool {
	if str, ok := v.(string); ok {
		return len(str) == d.Len
	}
	return false
}

// //@Data.Validate.Required
type Required struct{}

func (d Required) Validate(v interface{}) bool {
	if v == nil {
		return false
	}
	switch v.(type) {
	case string:
		return len(v.(string)) > 0
	case bool:
		return v.(bool)
	case int:
		return v.(int) != 0
	case time.Time:
		return !v.(time.Time).IsZero()
	case *time.Time:
		return !v.(*time.Time).IsZero()
	default:
		fmt.Println(fmt.Sprintf("Not implemented validation %T", v))
		return false
	}

}

// //@Data.Validate.Match
// email string
type Match struct {
	Regexp *regexp.Regexp
}

func NewMatch(match string) Match {
	return Match{regexp.MustCompile(match)}
}

func (d Match) Validate(v interface{}) bool {
	if str, ok := v.(string); ok {
		return d.Regexp.MatchString(str)
	}
	return false
}

// //@Data.Validate.Match
// email string
type Email struct {
	Match
}

func NewEmail() Email {
	// find a good email regex
	//http://stackoverflow.com/questions/13087755/can-anyone-tell-me-why-this-c-sharp-email-validation-regular-expression-regex
	var emailRegex = `(\w*[0-9a-zA-Z])*@`
	return Email{NewMatch(emailRegex)}
}
