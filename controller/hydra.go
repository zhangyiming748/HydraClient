package controller

import (
	"HydraClient/logic"
	"HydraClient/util/conf"
	"HydraClient/util/log"
	"HydraClient/util/net"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/dump"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
)

func Hydra(ctx *gin.Context) {

	serverHost := conf.GetVal("server", "host")
	serverPort := conf.GetVal("server", "port")
	server := strings.Join([]string{serverHost, serverPort}, ":")
	fileUrl := strings.Join([]string{server, "hydra", "upload"}, "/")

	UserId := 1
	TaskName := ctx.Request.FormValue("task_name")
	Address := ctx.Request.FormValue("address")
	Port := ctx.Request.FormValue("port")
	Protocol := ctx.Request.FormValue("protocol")
	Username := ctx.Request.FormValue("username")
	UsernameFile := ""
	UsernameType := ctx.Request.FormValue("username_type")
	Password := ctx.Request.FormValue("password")
	PasswordFile := ""
	PasswordType := ctx.Request.FormValue("password_type")
	Path := ctx.Request.FormValue("path")
	Form := ctx.Request.FormValue("form")
	Sid := ctx.Request.FormValue("sid")
	ut, _ := strconv.Atoi(UsernameType)
	pt, _ := strconv.Atoi(PasswordType)
	requestHost := strings.Join([]string{strings.Join([]string{conf.GetVal("client", "host"), conf.GetVal("client", "port")}, ":"), "hydra", "recv"}, "/")

	args := &logic.HydraArgs{
		TaskId:       rand.Intn(65535),
		TaskName:     TaskName,
		Address:      Address,
		Port:         Port,
		Protocol:     Protocol,
		Username:     Username,
		UsernameFile: UsernameFile,
		UserNameType: ut,
		Password:     Password,
		PasswordFile: PasswordFile,
		PasswordType: pt,
		UserId:       UserId,
		Path:         Path,
		Form:         Form,
		Sid:          Sid,
		RequestHost:  requestHost,
	}

	if UsernameType == "3" {
		file, fileHeader, err := ctx.Request.FormFile("username_file")
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		log.Info.Println(file, fileHeader)
		//todo 传给server
		m := make(map[string]string)
		m["type"] = "username"
		tid := strconv.Itoa(args.TaskId)
		m["task_id"] = tid
		UserNameUrl := strings.Join([]string{fileUrl, "username"}, "/")
		c, err := net.HttpProxyFileUploadCustom(fileHeader, "username_file", strconv.Itoa(args.TaskId), m, nil, UserNameUrl)
		log.Info.Printf("用户名字典:%s\n任务ID:%s\n上传地址%s\n", fileHeader.Size, m, UserNameUrl)
		if err != nil {
			log.Debug.Println("上传给server数据时产生的错误")
			dump.P(args)
			dump.P(fileUrl)
			dump.P(c)
			return
		} else {
			log.Info.Printf("返回的数据: %v\n", c)
			dump.P(args)
			dump.P(fileUrl)
			dump.P(c)
		}
	}
	if PasswordType == "3" {
		file, fileHeader, err := ctx.Request.FormFile("password_file")
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		log.Info.Println(file)
		//todo 传给server
		m := make(map[string]string)
		m["type"] = "password"
		tid := strconv.Itoa(args.TaskId)
		m["task_id"] = tid
		PasswordUrl := strings.Join([]string{fileUrl, "password"}, "/")
		c, err := net.HttpProxyFileUploadCustom(fileHeader, "password_file", strconv.Itoa(args.TaskId), m, nil, PasswordUrl)
		log.Info.Printf("密码字典:%s\n任务ID:%s\n上传地址%s\n", fileHeader.Size, m, PasswordUrl)
		if err != nil {
			log.Debug.Println("上传给server数据时产生的错误")
			dump.P(args)
			dump.P(fileUrl)
			dump.P(c)
			return
		} else {
			log.Info.Printf("返回的数据: %v\n", c)
			dump.P(args)
			dump.P(fileUrl)
			dump.P(c)
		}
	}
	dump.P(args)

	go logic.CreateTask(args)
	ctx.JSON(200, args)
}

func Recv(ctx *gin.Context) {
	log.Info.Println("recv接口收到消息")
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	log.Info.Printf("ctx.Request.body: %v", string(data))
	var r logic.Reflexion
	if err := json.Unmarshal(data, &r); err != nil {
		log.Debug.Println("解析为结构体失败")
		log.Report.Printf("当前任务返回结果异常\n")
	}
	if r.Status {
		logic.SaveReport(r.Report, r.TaskId)
	} else {
		log.Report.Printf("任务%d在密码破解命令执行的时候出错,详情需要查询日志\n", r.TaskId)
	}
	log.Info.Printf("任务%d返回的报告全文:%s\n", r.TaskId, r.Report)
	ctx.JSON(200, data)
}
