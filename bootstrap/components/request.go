package components

import (
    // "fmt"
    // "reflect"
    "io/ioutil"
    "bytes"
    "github.com/gin-gonic/gin"
    "github.com/tidwall/gjson"
)

// 方案三（优点是能继承老方法）
type Context struct {
    gin.Context
}

var Request *Context

func RequestInit(ctx *gin.Context) {
    Request = &Context{(*ctx)}
}

// 说明：方法名想模拟laravel的，但最终和gin统一
// 两个不带Raw的方法取出的类型强转成了string
func (c *Context) PostJson(key string) string {
    body, _ := c.GetRawData()
    value := gjson.Get(string(body), key)

    return value.String()
}

func (c *Context) PostJsonArray(key string) []string {
    body, _ := c.GetRawData()
    value := gjson.Get(string(body), key)

    ret := make([]string, 0)

    for _, val := range value.Array() {
        ret = append(ret, val.String())
    }

    return ret
}

// 两个带Raw的方法取出的类型与json一致
func (c *Context) PostJsonRaw(key string) gjson.Result {
    body, _ := c.GetRawData()
    value := gjson.Get(string(body), key)

    return value
}

func (c *Context) PostJsonArrayRaw(key string) []gjson.Result {
    body, _ := c.GetRawData()
    value := gjson.Get(string(body), key)

    return value.Array()
}

func (c *Context) GetRawData() ([]byte, error) {
    raw_data, err := ioutil.ReadAll(c.Request.Body)

    buf := bytes.NewBuffer(raw_data)
    c.Request.Body = ioutil.NopCloser(buf)

    return raw_data, err
}


// 方案一
// type Request []byte

// func (req Request) JsonGet(key string) string {
//     value := gjson.Get(string(req), key)
//     return value.String()
// }

// // 方案二（缺点是不能继承方法）
// type Context gin.Context

// var Request *Context

// func RequestInit(ctx *gin.Context) {
//     c := Context((*ctx))
//     Request = &c
// }

// 备份，想模拟laravel，但不容易（from和query都是字符串，但json有bool，int等）
// // 这个只涵盖了 string 返回型
// func (c *Context) Input(key string) string {
//     var value string

//     switch c.ContentType() {
//     case "application/json":
//         value = c.JsonGet(key)
//     case "application/x-www-form-urlencoded":
//         value = c.FormGet(key)
//         if value == "" {
//             value = c.QueryGet(key)
//         }
//     }

//     return value
// }

// func (c *Context) InputArray(key string) []string {
//     var values []string

//     switch c.ContentType() {
//     case "application/json":
//         value = c.JsonGet(key)
//     case "application/x-www-form-urlencoded":
//         values, _ = c.PostFormArray(key)
//         if values == []string{} {
//             values, _ = c.GetQueryArray(key)
//         }
//     }

//     return values
// }

// func (c *Context) QueryGet(key string) string {
//     value, _ := c.GetQuery(key)
//     return value
// }

// func (c *Context) FormGet(key string) string {
//     value, _ := c.GetPostForm(key)
//     return value
// }
