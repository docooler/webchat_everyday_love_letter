package main 
import (
            "log"
            "controllers"
            "net/http"
            )

// var client *wechatdb.dbClient


func main() {
   
   err := http.ListenAndServe(":80", nil)
   if err != nil {
      log.Fatal("God like")
   }
   log.Println("run server at 80")
}

func init(){
    http.HandleFunc("/weixin", controllers.Weixinhandler)
    controllers.TestFunc()
}
