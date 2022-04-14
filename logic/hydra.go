package logic

import (
	"HydraClient/util/conf"
	"HydraClient/util/log"
	"HydraClient/util/net"
	"encoding/json"
	"fmt"
	//"github.com/gookit/goutil/dump"
	"strconv"
	"strings"
)

type Report struct {
	Generator struct {
		Software          string `json:"software"`
		Version           string `json:"version"`
		Built             string `json:"built"`
		Server            string `json:"server"`
		Service           string `json:"service"`
		Jsonoutputversion string `json:"jsonoutputversion"`
		Commandline       string `json:"commandline"`
	} `json:"generator"`
	Results []struct {
		Port     int    `json:"port"`
		Service  string `json:"service"`
		Host     string `json:"host"`
		Login    string `json:"login"`
		Password string `json:"password"`
	} `json:"results"`
	Success       bool          `json:"success"`
	Errormessages []interface{} `json:"errormessages"`
	Quantityfound int           `json:"quantityfound"`
}
type Reflexion struct {
	TaskId int    `json:"task_id"`
	Report []byte `json:"report"`
	Status bool   `json:"status"`
}
type HydraArgs struct {
	TaskId       int    `json:"task_id"`
	TaskName     string `json:"task_name"`
	Address      string `json:"address"`
	Port         string `json:"port"`
	Protocol     string `json:"protocol"`
	Username     string `json:"username"`
	UsernameFile string `json:"username_file"`
	UserNameType int    `json:"user_name_type"` // 1 默认 2 手写 3 上传
	Password     string `json:"password"`
	PasswordFile string `json:"password_file"`
	PasswordType int    `json:"passwd_type"` // 1 默认 2 手写 3 上传
	UserId       int    `json:"user_id"`
	Path         string `json:"path"`
	Form         string `json:"form"`
	Sid          string `json:"sid"`
	RequestHost  string `json:"request_host"`
}
type Transport struct {
	TaskId       int      `json:"task_id"`
	TaskString   string   `json:"task_string"`
	Username     string   `json:"username"`
	UsernameType int      `json:"username_type"`
	Password     string   `json:"password"`
	PasswordType int      `json:"password_type"`
	CmdLine      []string `json:"cmd_line"`
}

func CreateTask(args *HydraArgs) {
	defer func() {
		if err := recover(); err != nil {
			net.HttpPost(nil, []byte{}, args.RequestHost)
		}
	}()
	cmdline, _ := splice(args)
	fmt.Printf("任务执行前最终要传递的命令是%v\n", cmdline)
	Trans(args, cmdline)
	log.Report.Println("新任务创建")
	log.Report.Printf("任务ID\t%v\n", args.TaskId)
	log.Report.Printf("任务名称\t%v\n", args.TaskId)

}
func Trans(args *HydraArgs, cmdline []string) {
	t := &Transport{
		TaskId:       args.TaskId,
		TaskString:   strconv.Itoa(args.TaskId),
		Username:     args.Username,
		UsernameType: args.UserNameType,
		Password:     args.Password,
		PasswordType: args.PasswordType,
		CmdLine:      cmdline,
	}
	m, err := json.Marshal(t)
	if err != nil {
		log.Debug.Println("结构体序列化失败")
	} else {
		log.Info.Printf("结构体序列化: %s\n", string(m))
	}
	host := conf.GetVal("server", "host")
	port := conf.GetVal("server", "port")
	master := strings.Join([]string{host, port}, ":")
	masterUrl := strings.Join([]string{master, "hydra", "create"}, "/")
	log.Info.Printf("创建任务上传参数的地址是:%s\n", masterUrl)
	if res, err := net.HttpPost(nil, m, masterUrl); err != nil {
		log.Debug.Printf("将参数上传到服务端时出现错误:%v\n错误内容%v\n", err, err.Error())
	} else {
		log.Info.Printf("将参数上传到服务端返回内容:%v\n", res)
	}
}

