package controller

import (
	json2 "encoding/json"
	"gorm_tcp/dao"
	"gorm_tcp/model"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetUsers 获取全部用户数据
func GetUsers(w http.ResponseWriter, r *http.Request) {
	//接收GET请求
	if strings.ToUpper(r.Method) == "GET" {
		users, err := dao.GetUsers() //将数据库中找到的内容赋给结构体users
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json, err := json2.Marshal(users) //将结构体users转化为json编码
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json") //设置Content-Type参数json
		w.WriteHeader(http.StatusOK)                       //Header添加状态码
		w.Write(json)                                      //输出json编码
	} else {
		http.Error(w, "请求方式错误", http.StatusNotFound)
	}
}

// GetUserName 通过查找名字，获取单个用户详细数据
func GetUserName(w http.ResponseWriter, r *http.Request) {
	//接收GET请求
	if strings.ToUpper(r.Method) == "GET" {
		//定义u结构体
		var u model.User
		body, err := ioutil.ReadAll(r.Body) //读取r.Body的json内容赋给body
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) //err不为空，则返回错误信息
			return
		}
		err = json2.Unmarshal(body, &u) //解析body的json编码的数据并将结果存在结构体u中
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) //err不为空，则返回错误信息
			return
		}

		err1 := json2.Valid([]byte(body)) //判断传入json编码格式
		if err1 == false {
			http.Error(w, "传入参数格式不对", http.StatusBadRequest)
			return
		}

		username := &model.User{
			Name: u.Name,
		}
		user, data, err := dao.GetUserName(username)
		if data != nil {
			http.Error(w, string(data), http.StatusBadRequest)
			return
		}
		if err != nil { //数据库怠机，返回错误
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json, err := json2.Marshal(user) //将user结构体转化为json编码
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json") //设置Content-Type参数json
		w.WriteHeader(http.StatusOK)                       //Header添加状态码
		w.Write(json)                                      //输出json编码
	} else {
		http.Error(w, "请求方式错误", http.StatusNotFound)
	}
}

//AddUser 添加用户数据
func AddUser(w http.ResponseWriter, r *http.Request) {
	//接收POST请求
	if strings.ToUpper(r.Method) == "POST" {

		//定义u结构体
		var u model.User
		body, err := ioutil.ReadAll(r.Body) //读取r.Body的json内容赋给body
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) //err不为空，则返回错误信息
			return
		}

		err = json2.Unmarshal(body, &u) //解析body的json编码的数据并将结果存在结构体u中
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) //err不为空，则返回错误信息
			return
		}

		err1 := json2.Valid([]byte(body)) //判断传入json编码格式
		if err1 == false {
			http.Error(w, "传入参数格式不对", http.StatusBadRequest)
			return
		}

		user := &model.User{
			Name: u.Name,
		}
		if len(u.Name) < 2 || len(u.Name) > 10 { //添加限制条件，对于输入的名字小于2或者大于10的值返回错误
			http.Error(w, "长度格式错误", http.StatusBadRequest)
			return
		}
		users, data, err := dao.AddUser(user) //进行数据库的增加操作
		if data != nil {                      //添加相同名字的限制条件，如果相同就返回错误
			http.Error(w, string(data), http.StatusBadRequest)
			return
		}
		if err != nil { //数据库怠机，返回错误
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json, err := json2.Marshal(users) //将user结构体转化为json编码
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json") //设置Content-Type参数json
		w.WriteHeader(http.StatusCreated)                  //Header添加状态码
		w.Write(json)                                      //输出json编码
	} else {
		http.Error(w, "请求方式错误", http.StatusNotFound)
	}
}

// DpdUserName 修改用户名字
func DpdUserName(w http.ResponseWriter, r *http.Request) {
	//接收PUT的修改请求
	if strings.ToUpper(r.Method) == "PUT" {

		body, err := ioutil.ReadAll(r.Body) //读取r.Body里的json编码传入body里
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var u model.Username            //定义结果体u(包含OldName和NewName)
		err = json2.Unmarshal(body, &u) //对body里的json编码进行解码操作，传入结构体u中
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) //err不为空，则返回错误信息
			return
		}
		err1 := json2.Valid([]byte(body)) //判断传入json编码格式
		if err1 == false {
			http.Error(w, "传入参数格式不对", http.StatusBadRequest)
			return
		}
		username := &model.Username{ //将结构体u中的值传入定义的username中
			OldName: u.OldName,
			NewName: u.NewName,
		}
		if len(u.NewName) < 2 || len(u.NewName) > 10 { //添加限制条件，对于输入的名字小于2或者大于10的值返回错误
			http.Error(w, "长度格式错误", http.StatusBadRequest)
			return
		}
		users, data, err := dao.DpdUser(username)
		if data != nil { //添加相同名字的限制条件，如果相同就返回错误
			http.Error(w, string(data), http.StatusBadRequest)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json, err := json2.Marshal(users) //对结构体users进行编码操作成json格式
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(json)
	} else {
		http.Error(w, "请求方式错误", http.StatusNotFound)
	}
}

// DelUser 删除用户，执行软删除操作
func DelUser(w http.ResponseWriter, r *http.Request) {
	//接收DELETE请求，对用户进行删除操作
	if strings.ToUpper(r.Method) == "DELETE" {

		var u model.User
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonData := body
		err1 := json2.Valid([]byte(jsonData)) //判断传入json编码格式
		if err1 == false {
			http.Error(w, "传入参数格式不对", http.StatusBadRequest)
			return
		}
		err = json2.Unmarshal(jsonData, &u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) //err不为空，则返回错误信息
			return
		}
		user := &model.User{
			Name: u.Name,
		}
		users, data, err := dao.DelUser(user)
		if data != nil {
			http.Error(w, string(data), http.StatusBadRequest)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json, err := json2.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(json)
	} else {
		http.Error(w, "请求方式错误", http.StatusNotFound)
	}

}

//func DelUserID(w http.ResponseWriter, r *http.Request) {
//	if strings.ToUpper(r.Method) == "DELETE" {
//		//id := r.PostFormValue("id")
//		//intid, _ := strconv.Atoi(id)
//		//uintid := uint(intid)
//		var u model.User
//		body, err1 := ioutil.ReadAll(r.Body)
//		if err1 != nil || len(body) == 0 {
//			w.WriteHeader(http.StatusBadRequest)
//			io.WriteString(w, "请求的语法错误")
//			return
//		}
//		jsonData := []byte(body)
//		json2.Unmarshal(jsonData, &u)
//		user := &model.User{
//			Model: gorm.Model{
//				ID: u.ID,
//			},
//		}
//		users, _ := dao.DelUserID(user)
//		json, err2 := json2.Marshal(users)
//		if err2 != nil {
//			http.Error(w, err2.Error(), http.StatusBadRequest)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusCreated)
//		w.Write(json)
//	} else {
//		http.Error(w, "请求方式错误", http.StatusNotFound)
//	}
//}
