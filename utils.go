package tcpx

import (
	"encoding/json"
	"strings"
)

type H map[string]interface{}

func Debug(src interface{}) string {
	buf, _ := json.MarshalIndent(src, "  ", "  ")
	return string(buf)
}

// Whether s in arr
// Support %%
func In(s string, arr []string) bool {
	for _, v := range arr {
		if strings.Contains(v, "%"){
			if strings.HasPrefix(v, "%") && strings.HasSuffix(v, "%") {
				if strings.Contains(s, string(v[1:len(v)-1])) {
					return true
				}
			} else if strings.HasPrefix(v, "%") {
				if strings.HasSuffix(s, string(v[1:])) {
					return true
				}
			} else if strings.HasSuffix(v, "%") {
				if strings.HasPrefix(s, string(v[:len(v)-1])) {
					return true
				}
			}
		}else {
			if v == s {
				return true
			}
		}
	}
	return false
}
