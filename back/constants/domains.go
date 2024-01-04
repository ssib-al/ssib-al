package constants

var (
	Domain = []string{
		"ssib.al",
		"zirr.al",
		"niae.me",
	}
)

func InsertionPrefix(prefix string) {
	for i, domain := range Domain {
		Domain[i] = prefix + "." + domain
	}
}
