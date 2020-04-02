package config

import (
	"fmt"
	"regexp"
	"strings"
)

var placeHolder = regexp.MustCompile(`\[\[\.(.*)+?\]\]`)

type placeHolderReplaceFunc func([]string) map[string]string

func ReplacePlaceHolder(content string, replacerFn placeHolderReplaceFunc) string {
	keys := placeHolder.FindAllStringSubmatch(content, -1)

	cache := make(map[string]struct{})
	oldKeys := make([]string, 0, len(keys))
	for _, v := range keys {
		if len(v) < 2 {
			continue
		}

		if _, ok := cache[v[1]]; ok {
			continue
		}

		oldKeys = append(oldKeys, v[1])
		cache[v[1]] = struct{}{}
	}

	newKeys := replacerFn(oldKeys)
	oldNewParis := make([]string, 0, 2*len(oldKeys))
	for _, old := range oldKeys {
		oldNewParis = append(oldNewParis, fmt.Sprintf("[[.%v]]", old), newKeys[old])
	}

	replacer := strings.NewReplacer(oldNewParis...)
	return replacer.Replace(content)
}
