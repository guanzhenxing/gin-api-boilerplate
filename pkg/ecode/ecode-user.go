package ecode

var (
	ErrBind       = add(20001)
	ErrDelete     = add(20002)
	ErrValidation = add(20003)
	ErrDatabase   = add(20004)
	ErrToken      = add(20005)

	// user errors
	ErrEncrypt           = add(20101)
	ErrUserNotFound      = add(20102)
	ErrTokenInvalid      = add(20103)
	ErrPasswordIncorrect = add(20104)
)
