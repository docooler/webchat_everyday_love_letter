package models

import (
    "github.com/astaxie/goredis"
    "fmt"
)

const (
    ALL_USER                = "all_user:subscriber"
    ACTIVE_USER             = "active_user:subscriber"
    UNSUBSCRIBE_USER        = "unsubscribe_user:subscriber"

    FW_LETTER               = "all_fw_letter:content"
    SEND_FW_LETTER          = "send_fw_letter:content"
    ORIGINAL_LETTER         = "original_leter:content"
    SEND_ORIGINAL_LETTER    = "send_original_letter:content"

    FW_LETTER_INDEX         = "fw_letter:index"
    ORIGINAL_LETTER_INDEX   = "original_leter:index"
)
type dbClient struct {
    grClient goredis.Client
}

func NewWebChatDbClient()(client *dbClient){
    client = &dbClient{}
    client.Addr = "127.0.0.1:6379"
    val, err = client.Get(FW_LETTER_INDEX)
    if err != nil {
        client.Set(FW_LETTER_INDEX, 0)
        val = 0
    }
    fmt.Println("fw index is ", val)

    _, err = client.Get(ORIGINAL_LETTER_INDEX)
    if err != nil {
        client.Set(ORIGINAL_LETTER_INDEX, 0)
        val = 0
    }
    fmt.Println("origin index is ", val)

    return
}
func dbInit() {
    
}

func (client dbClient)AddSubscriber(subscriberKey string) err error{
    err = nil
    return
}

func (client dbClient)DelSubscriber(subscriberKey string) err error{
    err = nil
    return
}

func (client dbClient)AddForwardLetter(letter string) err error{
    err = nil
    return
}

func (client dbClient)DelForwardLetter(letterKey string) err error {
    err = nil
    return
}

func (client dbClient)GetForwardLetterNum() (num int, err error)  {
    num = 0
    err = nil
    return
}

func (client dbClient)AddOriginLetter(letterKey string)err error {
    err = nil
    return
}

func (client dbClient)DelOriginLetter(letterKey string) err error {
    err = nil
    return
}

func (client dbClient)GetOriginLetterNum()(num int, err error) {
    err = nil
    return
}


