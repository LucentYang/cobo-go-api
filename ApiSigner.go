package main
type ApiSigner interface {
	Sign(message string) string
}
