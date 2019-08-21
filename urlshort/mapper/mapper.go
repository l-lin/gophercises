package mapper

import (
	"io/ioutil"
)

func fromFile(file string, unmarshal func([]byte, interface{}) error) (map[string]string, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var out []map[string]string
	err = unmarshal(content, &out)
	if err != nil {
		return nil, err
	}
	result := make(map[string]string)
	for _, m := range out {
		result[m["path"]] = m["url"]
	}

	return result, nil
}
