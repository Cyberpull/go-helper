package gotk

func one[T any](def T, attr []T) T {
	if len(attr) > 0 {
		return attr[0]
	}

	return def
}

func oneOfAny[T any](def T, attr []any) T {
	if len(attr) > 0 {
		return attr[0].(T)
	}

	return def
}
