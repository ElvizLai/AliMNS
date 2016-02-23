###AliMNS Go SDK

####Usage:
```go
// you own config
const (
	accessKeyId = "xxxxx"
	accessKeySecret = "xxxxx"
	url = "http://{id}.mns.{region}.aliyuncs.com"
)

func main() {
	//create an AliMNS client
	client := AliMNS.NewClient(accessKeyId, accessKeySecret, url)
	
	//list all Queues you have created
	fmt.Println(client.ListQueue())

	//send msg to queue called msgQueue with delay 0 and priority 8
	fmt.Println(client.SendMessage("msgQueue", AliMNS.NewMessage("你好, Elvizlai", 0, 8)))

	//receive with 30 seconds block and delete it if succeed, if not msg received, will return error: Message not exist.
	for {
		msg, err := client.ReceiveMessage("msgQueue", 30)
		if err != nil {
			fmt.Println(err)
		}else {
			fmt.Println(string(msg.MessageBody))
			fmt.Println(client.DeleteMessage("msgQueue", msg.ReceiptHandle))
		}
	}
}

```