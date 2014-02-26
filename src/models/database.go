package models

import (
    "github.com/astaxie/goredis"
    "fmt"
    "sync"
)

const (
    ALL_USER                = "all_user:subscriber"
    ACTIVE_USER             = "active_user:subscriber"
    UNSUBSCRIBE_USER        = "unsubscribe_user:subscriber"

    FW_LETTER               = "all_fw_letter:content"
    SEND_FW_LETTER          = "send_fw_letter:content"
    ORIGINAL_LETTER         = "original_leter:content"
    SEND_ORIGINAL_LETTER    = "send_original_letter:content"

    LETTER_INDEX            = "letter_index"
    FW_LETTER_INDEX         = "fw_letter:index"
    ORIGINAL_LETTER_INDEX   = "original_leter:index"
)


var ( 
        mutex sync.Mutex
    )
type dbClient struct {
    grClient goredis.Client
}

func NewWebChatDbClient()(client *dbClient){
    client = &dbClient{}
}
func DBInit() {
    client = &dbClient{}
    client.Addr = "127.0.0.1:6379"
    client.Db   = 13
    return
}

func getLetterIndex(c *dbClient, feild string ) (int64, error){
    mutex.Lock()
    defer mutex.Unlock()
    return c.goredis.Hincrby(LETTER_INDEX, feild, 1)
}


func (c *dbClient)AddSubscriber(subscriberKey string) error{
    ismem, err := c.goredis.Sismember(UNSUBSCRIBE_USER, subscriberKey)
    if err != nil {
        return err
    }
    if ismem {
       _, err =  c.goredis.Smove(UNSUBSCRIBE_USER, ALL_USER, subscriberKey )
       return err
    }

    _, err = c.goredis.Sadd(ALL_USER, subscriber)
    return err
}

func (c *dbClient)DelSubscriber(subscriberKey string) err error{
    _, err =  c.goredis.Smove(ALL_USER,UNSUBSCRIBE_USER, subscriberKey )
}

func (c *dbClient)AddForwardLetter(letter string) err error{
    index , err := getLetterIndex(c, FW_LETTER_INDEX)
    if err != nil {
        return 
    }
    _, err = c.goredis.Sadd(FW_LETTER, index)
    if err != nil {
        return
    }
    _, err = c.goredis.HSet(FW_LETTER, index, letter)
    return
}

func (c *dbClient)DelForwardLetter(index int64) err error {
    _, err = c.goredis.Smove(FW_LETTER, SEND_FW_LETTER, index)
    return
}

func (c *dbClient)GetForwardLetterNum() ( int,  error)  {
    return c.goredis.Scard(FW_LETTER)
}

func (c dbClient)AddOriginLetter(letterKey string)err error {
    index , err := getLetterIndex(c, ORIGINAL_LETTER_INDEX)
    if err != nil {
        return 
    }
    _, err = c.goredis.Sadd(ORIGINAL_LETTER, index)
    if err != nil {
        return
    }
    _, err = c.goredis.HSet(ORIGINAL_LETTER, index, letter)
    return
}

func (c dbClient)DelOriginLetter(letterKey string) err error {
    _, err = c.goredis.Smove(ORIGINAL_LETTER, SEND_ORIGINAL_LETTER, index)
    return
}

func (c dbClient)GetOriginLetterNum()(num int, err error) {
    return c.goredis.Scard(ORIGINAL_LETTER)
}


