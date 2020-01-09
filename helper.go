package PKGBUILD

import "strings"

func wrapString(source string, prefix string, suffix string) string {
	return prefix + source + suffix
}

func wrapStrings(source []string, join string, prefix string, suffix string) string {
	var wrapped []string
	for _, s := range source {
		wrapped = append(wrapped, wrapString(s, prefix, suffix))
	}
	return strings.Join(wrapped, join)
}
