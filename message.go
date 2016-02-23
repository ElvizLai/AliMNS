/**
 * Created by elvizlai on 2016/2/16 16:57
 * Copyright © PubCloud
 */

package AliMNS

import (
	"encoding/xml"
	"fmt"
)

//SendMessage,BatchSendMessage,ReceiveMessage,BatchReceiveMessage,PeekMessage,BatchPeekMessage,DeleteMessage,BatchDeleteMessage,ChangeMessageVisibility

type MessageResponse struct {
	//XMLName   xml.Name `xml:"Message" json:"-"`
	Code           string   `xml:"Code,omitempty" json:"code,omitempty"`
	Message        string   `xml:"Message,omitempty" json:"message,omitempty"`
	RequestId      string   `xml:"RequestId,omitempty" json:"request_id,omitempty"`
	HostId         string   `xml:"HostId,omitempty" json:"host_id,omitempty"`
	MessageId      string `xml:"MessageId" json:"message_id"`
	MessageBodyMD5 string `xml:"MessageBodyMD5" json:"message_body_md5"`
}

type ErrorResponse struct {
	//XMLName   xml.Name `xml:"Error"`
	Code      string   `xml:"Code,omitempty" json:"code,omitempty"`
	Message   string   `xml:"Message,omitempty" json:"message,omitempty"`
	RequestId string   `xml:"RequestId,omitempty" json:"request_id,omitempty"`
	HostId    string   `xml:"HostId,omitempty" json:"host_id,omitempty"`
}

type message struct {
	XMLName      xml.Name    `xml:"Message"`
	MessageBody  base64Bytes `xml:"MessageBody"` //need base64 encode
	DelaySeconds int64       `xml:"DelaySeconds"`
	Priority     int64       `xml:"Priority"`
}

type messages struct {
	XMLName  xml.Name  `xml:"Messages"`
	Messages []message `xml:"Message"`
}

type MessageReceive struct {
	MessageId        string      `xml:"MessageId"`
	ReceiptHandle    string      `xml:"ReceiptHandle" json:"receipt_handle"`
	MessageBody      base64Bytes `xml:"MessageBody" json:"message_body"`
	MessageBodyMD5   string      `xml:"MessageBodyMD5" json:"message_body_md5"`
	EnqueueTime      int64       `xml:"EnqueueTime" json:"enqueue_time"`
	NextVisibleTime  int64       `xml:"NextVisibleTime" json:"next_visible_time"`
	FirstDequeueTime int64       `xml:"FirstDequeueTime" json:"first_dequeue_time"`
	DequeueCount     int64       `xml:"DequeueCount" json:"dequeue_count"`
	Priority         int64       `xml:"Priority" json:"priority"`
}

func NewMessage(content string, delaySeconds, priority int64) message {
	return message{MessageBody: []byte(content), DelaySeconds: delaySeconds, Priority: priority}
}

//send message to queue
func (c aliClient) SendMessage(queueName string, msg message) (MessageResponse, error) {
	msgXmlBytes, _ := xml.Marshal(&msg)
	path := fmt.Sprintf("/queues/%s/messages", queueName)
	m := MessageResponse{}
	return m, c.respHandler("POST", path, msgXmlBytes, &m)
}

func BatchSendMessage() {

}

func (c aliClient) ReceiveMessage(queueName string, waitSec int64) (MessageReceive, error) {
	path := ""
	if waitSec < 0 {
		path = fmt.Sprintf("/queues/%s/messages", queueName)
	} else {
		path = fmt.Sprintf("/queues/%s/messages?waitseconds=%v", queueName, waitSec)
	}
	msg := MessageReceive{}
	return msg, c.respHandler("GET", path, nil, &msg)
}

func BatchReceiveMessage() {

}

func PeekMessage() {

}

func BatchPeekMessage() {

}

func (c aliClient) DeleteMessage(queueName, receiptHandle string) error {
	path := fmt.Sprintf("/queues/%s/messages?ReceiptHandle=%s", queueName, receiptHandle)
	return c.respHandler("DELETE", path, nil, nil)
}

func BatchDeleteMessage() {

}

func ChangeMessageVisibility() {

}
