package net

import (
	"HydraClient/util/log"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

func HttpPost(addHeaders map[string]string, data []byte, urlPath string) (body []byte, err error) {
	//bytesData, err := json.Marshal(data)
	//if err != nil {
	//	log.Debug.Println.Printf("marshal%v",err)
	//	return
	//}
	reader := bytes.NewReader(data)
	req, err := http.NewRequest("POST", urlPath, reader)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

// 文件上传
func HttpProxyFileUpload(file *multipart.FileHeader, fileKey string, addFields map[string]string,
	addHeaders map[string]string, urlPath string) (body []byte, err error) {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	formFile, err := writer.CreateFormFile(fileKey, file.Filename)
	if err != nil {
		log.Debug.Println("Upload Create form file failed")
		return
	}

	// 从文件读取数据，写入表单
	srcFile, err := file.Open()
	if err != nil {
		log.Debug.Println("Upload Create form file failed")
		return
	}
	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		log.Debug.Println("Write to form file failed")
		return
	}
	for fieldKey, fieldVal := range addFields {
		if err = writer.WriteField(fieldKey, fieldVal); err != nil {
			log.Debug.Println("WriteField failed")
			return
		}
	}
	// 发送表单
	contentType := writer.FormDataContentType()
	writer.Close() // 发送之前必须调用Close()以写入结尾行
	req, err := http.NewRequest("POST", urlPath, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Debug.Println("Post failed")

	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, _ = ioutil.ReadAll(resp.Body)
	return
}

func HttpPostJson(addHeaders map[string]string, data interface{}, urlPath string) (body []byte, err error) {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return
	}
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest("POST", urlPath, reader)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err = ioutil.ReadAll(resp.Body)
	return
}
func HttpPostJsoDownload(addHeaders map[string]string, data interface{}, urlPath string, filePathName string) error {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest("POST", urlPath, reader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	// Create output file
	out, err := os.Create(filePathName)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func HttpPostJsonPut(addHeaders map[string]string, data interface{}, urlPath string) (body []byte, err error) {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return
	}
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest("PUT", urlPath, reader)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

func HttpGet(addHeaders map[string]string, data map[string]string, urlPath string) (body []byte, err error) {
	params := url.Values{}
	urlInfo, err := url.Parse(urlPath)
	if err != nil {
		panic(err.Error())

	}
	for dataKey, dataVal := range data {
		params.Set(dataKey, dataVal)
	}
	urlInfo.RawQuery = params.Encode()
	fullUrl := urlInfo.String()
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return
	}
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err = ioutil.ReadAll(resp.Body)
	return
}
func HttpProxyFileUploadCustom(file *multipart.FileHeader, fileKey, filename string, addFields map[string]string,
	addHeaders map[string]string, urlPath string) (body []byte, err error) {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	formFile, err := writer.CreateFormFile(fileKey, filename)
	if err != nil {
		log.Debug.Println("Upload Create form file failed")
		return
	}

	// 从文件读取数据，写入表单
	srcFile, err := file.Open()
	if err != nil {
		log.Debug.Println("Upload Create form file failed")
		return
	}
	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		log.Debug.Println("Write to form file failed")
		return
	}
	for fieldKey, fieldVal := range addFields {
		if err = writer.WriteField(fieldKey, fieldVal); err != nil {
			log.Debug.Println("WriteField failed")
			return
		}
	}
	// 发送表单
	contentType := writer.FormDataContentType()
	writer.Close() // 发送之前必须调用Close()以写入结尾行
	req, err := http.NewRequest("POST", urlPath, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Debug.Println("Post failed")

	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, _ = ioutil.ReadAll(resp.Body)
	return
}
