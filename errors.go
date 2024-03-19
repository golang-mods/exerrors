package exerrors

func Flatten(errs ...error) []error {
	var result []error

	for _, err := range errs {
		if unwrapper, ok := err.(interface{ Unwrap() []error }); ok {
			result = append(result, Flatten(unwrapper.Unwrap()...)...)
		} else if err != nil {
			result = append(result, err)
		}
	}

	return result
}
