package errors

// Is reports that err.Kind and kind are equal or not
func IsKind(err error, kind Kind) bool {
	e := Match(err, kind)
	if e != nil {
		return true
	}
	return false
}
