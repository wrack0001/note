# base64与图片互转.md
>> 此文章主要是演示图片转base64（base64转图片）

```go
package main

import (
	"encoding/base64"
	"io/ioutil"
)

const (
	picPath = `base64_img/image.jpg`
	base64Path = `base64_img/test.txt`
	picNewPath = `base64_img/image_new.jpg`
)

func main() {
	// 1. 将图片转成字符串
	data,err := img2base64(picPath)
	if err != nil {
		panic(err)
	}
	// 2. 文件写成 base64   imgFile -> base64
	_ = ioutil.WriteFile(base64Path, []byte(data), 0666)
	// 如果想增加图片头，下面增加
	// _ = ioutil.WriteFile(base64Path, []byte("data:image/jpeg;base64," + data), 0666)
	// 3. 读取文件内容
	f, _ := ioutil.ReadFile(base64Path) 
	// 4. 将读取出来的文件存成文件
	if err = base642img(string(f),picNewPath);err != nil {
		panic(err)
	}

}

// img2base64 图片转base64
func img2base64(picPath string) (string,error) {
	// 1. 读取文件到内存
	f, err := ioutil.ReadFile(picPath)
	if err != nil {
		return "",err
	}
	// 2. 将图片转成 base64 字符串
	base64Desc := base64.StdEncoding.EncodeToString(f)
	return base64Desc,nil
}

// base642img 将base64字符串转图片 [注意：注意转码的时候若有有 【data:image/jpeg;base64,】 需要将其去掉]
func base642img(base64Desc,filePath string) error  {
	// 1. 将base64字符串解码成[]byte
	data, err := base64.StdEncoding.DecodeString(base64Desc)
	if err != nil {
		return err
	}
	// 2. buffer输出到jpg文件中（不做处理，直接写到文件）
	return ioutil.WriteFile(filePath, data, 0666)
}
```