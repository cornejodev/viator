package errs

// MissingField is an error type that can be used when
// validating input fields that do not have a value, but should
func MissingField(f string) string {
	return f + " is required"
}

// InputUnwanted is an error type that can be used when
// validating input fields that have a value, but should should not
func InputUnwanted(in string) string {
	return in + " has a value, but should be nil"
}
