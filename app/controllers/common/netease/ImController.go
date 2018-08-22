package netease

import (
    "fmt"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/tidwall/sjson"
    . "wego3/bootstrap/components"
)

// class
type ImController string

// public
func (ctl *ImController) Notify(ctx *gin.Context) {
    // 简单的查询（无模型）
    var (
        bio string
        im_user string
        im_staff_tid string
    )

    row := DB.Table("base_user_profile").Where("user_id = ?", 7).Select("bio, im_user, im_staff_tid").Row()
    row.Scan(&bio, &im_user, &im_staff_tid)

    fmt.Println(bio)
    fmt.Println(im_user)
    fmt.Println(im_staff_tid)

    ctx.String(http.StatusOK, "测试网易云IM通知")
}

func (ctl *ImController) NotifyPost(ctx *gin.Context) {
    // try/catch
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
            ctx.String(http.StatusOK, "success")
        }
    }()

    RequestInit(ctx)
    // 测试PostJson的取值能力，去单值和取数组，
    fmt.Println(Request.PostJsonArray("convType"))
    fmt.Println("eventType: " + Request.PostJson("eventType"))
    fmt.Println("convType: " + Request.PostJson("convType.0"))

    // 假设要监控的accid(由于网易云小程序不支持群聊，所以暂时用单聊当客服)
    staff_accid := "bu-0000000007"

    if Request.PostJson("eventType") == "1" && Request.PostJson("convType") == "PERSON" {
        from_accid := Request.PostJson("fromAccount")
        to_accid   := Request.PostJson("to")

        if from_accid != staff_accid && to_accid != staff_accid {
            panic("不是客服咨询，不处理")
        }

        // 用户找客服聊天
        if to_accid == staff_accid {
            im_staff_key := "im:monitor_staff:" + from_accid

            // 如果存在就删除（说明已经看了，如果有已读抄送就更好了）
            if len, _ := Redis.HLen(im_staff_key).Result(); len > 0 {
                Redis.Del(im_staff_key)
            }

            // 新建一个（加q就是question的意思，其实是为了排序方便，排在staff前面）
            json, _ := Request.GetRawData()
            Redis.HSet(im_staff_key, "quser_" + Request.PostJson("msgTimestamp"), json)

            // 7天后过期
            Redis.Expire(im_staff_key, 3600*24*7*time.Second)

        // 客服找用户聊天
        } else {
            im_staff_key := "im:monitor_staff:" + to_accid

            // 允许客服主动发起聊天
            if len, _ := Redis.HLen(im_staff_key).Result(); len == 0 {
                // abort(403, '用户未咨询客服主动发起聊天，不处理');
                json_tpl := `{"fromAccount":"","last":"洗护咨询"}`
                json, _  := sjson.Set(json_tpl, "fromAccount", to_accid)

                Redis.HSet(im_staff_key, "quser_" + Request.PostJson("msgTimestamp"), json)
            }

            json, _ := Request.GetRawData()
            Redis.HSet(im_staff_key, "staff_" + Request.PostJson("msgTimestamp"), json)

            // 判断聊天信息的关键词，然后加个 'type' 进去
            if exists, _ := Redis.HExists(im_staff_key, "type").Result(); exists {
                Redis.HSet(im_staff_key, "type", "care")
            }
        }
    }

    // IM登入时记录登入记录
    if Request.PostJson("eventType") == "2" {
        im_user       := Request.PostJson("accid")
        im_online_key := "im:monitor_online:" + im_user

        // im很难24小时在线，如果有应该是出错了，所以设置一个过期时间
        Redis.Set(im_online_key, Request.PostJson("timestamp"), 3600*24*time.Second)
    }

    // IM登出时清除登入记录
    if Request.PostJson("eventType") == "3" {
        im_user       := Request.PostJson("accid")
        im_online_key := "im:monitor_online:" + im_user

        Redis.Del(im_online_key)
    }

    ctx.String(http.StatusOK, "success")
}
