package main 

import (
         "io"
         "net/http"
         "log"
         "fmt"
         "crypto/sha1"
         "sort"
         "encoding/xml"
         "time"
         "strings"
         )

const (
   TOKEN = "astaxiefromqqweixin"
   Text = "text"
   Location = "location"
   Image = "image"
   Link = "link"
   Event = "event"
   Music = "music"
   News = "news"
)

type msgBase struct {
ToUserName string
FromUserName string
CreateTime time.Duration
MsgType string
Content string
}


type Request struct {
XMLName xml.Name `xml:"xml"`
msgBase // base struct
Location_X, Location_Y float32
Scale int
Label string
PicUrl string
MsgId int
}


func Signature(timestamp, nonce string) string {
   strs := sort.StringSlice{TOKEN, timestamp, nonce}
   sort.Strings(strs)
   str := ""
   for _, s := range strs {
      str += s
   }
   h := sha1.New()
   h.Write([]byte(str))
   return fmt.Sprintf("%x", h.Sum(nil))
}

func getReqHeader_stub(str string) string{
  return str
}
func handShakeGet(w http.ResponseWriter, r * http.Request) {
   signature := getReqHeader_stub("signature")
   timestamp := getReqHeader_stub("timestamp")
   nonce     := getReqHeader_stub("nonce")
   echostr   := getReqHeader_stub("echostr")
   if Signature(timestamp, nonce) == signature {
      io.WriteString(w, echostr)
   }else{
      io.WriteString(w, " ")
   }
}

func DecodeRequest(data []byte) (req *Request, err error) {
   req = &Request{}
   if err = xml.Unmarshal(data, req); err != nil {
      return
   }
   req.CreateTime *= time.Second
   return
}

func handlerPost(w http.ResponseWriter, r * http.Request) {
   
}
func weixinhandler( w http.ResponseWriter, r * http.Request) {
   // io.WriteString(w, "Hello, world!")
   log.Println("method is :", r.Method)
   if r.Method == "GET" {
    r.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
    //注意:如果没有调用ParseForm方法，下面无法获取表单的数据
    fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }

    handShakeGet(w, r)

    io.WriteString(w, "Hello, world!")
   }else {
      handlerPost(w, r)
   }
}

func main() {
   http.HandleFunc("/weixin", weixinhandler)
   err := http.ListenAndServe(":8080", nil)
   if err != nil {
      log.Fatal("God like")
   }
   log.Println("run server at 8080")
}