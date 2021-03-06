package controllers

import (
         "io"
         "io/ioutil"
         "net/http"
         "log"
         "fmt"
         "crypto/sha1"
         "sort"
         "encoding/xml"
         "time"
         "errors"
         )

const (
   TOKEN = "every_date_love_letter"
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

type Response struct {
    XMLName xml.Name `xml:"xml"`
    msgBase
}
func TestFunc(){
  fmt.Println("hello package")
}

func createSignature(timestamp, nonce string) string {
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

// func getReqHeader(str string, r * http.Request) string{
//   r.ParseForm()
//   for k, v := range r.Form {
//     if str == k {
//        return strings.Join(v, "")
//     }
//   }
//   return " "
// }
func CheckSignature(w http.ResponseWriter, r * http.Request) ( string,  error) {
    signature := r.FormValue("signature")
  
    timestamp := r.FormValue("timestamp")
   
    nonce     := r.FormValue("nonce")
   
    echostr   := r.FormValue("echostr")
    if createSignature(timestamp, nonce) == signature {
        return echostr, nil
    }
    log.Println("signature isn't right")
    return "error signature", errors.New("error signature")
}
func handShakeGet(w http.ResponseWriter, r * http.Request) {
   
   if echostr, err := CheckSignature(w, r); err == nil {
      log.Println("write echostr to requestor")
      io.WriteString(w, echostr)
      
   }else{
      io.WriteString(w, "500")
   }
}

func NewResponse() (resp *Response) {
    resp = &Response{}
    resp.CreateTime = time.Duration(time.Now().Unix())
    return
}

func (resp Response) Encode() (data []byte, err error) {
    resp.CreateTime = time.Second
    data, err = xml.Marshal(resp)
    return
}

func DecodeRequest(data []byte) (req *Request, err error) {
   req = &Request{}
   if err = xml.Unmarshal(data, req); err != nil {
      return
   }
   req.CreateTime *= time.Second
   return 
}

func handlerUserReq(wreq * Request)(resp *Response, err error){
    log.Println("enter handlerUserReq")
    resp = NewResponse()
    resp.ToUserName = wreq.FromUserName
    resp.FromUserName = wreq.ToUserName
    resp.MsgType = Text
    log.Println("request msg type is ", wreq.MsgType)
    log.Println("request content is ", wreq.Content)
    if wreq.MsgType == Text {
        resp.Content = "愿得一心人，白首不相离。 每天一段情话送给你的爱人！ 天下有情人终成眷属.自动回复功能测试中."
    }
    err = nil
    return 
}
func handlerPost(w http.ResponseWriter, r * http.Request) {

    log.Println("entry handlerPost")
    if _, err := CheckSignature(w, r); err != nil{
        io.WriteString(w, "500")
        log.Println("handlerPost error signature")
        return
    }
    r.ParseForm()
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return 
    }
    log.Println(string(body))

    wechatrep, err := DecodeRequest([]byte(body))
    if err != nil {
        log.Println("DecodeRequest error")
        io.WriteString(w, "500")
        return
    }
    wresp, err := handlerUserReq(wechatrep)
    if err != nil{
        log.Println("handlerUserReq error")
        io.WriteString(w, "500")
        return
    }
    
    data, err := wresp.Encode()
    if  err != nil{
        log.Println("Encode error")
        io.WriteString(w, "500")
        return
    }
    io.WriteString(w, string(data))
}
func Weixinhandler( w http.ResponseWriter, r * http.Request) {
   // io.WriteString(w, "Hello, world!")
    // log.Println("method is :", r.Method)
    // r.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
    // //注意:如果没有调用ParseForm方法，下面无法获取表单的数据
    // log.Println(r.Form) //这些信息是输出到服务器端的打印信息
    // log.Println("path", r.URL.Path)
    // log.Println("scheme", r.URL.Scheme)
    log.Println("signature is ", r.FormValue("signature"))
    log.Println("nonce is ", r.FormValue("nonce"))
    log.Println("echostr is ", r.FormValue("echostr"))
    log.Println("timestamp is  ", r.FormValue("timestamp"))
    // for k, v := range r.Form {
    //     log.Println("key:", k)
    //     log.Println("val:", strings.Join(v, ""))
    // }

   if r.Method == "GET" {
      handShakeGet(w, r)
   }else {
      handlerPost(w, r)
   }
}