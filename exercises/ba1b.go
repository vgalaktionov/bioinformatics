package exercises

func foobar(a []byte, b []byte) string {
	return ""
}

func (ex Exercise) BA1B(params [][]byte) string {
	a := params[0]
	b := params[1]

	result := foobar(a, b)

	return result
}
