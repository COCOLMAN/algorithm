package DefangingAnIPAddress

import "strings"

func defangIPaddr(address string) string {
	sb := strings.Builder{}
	for _, ch := range address {
		if ch == 46 {
			sb.WriteString("[.]")
		} else {
			sb.WriteRune(ch)
		}

	}
	return sb.String()
}
