package testcontrollers

import (
	// "bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/astaxie/beego"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"

	_ "app/server_side/routers"
)

var jsonDirectory = baseDirectory + "/tests/test_controllers/test_json/"

func camelToSnake(camelCaseName string) (snakeCaseName string) {
	slice := strings.Split(camelCaseName, "")
	upperAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, c := range slice {
		if strings.Contains(upperAlphabet, c) {
			c = "_" + strings.ToLower(c)
		}
		snakeCaseName = snakeCaseName + c
	}
	return
}

// Test Get
func testGet(t *testing.T, controllerNameByCamelCase string, id int64, expectJSON string) {
	tableName := camelToSnake(controllerNameByCamelCase)
	response := get(t, tableName, id)
	assert.Equal(t, 200, response.StatusCode)
	bodyByte, _ := ioutil.ReadAll(response.Body)
	actualJSON := string(bodyByte)
	compareJSONs(t, expectJSON, actualJSON, id)
}

// 指定したレコードのGet
func get(t *testing.T, tableName string, id int64) (response *http.Response) {
	url := "/api/v1" + tableName + "/" + strconv.Itoa(int(id))
	httpRequest := httptest.NewRequest("GET", url, nil)
	recorder := httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(recorder, httpRequest)
	response = recorder.Result()

	return
}

// JSONファイルとの比較
func compareJSONs(t *testing.T, expectJSON string, actualJSON string, id int64) {
	var testCase []string

	testCase = jsonKeysToStringSlice(expectJSON, "", false)

	if id != -1 {
		assert.Equal(t, strconv.Itoa(int(id)), gjson.Get(actualJSON, "id").String(), "id")
	}

	for _, path := range testCase {
		CompareJSONElement(t, expectJSON, path, actualJSON, path)
	}
}

// JSONの各キー名を、ネストを考慮してstringのスライスにする
func jsonKeysToStringSlice(jsonStr string, parentKeyName string, isArray bool) (slice []string) {
	var jsonArray []gjson.Result
	var jsonMap map[string]gjson.Result

	// 配列の場合
	if isArray {
		jsonArray = gjson.Parse(jsonStr).Array()

		if parentKeyName != "" {
			parentKeyName = parentKeyName + ".#"
		}

		slice = append(slice, jsonKeysToStringSlice(jsonArray[0].String(), parentKeyName, false)...)
	} else {
		jsonMap = gjson.Parse(jsonStr).Map()
		for key, value := range jsonMap {
			if key == "createAt" || key == "updateAt" || key == "deleteAt" {
				continue
			}
			if parentKeyName != "" {
				key = parentKeyName + "." + key
			}
			if gjson.Valid(value.String()) {
				_, err := strconv.Atoi(value.String())
				if err != nil {
					if value.IsArray() {
						slice = append(slice, jsonKeysToStringSlice(value.String(), key, true)...)
					} else {
						slice = append(slice, jsonKeysToStringSlice(value.String(), key, false)...)
					}
				} else {
					slice = append(slice, key)
				}
			} else {
				slice = append(slice, key)
			}
		}
	}
	return
}

// JSONの比較
func CompareJSONElement(t *testing.T, expectJSON string, expectPath string, resultJSON string, resultPath string) {
	expectStr := extractDate(gjson.Get(expectJSON, expectPath).String())
	resultStr := extractDate(gjson.Get(resultJSON, resultPath).String())
	assert.Equal(t, expectStr, resultStr, expectPath)
}

func extractDate(str string) (ret string) {
	runes := []rune(str)
	return string(runes[0:10])
}
