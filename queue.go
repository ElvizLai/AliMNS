/**
 * Created by elvizlai on 2016/2/17 14:31
 * Copyright © PubCloud
 */

package AliMNS

import "encoding/xml"

type queue struct {
	QueueURL string `xml:"QueueURL" json:"url"`
}

type queues struct {
	XMLName xml.Name `xml:"Queues" json:"-"`
	Queues  []queue `xml:"Queue" json:"queues"`
}

//CreateQueue,DeleteQueue,ListQueue,GetQueueAttributes,SetQueueAttributes

func CreateQueue() {

}

func DeleteQueue() {

}

func (c aliClient) ListQueue() ([]queue, error) {
	qs := queues{}
	return qs.Queues, c.respHandler("GET", "/queues", nil, &qs)
}

func GetQueueAttributes() {

}

func SetQueueAttributes() {

}
