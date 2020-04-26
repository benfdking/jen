package url

const (
	a = "a"
	b = "b"
	c = "c"
)

//ReturnVersions returns the available options
func ReturnVersions() []string {
	return []string{a, b, c}
}

// IsVersion returns true if the passed string is one of the versions
func IsVersion(s string) bool {
	vs := ReturnVersions()
	for _, v := range vs{
		if v == s {
			return true
		}
	}
	return false
}