package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func SaveData(filePath string, data []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func Utf8ToGbk(s []byte) (string, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return "", e
	}
	return string(d), nil
}

// Output struct as json format
func PrintStruct(s interface{}) string {
	b, err := json.Marshal(s)

	if err != nil {
		return fmt.Sprintf("%+v", s)
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", s)
	}

	return out.String()
}

func GetChineseCharacter(s string) string {
	var result string
	for _, v := range s {
		if v >= 0x4e00 && v <= 0x9fa5 {
			result += string(v)
		}
	}
	return result

	// return regexp.MustCompile("[^\u4e00-\u9fa5]").ReplaceAllString(s, "")
	// 经过测试，这两者没有什么区别
}

func RemoveDuplicate(data interface{}) interface{} {
	inArr := reflect.ValueOf(data)
	if inArr.Kind() != reflect.Slice && inArr.Kind() != reflect.Array {
		return data // 不是数组/切片
	}

	existMap := make(map[interface{}]bool)
	outArr := reflect.MakeSlice(inArr.Type(), 0, inArr.Len())

	for i := 0; i < inArr.Len(); i++ {
		iVal := inArr.Index(i)

		if _, ok := existMap[iVal.Interface()]; !ok {
			outArr = reflect.Append(outArr, inArr.Index(i))
			existMap[iVal.Interface()] = true
		}
	}

	return outArr.Interface()
}

func Base64EncodeHTTPImage(data []byte) string {
	return "data:" + http.DetectContentType(data) + "base64," + base64.StdEncoding.EncodeToString(data)
}

func JSONUnmarshalFromFile(filePath string, v any) error {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}