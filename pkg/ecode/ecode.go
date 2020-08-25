package ecode

import (
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/pkg/errors"
)

var (
	_messages atomic.Value         // NOTE: stored map[string]map[int]string
	_codes    = map[int]struct{}{} // register codes.
)

// Register register ecode message map.
func Register(cm map[int]string) {
	_messages.Store(cm)
}

// New new a ecode.ECodes by int value.
// NOTE: ecode must unique in global, the New will check repeat and then panic.
func New(e int) ECode {
	if e <= 0 {
		panic("business ecode must greater than zero")
	}
	return add(e)
}

func add(e int) ECode {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = struct{}{}
	return Int(e)
}

// ECodes ecode error interface which has a code & message.
type ECodes interface {
	// sometimes Error return ECode in string form
	// NOTE: don't use Error in monitor report even it also work for now
	Error() string
	// ECode get error code.
	Code() int
	// Message get code message.
	Message() string
	//Detail get error detail,it may be nil.
	Details() []interface{}
	// Equal for compatible.
	// Deprecated: please use ecode.EqualError.
	Equal(error) bool
}

// A ECode is an int error code spec.
type ECode int

func (e ECode) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

// ECode return error code
func (e ECode) Code() int { return int(e) }

// Message return error message
func (e ECode) Message() string {
	if cm, ok := _messages.Load().(map[int]string); ok {
		if msg, ok := cm[e.Code()]; ok {
			return msg
		}
	}
	return e.Error()
}

// Details return details.
func (e ECode) Details() []interface{} { return nil }

// Equal for compatible.
// Deprecated: please use ecode.EqualError.
func (e ECode) Equal(err error) bool { return EqualError(e, err) }

// Int parse code int to error.
func Int(i int) ECode { return ECode(i) }

// String parse code string to error.
func String(e string) ECode {
	if e == "" {
		return OK
	}
	// try error string
	i, err := strconv.Atoi(e)
	if err != nil {
		return InternalServerError
	}
	return ECode(i)
}

// Cause cause from error to ecode.
func Cause(e error) ECodes {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(ECodes)
	if ok {
		return ec
	}
	return String(e.Error())
}

// Equal equal a and b by code int.
func Equal(a, b ECodes) bool {
	if a == nil {
		a = OK
	}
	if b == nil {
		b = OK
	}
	return a.Code() == b.Code()
}

// EqualError equal error
func EqualError(code ECodes, err error) bool {
	return Cause(err).Code() == code.Code()
}
