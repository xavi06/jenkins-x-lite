package util

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// DoBytesPost func
// 以二进制POST
func DoBytesPost(url string, data []byte) ([]byte, error) {

	body := bytes.NewReader(data)
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return []byte(""), err
	}
	//request.Header.Set("Connection", "Keep-Alive")
	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	return b, err
}

// UploadFile func
func UploadFile(filename string, name string, url string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadFile", filename)
	if err != nil {
		return err
	}
	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		return err
	}

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	fieldWriter, err := bodyWriter.CreateFormField("filename")
	if err != nil {
		return err
	}
	_, err = fieldWriter.Write([]byte(name))
	if err != nil {
		return err
	}
	bodyWriter.Close()

	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//fmt.Println(resp.Status)
	//fmt.Println(string(respBody))
	return nil
}

// DoDelete func
func DoDelete(url string) error {

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	return err

}
