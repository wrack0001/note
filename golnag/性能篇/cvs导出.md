# 论：CVS导出

> ** 导论：** 我们经常会给PM或者运营导出数以百万计的数据
>> 我们在导出的时候一般分两「在线」和「离线」导出，此处我会已在线导出的方式给出例子


## 思路
    1. 数据量比较大需要尽量多的复用同一块内存
    2. 分批将数据Flush给浏览器
    3. 内存使用完成后及时GC

## 注意
    - 此处只针对管理后台使用：如果是大量并发肯定会出问题。如果是针对外部下载，还是使用离线+COS的方式比较靠谱
    - 管理后台也是有并发的：也可以增加个锁啥的，如果很多个导出也是一个服务呢，这就看自己的场景了。

## 代码

```go
// Student 学生
type Student struct {
	ID string
	Name string
	Age string
	Sex string
}

const (
    // ExportFileName 导出文件名
    ExportFileName = "敏捷调研明细-%s.csv"
    // ExportLimit 导出单次限制
    ExportLimit = 50000
)

// CVSExport CVS 导出方法
func CVSExport(ctx context.Context, req *http.Request, rsp http.ResponseWriter) error {
	// 方式并发导出有问题的话，可以增加个锁
	id := req.URL.Query().Get("id")
	if id == "" {
		return fmt.Errorf("ID不能为空")
	}
	e := export.NewClient(rsp, fmt.Sprintf(ExportFileName, surveyID))
	data := []string{
		"ID", "姓名", "年龄", "性别",
	}
	if err := e.Writer(data); err != nil {
		return err
	}
	list := make([]*Student,0,ExportLimit)
	for i := 0; i < 100; i++ {
		// 查询CH中数据【想要复用内存，可以将返回值传入的方式】
		if err := mysqlCli.Query(ctx,list, id, int64(ExportLimit*i), ExportLimit);err != nil {
			return err
		}
		if len(list) == 0 {
			break
		}
		for _, v := range list {
			// 内存复用
			data = append(data[:0], v.ContentID, v.Qimei, v.UID, v.ABTestID, question, v.AType, v.ExpEventTime, v.ClickEventTime)
			if err = e.Writer(data); err != nil {
				return err
			}
		}
        // 重置切片
        list = list[:0]
		// 数据Flush
		e.Flush()
	}
	return nil
}


```

```go
// Package export 导出包文件
package export

import (
	"bytes"
	"encoding/csv"
	"net/http"
	"net/url"
	"runtime"
)

// Client CVS 接口导出
type Client interface {
	Writer(data []string) error
	Flush()
}

//go:generate mockgen -source=csv.go -destination=./mock/csv.go -package=exportMock

// client CVS客户端
type client struct {
	rw     http.ResponseWriter
	buffer *bytes.Buffer
	csv    *csv.Writer
}

// NewClient 实例化cvs导出
func NewClient(rw http.ResponseWriter, fileName string) Client {
	buffer := &bytes.Buffer{}
	rw.Header().Add("Content-type", "text/csv")
	rw.Header().Add("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
	// 解决导出UTF-8乱码问题
	buffer.Write([]byte("\xEF\xBB\xBF"))
	return &client{
		rw:     rw,
		buffer: buffer,
		csv:    csv.NewWriter(buffer),
	}
}

// Writer 数据写入缓存
func (s *client) Writer(data []string) error {
	// 写数据到csv文件
	if err := s.csv.Write(data); err != nil {
		return err
	}
	// 写数据到csv文件到buffer
	s.csv.Flush()
	return nil
}

// Flush 数据Flush【不要等所有的数据都loading后再Flush】
func (s *client) Flush() {
	_, _ = s.rw.Write(s.buffer.Bytes())
	// response := http.NewResponseController(s.rw)
	// response.Flush()
	// 数据Flush返回后重置buffer
	s.buffer.Reset()
	// 内存GC 【是否需要手动GC看机器本身内存大小】因为被动GC并不是实时触发
	runtime.GC()
}

```

## 参考文档
- [Go垃圾收集器指南](https://tip.golang.org/doc/gc-guide)
- [Go Slices：用法和内部结构](https://tip.golang.org/blog/slices-intro)