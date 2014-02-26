package models

import (
    "github.com/astaxie/goredis"
    "strconv" 
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
    return
}
func DBInit() {
    client := &dbClient{}
    client.grClient.Addr = "127.0.0.1:6379"
    client.grClient.Db   = 13
    return
}

func getLetterIndex(c *dbClient, feild string ) (int, error){
    mutex.Lock()
    defer mutex.Unlock()
    index , err := c.grClient.Hincrby(LETTER_INDEX, feild, 1)
    return (int)(index), err
}


func (c *dbClient)AddSubscriber(subscriberKey string) error{
    ismem, err := c.grClient.Sismember(UNSUBSCRIBE_USER, []byte(subscriberKey))
    if err != nil {
        return err
    }
    if ismem {
       _, err =  c.grClient.Smove(UNSUBSCRIBE_USER, ALL_USER, []byte(subscriberKey))
       return err
    }

    _, err = c.grClient.Sadd(ALL_USER, []byte(subscriberKey))
    return err
}

func (c *dbClient) DelSubscriber(subscriberKey string) (err error){
    _, err =  c.grClient.Smove(ALL_USER,UNSUBSCRIBE_USER, []byte(subscriberKey))
    return
}

func (c *dbClient) AddForwardLetter(letter string) (err error){
    index , err := getLetterIndex(c, FW_LETTER_INDEX)
    key := strconv.Itoa(index)
    if err != nil {
        return 
    }
    _, err = c.grClient.Sadd(FW_LETTER, []byte(key))
    if err != nil {
        return
    }
    _, err = c.grClient.Hset(FW_LETTER, key, []byte(letter))
    return
}

func (c *dbClient) DelForwardLetter(index string) (err error) {
    _, err = c.grClient.Smove(FW_LETTER, SEND_FW_LETTER, []byte(index))
    return
}

func (c *dbClient) GetForwardLetterNum() ( int,  error)  {
    return c.grClient.Scard(FW_LETTER)
}

func (c *dbClient) AddOriginLetter(letter string)(err error) {
    index , err := getLetterIndex(c, ORIGINAL_LETTER_INDEX)
    key := (strconv.Itoa(index))
    if err != nil {
        return 
    }
    _, err = c.grClient.Sadd(ORIGINAL_LETTER, []byte(key))
    if err != nil {
        return
    }
    _, err = c.grClient.Hset(ORIGINAL_LETTER, key, []byte(letter))
    return
}

func (c dbClient)DelOriginLetter(index string) (err error) {
    _, err = c.grClient.Smove(ORIGINAL_LETTER, SEND_ORIGINAL_LETTER, []byte(index))
    return
}

func (c dbClient) GetOriginLetterNum()(num int, err error) {
    return c.grClient.Scard(ORIGINAL_LETTER)
}


