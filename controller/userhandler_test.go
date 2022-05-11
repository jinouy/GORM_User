package controller

import (
	"bytes"
	json2 "encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm_tcp/model"
	"net/http"
	"testing"
)

func TestGetUsers(t *testing.T) {

	testCases := []struct { //定义测试的结构体(测试不同请求)
		TestName   string
		HttpDate   string
		StatusCode int
	}{
		//测试组
		{"TestCase1_GET", "GET", 200},
		{"TestCase2_POST", "POST", 404},
		{"TestCase3_DEL", "DELETE", 404},
	}
	for _, testCase := range testCases { //进行三次测试
		t.Run(testCase.TestName, func(t *testing.T) { //子测试

			//req设定了结构体中的请求
			req, err := http.NewRequest(testCase.HttpDate, "http://localhost:8080/getUsers", nil)
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json; charset-UTF-8")
			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			assert.Equal(t, testCase.StatusCode, resp.StatusCode, "They should be equal")
		})
	}
}

func TestGetUserName(t *testing.T) {

	testCases := []struct { //定义测试的结构体(测试不同请求)
		TestName   string
		UserDate   string
		StatusCode int
	}{
		//测试组
		{"TestCase1_Right", "tom", 200},
		{"TestCase2_Unknown", "ttt", 400},
	}
	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			u := &model.User{ //将UserDate的内容传入结构体u中
				Name: testCase.UserDate,
			}
			json, err := json2.Marshal(u) //对结构体u进行json的编码操作
			require.NoError(t, err)

			//req是GET的响应对象
			req, err := http.NewRequest("GET", "http://localhost:8080/getUserName", bytes.NewReader(json))
			require.NoError(t, err)

			req.Header.Set("Content-Type", "application/json; charset-UTF-8")
			resp, err := http.DefaultClient.Do(req)

			require.NoError(t, err)
			assert.Equal(t, testCase.StatusCode, resp.StatusCode, "They should be equal")
		})

	}
}

func TestAddUser(t *testing.T) {

	testCases := []struct { //定义测试的结构体
		TestName   string
		UserDate   string
		StatusCode int
	}{
		//测试组
		{"TestCase1_Repeat", "joy", 400},
		{"TestCase2_Length", "k", 400},
		{"TestCase3_Right", "test1111", 201},
	}
	for _, testCase := range testCases { //进行三次测试
		t.Run(testCase.TestName, func(t *testing.T) { //子测试
			u := &model.User{Name: testCase.UserDate}
			json, err := json2.Marshal(u) //将结构体u转化为json编码
			require.NoError(t, err)

			// req为POST请求的具体对象
			req, err := http.NewRequest("POST", "http://localhost:8080/addUser", bytes.NewReader(json))
			require.NoError(t, err)

			req.Header.Set("Content-Type", "application/json; charset-UTF-8")
			resp, err := http.DefaultClient.Do(req) //resp返回http的响应
			require.NoError(t, err)

			//通过判断结构体中设定的status与实际的传入的status来测试是否通过
			assert.Equal(t, testCase.StatusCode, resp.StatusCode, "They should be equal")
		})
	}
}

func TestDpdUserName(t *testing.T) {
	testCases := []struct { //定义测试的结构体
		TestName   string
		OldDate    string
		NewDate    string
		StatusCode int
	}{
		//测试组
		{"TestCase1_Right", "joy", "joy1", 201},
		{"TestCase2_Length", "1111", "k", 400},
		{"TestCase3_Unknown", "te", "22", 400},
		{"TestCase3_Repeat", "jack", "mark", 400},
	}
	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			u := &model.Username{
				OldName: testCase.OldDate,
				NewName: testCase.NewDate,
			}
			json, err := json2.Marshal(u) //将结构体u转化为json编码
			require.NoError(t, err)
			// req为POST请求的具体对象
			req, err := http.NewRequest("PUT", "http://localhost:8080/updateUserName", bytes.NewReader(json))
			require.NoError(t, err)

			req.Header.Set("Content-Type", "application/json; charset-UTF-8")
			resp, err := http.DefaultClient.Do(req) //resp返回http的响应
			require.NoError(t, err)

			//通过判断结构体中设定的status与实际的传入的status来测试是否通过
			//body, _ := ioutil.ReadAll(resp.Body)   输出错误
			//fmt.Println(string(body))
			assert.Equal(t, testCase.StatusCode, resp.StatusCode, "They should be equal")
		})
	}
}

func TestDelUser(t *testing.T) {
	testCases := []struct { //定义测试的结构体
		TestName   string
		UserDate   string
		StatusCode int
	}{
		//测试组
		{"TestCase1_Unknown", "want", 400},
		{"TestCase3_Right", "222", 201},
	}
	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			u := &model.User{ //将测试的UserDate内容传入结构体u中
				Name: testCase.UserDate,
			}
			json, err := json2.Marshal(u) //对结构体u进行json编码
			require.NoError(t, err)

			//执行DELETE操作，将返回的值传入req中，req为DELETE请求的具体对象
			req, err := http.NewRequest("DELETE", "http://localhost:8080/deleteUserName", bytes.NewReader(json))
			require.NoError(t, err)

			req.Header.Set("Content-Type", "application/json; charset-UTF-8")
			resp, err := http.DefaultClient.Do(req) //resp返回http的响应
			require.NoError(t, err)

			//通过判断结构体中设定的status与实际的传入的status来测试是否通过
			assert.Equal(t, testCase.StatusCode, resp.StatusCode, "They should be equal")
		})
	}
}
