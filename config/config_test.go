package config

import "testing"

func TestReplacePlaceHolder(t *testing.T) {
	content, err := GetConfigFromFile("./app.conf")
	if err != nil {
		t.Error(err)
		return
	}

	ret := ReplacePlaceHolder(string(content), func(input []string) map[string]string {
		ret := make(map[string]string)
		for _, v := range input {
			ret[v] = v
		}
		return ret
	})

	expected := `{
    "taobao_api":{
        "app_key":"27755801",
        "app_secret":"order-taobao_secret"
    },
    "config":{
        "tests":[
            "name",
            "password"
        ]
    }
}`
	if ret != expected {
		t.Errorf("expected %v but got %v", expected, ret)
	}
}
