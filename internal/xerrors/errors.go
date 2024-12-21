package xerrors

import "errors"

var NotImplementedError = errors.New("this functionality has not been implemented")

var SessionTokenMissingError = errors.New("session token is required")
