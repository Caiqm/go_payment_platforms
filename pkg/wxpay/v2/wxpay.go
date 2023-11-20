package v2

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

var (
	ErrWxPemKeyNotFound = errors.New("wxpay: wxpay pem or key cert not found")
)

type Client struct {
	appId          string
	secret         string
	mchId          string
	mchSecret      string
	host           string
	SignType       string
	pemCert        []byte
	keyCert        []byte
	location       *time.Location
	client         *http.Client
	onReceivedData func(method string, data []byte)
}

type OptionFunc func(c *Client)

func WithApiHost(host string) OptionFunc {
	return func(c *Client) {
		if host != "" {
			c.host = host
		}
	}
}

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

// 设置商户号信息
func (c *Client) SetMchInformation(id, secret string) error {
	c.mchId = id
	c.mchSecret = secret
	c.SignType = "MD5"
	return nil
}

// 加载证书文件
func (c *Client) LoadAppCertPemKeyFromFile(pemCertPath, keyCertPath string) (err error) {
	p, err := os.ReadFile(pemCertPath)
	if err != nil {
		err = fmt.Errorf("wxpay: read pem cert fail, err = %s", err.Error())
		return
	}
	k, err := os.ReadFile(keyCertPath)
	if err != nil {
		err = fmt.Errorf("wxpay: read key cert fail, err = %s", err.Error())
		return
	}
	c.pemCert = p
	c.keyCert = k
	return
}

// 加载证书tls配置
func (c *Client) LoadTlsCertConfig() (tlsConfig *tls.Config, err error) {
	if len(c.pemCert) == 0 || len(c.keyCert) == 0 {
		err = ErrWxPemKeyNotFound
		return
	}
	var certificate tls.Certificate
	// 解析证书内容
	if certificate, err = tls.X509KeyPair(c.pemCert, c.keyCert); err != nil {
		err = fmt.Errorf("wxpay: tls.LoadX509KeyPair, %s", err.Error())
		return
	}
	tlsConfig = &tls.Config{
		Certificates:       []tls.Certificate{certificate},
		InsecureSkipVerify: true,
	}
	return
}

// 请求参数
func (c *Client) URLValues(param Param) (value url.Values, err error) {
	var values = url.Values{}
	values.Add(kFieldAppId, c.appId)
	var params = param.Params()
	for k, v := range params {
		if v == "" {
			continue
		}
		values.Add(k, v)
	}
	signature := c.sign(values)
	// 添加签名
	values.Add(kFieldSign, signature)
	return
}

// 生成签名
func (c *Client) sign(parameters url.Values) string {
	signStr := c.formatBizQueryParaMap(parameters)
	signStr = fmt.Sprintf("%s&key=%s", signStr, c.mchSecret)
	h := md5.New()
	h.Write([]byte(signStr))
	sign := hex.EncodeToString(h.Sum(nil))
	return strings.ToUpper(sign)
}

// 格式化参数，签名过程需要使用
func (c *Client) formatBizQueryParaMap(parameters url.Values) string {
	// 将key值提取出来
	var strs []string
	for k := range parameters {
		strs = append(strs, k)
	}
	// 排序
	sort.Strings(strs)
	// 赋值新map
	var signStr string
	var dot string
	for _, k := range strs {
		if parameters[k][0] == "" {
			continue
		}
		signStr += dot + k + "=" + parameters[k][0]
		dot = "&"
	}
	return signStr
}

// 产生随机字符串，不长于32位
func (c *Client) createNonceStr() string {
	length := 32
	strByte := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	bytes := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		bytes[i] = strByte[r.Intn(len(strByte))]
	}
	return string(bytes)
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
		xmlByte, _ := xml.Marshal(values)
		req.Body = io.NopCloser(bytes.NewBuffer(xmlByte))
	}
	// 添加header头
	req.Header.Add("Content-Type", kContentType)
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
}