func splice(args *HydraArgs) ([]string, error) {
	log.Debug.Printf("docker recv args is %+v\n", args)

	var cmdline []string
	switch args.Protocol {
	case "http-form-get",
		"http-form-post",
		"https-form-get",
		"https-form-post": // 需要表单和路径
		//cmdline = Patch4HttpWithForm(*args)
	case "http-get",
		"https-get",
		"http-post",
		"https-post",
		"http-head",
		"https-head": // 需要路径
		//cmdline = Patch4HttpWithoutForm(*args)
	case "ftp",
		"imap",
		"ldap",
		"mssql",
		"mysql",
		"postgres",
		"rdp",
		"rexec",
		"rlogin",
		"rsh",
		"rtsp",
		"smb",
		"smtp",
		"ssh",
		"telnet",
		"pop3": //需要用户名和密码
		cmdline = Patch4Other(*args)
	case "snmp",
		"vnc",
		"oracle-listener": //只需要密码
		//cmdline = Patch4OnlyPasswd(*args)
	case "oracle-sid": // 只需要用户名
		//cmdline = Patch4OracleSid(*args)
	}
	log.Info.Printf("通过switch,拼接后的命令为%v\n", cmdline)
	return cmdline, nil
}

const (
	thread = "1"
	time   = "30" // 单位 秒
)

func Patch4Other(h HydraArgs) []string {
	var line = []string{"hydra"}
	UsernameFullPath := strings.Join([]string{"username", strconv.Itoa(h.TaskId)}, "/")
	PasswordFullPath := strings.Join([]string{"password", strconv.Itoa(h.TaskId)}, "/")
	log.Debug.Printf("传入docker的unType参数: %v\n", h.UserNameType)
	switch h.UserNameType { // 1:默认字典 2:手动录入 3:自定义上传
	case 1:
		line = append(line, "-L")
		line = append(line, strings.Join([]string{"default", "username"}, "/"))
	case 2:
		line = append(line, "-L")
		line = append(line, strings.Join([]string{"username", strconv.Itoa(h.TaskId)}, "/"))
	case 3:
		line = append(line, "-L")
		line = append(line, UsernameFullPath)
	}

	switch h.PasswordType { // 0:默认字典 1:手动录入 2:自定义上传
	case 1:
		line = append(line, "-P")
		line = append(line, strings.Join([]string{"default", "password"}, "/"))
	case 2:
		line = append(line, "-P")
		line = append(line, strings.Join([]string{"password", strconv.Itoa(h.TaskId)}, "/"))
	case 3:
		line = append(line, "-P")
		line = append(line, PasswordFullPath)
	}

	outfile := strings.Join([]string{"report", strconv.Itoa(h.TaskId)}, "/")
	outfile = strings.Join([]string{outfile, "json"}, ".")
	if h.Port != "" {
		line = append(line, "-s", h.Port)
	}
	line = append(line, "-t", thread, h.Address, h.Protocol, "-b", "json", "-o", outfile)
	log.Info.Printf("生成返回之前的命令:%v\n", line)
	return line
}
func SaveReport(report []byte, tid int) {
	var r Report
	if err := json.Unmarshal(report, &r); err != nil {
		log.Debug.Printf("任务%d报告解析为结构体失败\n", tid)
		log.Report.Printf("当前任务返回结果异常\n")
	}
	prefix := strings.Split(r.Generator.Commandline, "report/")[1]
	suffix := strings.Split(prefix, ".json")[0]
	log.Report.Printf("任务ID\t%v\n", suffix)
	log.Report.Printf("创建时间\t%v\n", r.Generator.Built)
	log.Report.Printf("状态\t%v\n", r.Success)
	if r.Success {
		for i,result:= range r.Results{
			log.Report.Printf("成功记录第%d条\t%v\n",i+1,result)
		}
	} else {
		for i,errormessage:= range r.Errormessages{
			log.Report.Printf("错误记录第%d条\t%v\n",i+1,errormessage)
		}
	}
}
