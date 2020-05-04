package testcontrollers

import (
	"io/ioutil"
	// "strings"
	_ "app/server_side/routers"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

const User = "user"

func TestGetUser(t *testing.T) {
	pathOfGetDir := jsonDirectory + camelToSnake(User) + "/get/"

	expectJSONDir, dirErr := ioutil.ReadDir(pathOfGetDir)
	if dirErr != nil {
		assert.Fail(t, dirErr.Error())
	}

	for _, expectJSONFile := range expectJSONDir {
		expectJSONByte, eErr := ioutil.ReadFile(pathOfGetDir + expectJSONFile.Name())
		if eErr != nil {
			assert.Fail(t, eErr.Error())
		}

		expectJSON := string(expectJSONByte)

		id := gjson.Get(expectJSON, "id").Int()

		testGet(t, User, id, expectJSON)
	}
}
