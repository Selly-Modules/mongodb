package mongodb

import "encoding/base64"

func base64DecodeToBytes(text string) []byte {
	s, _ := base64.StdEncoding.DecodeString(text)
	return s
}

func base64DecodeToString(text string) string {
	return string(base64DecodeToBytes(text))
}
