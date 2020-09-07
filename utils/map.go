package utils

import (
	"sort"
	"strconv"
)

func ParseMap(args map[string]interface{}) string {
	var (
		stringA string
		keys    = make([]string, 0)
	)

	for k, _ := range args {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		switch v := args[key].(type) {
		case string:
			stringA += "&" + key + "=" + v

		case int:
			stringA += "&" + key + "=" + strconv.Itoa(v)

		case int32:
			stringA += "&" + key + "=" + strconv.Itoa(int(v))

		case int64:
			stringA += "&" + key + "=" + strconv.Itoa(int(v))

		default:
			return ""
		}
	}

	return stringA[1:]
}

