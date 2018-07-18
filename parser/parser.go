package parser

type Config struct {
	Url         string
	HouseNumber string
	Postcode    string
}
type Parser interface {
	Parse(Config Config) []string
}
