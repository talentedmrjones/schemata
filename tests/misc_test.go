package schemata_test

func cmpError(a, b error) bool {
	if a == nil {
		return b == nil
	}
	if b == nil {
		return a == nil
	}
	return a.Error() == b.Error()
}
