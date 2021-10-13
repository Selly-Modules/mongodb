package mongodb

import (
	"regexp"
	"strings"
	"unicode"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// GenerateQuerySearchString ...
func GenerateQuerySearchString(s string) bson.M {
	return bson.M{
		"$regex": bsonx.Regex(NonAccentVietnamese(s), "i"),
	}
}

// NonAccentVietnamese ...
func NonAccentVietnamese(str string) string {
	str = strings.ToLower(str)
	str = replaceStringWithRegex(str, `đ`, "d")
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, str)
	result = replaceStringWithRegex(result, `[^a-zA-Z0-9\s]`, "")

	return result
}

// replaceStringWithRegex ...
func replaceStringWithRegex(src string, regex string, replaceText string) string {
	reg := regexp.MustCompile(regex)
	return reg.ReplaceAllString(src, replaceText)
}
