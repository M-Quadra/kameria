package kameria

import "encoding/json"

//JSONStrToMap json to map[string]interface{}
func JSONStrToMap(jsonStr string) (map[string]interface{}, error) {
	var opt map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &opt)
	if err != nil {
		return nil, err
	}

	return opt, nil
}
