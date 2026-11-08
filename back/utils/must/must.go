package utils_must

func Must[O any](v O, err error) O {
	if err != nil {
		panic(err)
	}
	return v
}
