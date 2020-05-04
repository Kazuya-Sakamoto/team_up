package testcontrollers

import (
	// "log"

	"os"
	_ "strings"
	"testing"

	"app/server_side/models"
	_ "app/server_side/routers"
	"app/server_side/services"
	_ "path/filepath"

	// "github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/stretchr/testify/assert"
	// "github.com/tidwall/gjson"
)

var db *gorm.DB

const baseURL = "http://localhost:8080/api/v1/"

var baseDirectory = models.BaseDirectory

const strictTest = true

var apppath = baseDirectory

func TestMain(m *testing.M) {
	beego.TestBeegoInit(apppath)

	db = models.InitTestDataBase()
	db = services.InitTestDataBase()
	os.Exit(m.Run())
}

// func CompareJSONElement(t *testing.T, expectJSON string, expectPath string, resultJSON string, resultPath string) {
// 	expectStr := gjson.Get(expectJSON, expectPath).String()
// 	resultStr := gjson.Get(resultJSON, resultPath).String()

// 	if strictTest {
// 		assert.NotContains(t, expectStr, "[]", expectPath)
// 		assert.NotContains(t, resultStr, "[]", expectPath)
// 	}
// 	assert.Equal(t, expectStr, resultStr, expectPath)
// }

// func CompareDate(t *testing.T, expectJSON string, expectPath string, resultJSON string, resultPath string) {
// 	expectStr := extractDate(gjson.Get(expectJSON, expectPath).String())
// 	resultStr := extractDate(gjson.Get(resultJSON, resultPath).String())
// 	assert.Equal(t, expectStr, resultStr, expectPath)
// }

// func extractDate(str string) (ret string) {
// 	runes := []rune(str)
// 	return string(runes[0:10])
// }
