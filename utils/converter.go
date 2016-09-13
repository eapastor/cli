package utils

// BToS - Convert bool type to string (yes|no)
func BToS(data bool) string {

	if data {
		return "yes"
	}
	return "no"

}
