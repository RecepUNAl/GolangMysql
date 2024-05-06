package helper

import "encoding/json"

// Fonksiyon, gelen struct verisini interface{}'e dönüştürür.
func StructToInterface(s interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	var i map[string]interface{}
	err = json.Unmarshal(jsonData, &i)
	if err != nil {
		return nil, err
	}
	return i, nil
}
