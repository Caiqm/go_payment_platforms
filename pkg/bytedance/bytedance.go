package bytedance

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

var (
	ErrByteDanceErrNoNotFound = errors.New("bytedance: err_no not found")
)

type Client struct {
	appId          string
	secret         string
	salt           string
	token          string
	host           string
	location       *time.Location
	client         *http.Client
	onReceivedData func(method string, data []byte)
}

type OptionFunc func(c *Client)

// 设置请求连接
func WithApiHost(host string) OptionFunc {
	return func(c *Client) {
		if host != "" {
			c.host = host
		}
	}
}

// 设置支付系统秘钥SALT
func WithSalt(salt string) OptionFunc {
	return func(c *Client) {
		if salt != "" {
			c.salt = salt
		}
	}
}

// 设置支付系统回调TOKEN
func WithToken(token string) OptionFunc {
	return func(c *Client) {
		if token != "" {
			c.token = token
		}
	}
}

// 初始化
func (c *Client) New(appId, secret string, opts ...OptionFunc) (nClient *Client) {
	nClient = &Client{}
	nClient.appId = appId
	nClient.secret = secret
	nClient.client = http.DefaultClient
	nClient.location = time.Local
	for _, opt := range opts {
		if opt != nil {
			opt(nClient)
		}
	}
	return
}

// 请求参数
func (c *Client) URLValues(param Param) (value url.Values, err error) {
	var values = url.Values{}
	// 是否需要APPID
	if param.NeedAppId() {
		values.Add(kFieldAppId, c.appId)
	}
	// 是否需要密钥
	if param.NeedSecret() {
		values.Add(kFieldSecret, c.secret)
	}
	// 结构体转MAP
	var params = c.structToMap(param)
	for k, v := range params {
		if v == "" {
			continue
		}
		values.Add(k, v.(string))
	}
	// 判断是否需要签名
	if param.NeedSign() {
		signature := c.sign(values)
		// 添加签名
		values.Add(kFieldSign, signature)
	}
	return
}

// 生成签名
func (c *Client) sign(paramsMap url.Values) string {
	var paramsArr []string
	for k, v := range paramsMap {
		if k == "other_settle_params" || k == "app_id" || k == "thirdparty_id" || k == "sign" {
			continue
		}
		value := strings.TrimSpace(fmt.Sprintf("%v", v[0]))
		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") && len(value) > 1 {
			value = value[1 : len(value)-1]
		}
		value = strings.TrimSpace(value)
		if value == "" || value == "null" {
			continue
		}
		paramsArr = append(paramsArr, value)
	}
	paramsArr = append(paramsArr, c.salt)
	sort.Strings(paramsArr)
	return fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(paramsArr, "&"))))
}

// 请求主方法
func (c *Client) doRequest(method string, param Param, result interface{}) (err error) {
	// 创建一个请求
	req, _ := http.NewRequest(method, c.host, nil)
	// 判断参数是否为空
	if param != nil {
		var values url.Values
		values, err = c.URLValues(param)
		if err != nil {
			return err
		}
		// 根据类型转换
		if param.ContentType() == kContentTypeJson {
			reqByte, _ := json.Marshal(values)
			req.Body = io.NopCloser(bytes.NewBuffer(reqByte))
		} else {
			req.PostForm = values
		}
		// 添加token头
		if param.NeedAccessToken() {
			req.Header.Add("access-token", values.Get(kFieldAccessToken))
		}
	}
	// 添加header头
	req.Header.Add("Content-Type", param.ContentType())
	// 发起请求数据
	rsp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	err = c.decode(bodyBytes, result)
	return
}

// 解密返回数据
func (c *Client) decode(data []byte, result interface{}) (err error) {
	var raw = make(map[string]json.RawMessage)
	if err = json.Unmarshal(data, &raw); err != nil {
		return
	}
	// 判断是否成功
	var errNBytes = raw[kFieldErrNo]
	if len(errNBytes) > 0 {
		var rErr *Error
		if err = json.Unmarshal(errNBytes, &rErr); err != nil {
			return
		}
		return rErr
	}
	if err = json.Unmarshal(data, result); err != nil {
		return
	}
	return
}

// 结构体转map
func (c *Client) structToMap(stu interface{}) map[string]interface{} {
	// 结构体转map
	m, _ := json.Marshal(&stu)
	var parameters map[string]interface{}
	_ = json.Unmarshal(m, &parameters)
	return parameters
}

// 返回内容
func (c *Client) OnReceivedData(fn func(method string, data []byte)) {
	c.onReceivedData = fn
}

// 验证回调签名
func (c *Client) VerifySign(values url.Values) bool {
	msgSignature := values.Get(kFieldMsgSignature)
	sortedString := []string{values.Get(kFieldTimestamp), values.Get(kFieldNonce), values.Get(kFieldMsg), c.token}
	sort.Strings(sortedString)
	h := sha1.New()
	h.Write([]byte(strings.Join(sortedString, "")))
	signature := fmt.Sprintf("%x", h.Sum(nil))
	return msgSignature == signature
}
