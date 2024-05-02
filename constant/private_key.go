package constant

import "os"

// Private Key for JWT
func PrivateKeyJWT() string {
	return os.Getenv("PRIVATE_KEY_JWT")
}

func ImageApiKey() string {
	return os.Getenv("PRIVATE_KEY_PEXELS")
}
