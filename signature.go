/**
 * Created by elvizlai on 2016/2/17 11:39
 * Copyright © PubCloud
 */

package AliMNS

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"sort"
	"strings"
	"time"
)

/**
签名机制

Access Key ID和Access Key Secret 由阿里云官方颁发给访问者(可以通过阿里云官方网站申请和管理 ),其中 Access Key ID 用于标识访问者的身份;Access Key Secret 是用于加密签名字符串和服务器端 验证签名字符串的密钥,必须严格保密,只有阿里云和用户知道。
MNS 服务会对每个访问的请求进行验证,每个向 MNS 提交的请求,都需要在请求中包含签名 (Signature)信息。MNS 通过使用 Access Key ID 和 Access Key Secret 进行对称加密的方法来验证 请求的发送者身份。如果计算出来的验证码和提供的一样即认为该请求是有效的;否则,MNS 将拒绝处 理这次请求,并返回 HTTP 403 错误。
用户可以在 HTTP 请求中增加 Authorization(授权)的 Head 来包含签名信息,表明这个消息已被授权 。

MNS 要求将签名包含在 HTTP Header 中,格式如下: Authorization: MNS AccessKeyId:Signature

Signature计算方法如下:

Signature = base64(hmac-sha1(VERB + "\n" + CONTENT-MD5 + "\n"
+ CONTENT-TYPE + "\n"
+ DATE + "\n"
+ CanonicalizedMNSHeaders
+ CanonicalizedResource))

● VERB 表示HTTP的Method(如:PUT,POST)
● Content-Md5 表示请求内容数据的MD5值(需base64,参考 https://tools.ietf.org/html/rfc1864)
● CONTENT-TYPE 表示请求内容的类型
● DATE 表示此次操作的时间,不能为空,目前只支持GMT格式,如果请求时间和MNS服务器时间相差 超过15分钟,MNS会判定此请求不合法,返回400错误,错误信息及错误码详见本文档第5部分。 (如示例中:Thu, 17 Mar 2012 18:49:58 GMT)
● CanonicalizedMNSHeaders表示 http中的x-mns-开始的字段组合。(见下文注意事项)
● CanonicalizedResource表示http所请求资源的URI(统一资源标识符)。(如示例中:/queues/$queueName?metaOverride=true)

注意:
CanonicalizedMNSHeaders(即x-mns-开头的head)在签名验证前需要符合以下规范:
● head的名字需要变成小写
● head自小到大排序
● 分割head name和value的冒号前后不能有空格
● 每个Head之后都有一个\n,如果没有以x-mns-开头的head,则在签名时CanonicalizedMNSHeaders就设置为空

其他:
1. 用来签名的字符串为UTF-8格式。
2. 签名的方法用RFC 2104中定义的HMAC-SHA1方法,其中Key为AccessKeySecret。
3. content-type和content-md5在请求中不是必须的,没有情况下,请使用''代替。

**/
func signature(key string, method string, headers map[string]string, resource string) (signature string, err error) {
	contentMD5 := ""
	contentType := ""
	date := time.Now().UTC().Format(http.TimeFormat)

	if v, exist := headers[_CONTENT_MD5]; exist {
		contentMD5 = v
	}

	if v, exist := headers[_CONTENT_TYPE]; exist {
		contentType = v
	}

	if v, exist := headers[_DATE]; exist {
		date = v
	}

	mnsHeaders := []string{}

	for k, v := range headers {
		if strings.HasPrefix(k, "x-mns-") {
			mnsHeaders = append(mnsHeaders, k+":"+strings.TrimSpace(v))
		}
	}

	sort.Sort(sort.StringSlice(mnsHeaders))

	stringToSign := string(method) + "\n" +
		contentMD5 + "\n" +
		contentType + "\n" +
		date + "\n" +
		strings.Join(mnsHeaders, "\n") + "\n" +
		resource

	sha1Hash := hmac.New(sha1.New, []byte(key))
	if _, e := sha1Hash.Write([]byte(stringToSign)); e != nil {
		panic(e)
		return
	}

	signature = base64.StdEncoding.EncodeToString(sha1Hash.Sum(nil))

	return
}
