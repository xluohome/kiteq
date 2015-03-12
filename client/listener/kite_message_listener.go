package listener

import (
	"kiteq/protocol"
	"log"
)

type IListener interface {
	//接受投递消息的回调
	OnMessage(msg *protocol.QMessage) bool
	//接收事务回调
	// 除非明确提交成功、其余都为不成功
	// 有异常或者返回值为false均为不提交
	OnMessageCheck(tx *protocol.TxResponse) error
}

type MockListener struct {
}

func (self *MockListener) OnMessage(msg *protocol.QMessage) bool {
	log.Println("MockListener|OnMessage", msg.GetHeader(), msg.GetBody())
	return true
}

func (self *MockListener) OnMessageCheck(tx *protocol.TxResponse) error {
	log.Printf("MockListener|OnMessageCheck|%s\n", tx.MessageId)
	v, _ := tx.GetProperty("tradeno")
	log.Printf("MockListener|OnMessageCheck|PROP|%s\n", v)
	tx.Commit()
	return nil
}
