package helper

func Rerr[T any](data T, err error) T {
	if err != nil {
		panic(err)
	}

	return data
}

func Err(err error) {
	if err != nil {
		panic(err)
	}
}
