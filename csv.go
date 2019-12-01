package middlewares

import (
	"bytes"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
)

func csvFromJSON(b []byte) (result []byte, err error) {
	var data []map[string]interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return
	}
	if len(data) == 0 {
		return
	}
	keys := make([]string, 0, len(data[0]))
	for k := range data[0] {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var strOut bytes.Buffer
	line := strings.Join(keys, ";") + "\n"
	strOut.WriteString(line)
	for _, v := range data {
		values := decodeLine(keys, v)
		line := strings.Join(values, ";") + "\n"
		strOut.WriteString(line)
	}
	return strOut.Bytes(), nil
}

func decodeLine(keys []string, data map[string]interface{}) []string {
	values := make([]string, 0, len(data))
	for _, k := range keys {
		v := data[k]
		switch vv := v.(type) {
		case map[string]interface{}:
			values = append(values, "nil")
		case string:
			vv = `"` + strings.ReplaceAll(vv, `"`, " ") + `"`
			values = append(values, vv)
		case float64:
			values = append(values, strconv.FormatFloat(vv, 'f', -1, 64))
		case []interface{}:
			values = append(values, "nil")
		case bool:
			values = append(values, strconv.FormatBool(vv))
		default:
			values = append(values, "nil")
		}
	}
	return values
}
