package domain

func DomainCompliant(domain string) bool {
	switch domain {
	case "EA":
		return true
	case "MFG":
		return true
	case "ENG":
		return true
	default:
		return false
	}
}
