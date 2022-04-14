package nature

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

var (
	HttpClient *http.Client
)

func InitHttpClient() {
	HttpClient = &http.Client{
		Timeout: 30 * time.Second, // 请求超时时间
	}
}

// UploadFile 发送文件上传请求
func UploadFile(url string, params map[string]string, nameField, filename string, file io.Reader) ([]byte, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	// 在表单中创建一个文件字段
	formFile, err := writer.CreateFormFile(nameField, filename)
	if err != nil {
		return nil, err
	}
	// 读取文件内容到表单文件字段
	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, err
	}
	// 将其他参数写入到表单
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	if err = writer.Close(); err != nil {
		return nil, err
	}
	// 构造请求对象
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("Content-Type","multipart/form-data")
	req.Header.Add("Content-Type", writer.FormDataContentType())
	// 发送请求
	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
