package v2

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
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
	ErrWxErrCodeNotFound    = errors.New("wxpay: errcode not found")
	ErrWxReturnCodeNotFound = errors.New("wxpay: return_code not found")
	ErrWxPemKeyNotFound     = errors.New("wxpay: wxpay pem or key cert not found")
)

type Client struct {
	appId          string
	secret         string
	mchId          string
	mchSecret      string
	host           string
	signType       string
	pemCert        []byte
	keyCert        []byte
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

// 设置支付请求连接
func WithPayHost() OptionFunc {
	return func(c *Client) {
		c.host = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	}
}

// 设置申请退款请求连接
func WithRefundHost() OptionFunc {
	return func(c *Client) {
		c.host = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	}
}

// 设置商户号信息
func WithMchInformation(id, secret string) OptionFunc {
	return func(c *Client) {
		if id != "" {
			c.mchId = id
		}
		if secret != "" {
			c.mchSecret = secret
		}
		c.signType = "MD5"
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

// 加载证书文件
func (c *Client) LoadAppCertPemKeyFromFile(pemCertPath, keyCertPath string) (err error) {
	p, err := os.ReadFile(pemCertPath)
	if err != nil {
		err = fmt.Errorf("wxpay: read pem cert fail, err = %s", err.Error())
		return
	}
	k, err := os.ReadFile(keyCertPath)
	if err != nil {
		err = fmt.Errorf("wxpay: read key cert fail, %s", err.Error())
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
	// 是否需要APPID
	if param.NeedAppId() {
		values.Add(kFieldAppId, c.appId)
	}
	// 是否需要密钥
	if param.NeedSecret() {
		values.Add(kFieldSecret, c.secret)
	}
	var params = c.structToMap(param)
	for k, v := range params {
		if v == "" {
			continue
		}
		values.Add(k, v)
	}
	// 判断是否需要签名
	if param.NeedSign() {
		values.Add(kFieldMchId, c.mchId)
		values.Add(kFieldNonceStr, c.createNonceStr())
		values.Add(kFieldSignType, c.signType)
		signature := c.sign(values)
		// 添加签名
		values.Add(kFieldSign, signature)
	}
	return
}

// 结构体转map
func (c *Client) structToMap(stu interface{}) map[string]string {
	// 结构体转map
	m, _ := json.Marshal(&stu)
	var parameters map[string]string
	_ = json.Unmarshal(m, &parameters)
	return parameters
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
		if parameters[k][0] == "" || k == kFieldSign {
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
		if method == http.MethodPost {
			// 根据类型转换
			var reqByte []byte
			if strings.ToLower(param.ReturnType()) == "json" {
				reqByte, _ = json.Marshal(values)
			} else {
				reqByte, _ = xml.Marshal(values)
			}
			req.Body = io.NopCloser(bytes.NewBuffer(reqByte))
		} else if method == http.MethodGet {
			req.Host = c.host + "?" + values.Encode()
		}
	}
	// 是否需要证书
	if param.NeedTlsCert() {
		tlsConfig, err1 := c.LoadTlsCertConfig()
		if err1 != nil {
			err = err1
			return
		}
		c.client.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
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
	err = c.decode(bodyBytes, param.ReturnType(), param.NeedVerify(), result)
	return
}

// 解密返回数据
func (c *Client) decode(data []byte, returnType string, needVerifySign bool, result interface{}) (err error) {
	if strings.ToLower(returnType) == "json" || returnType == "" {
		var raw = make(map[string]json.RawMessage)
		if err = json.Unmarshal(data, &raw); err != nil {
			return
		}
		// 判断是否成功
		var errNBytes = raw[kFieldErrCode]
		if len(errNBytes) > 0 {
			var aErr *AppletError
			if err = json.Unmarshal(data, &aErr); err != nil {
				return
			}
			return aErr
		}
		if err = json.Unmarshal(data, result); err != nil {
			return
		}
	} else {
		var tmpResult interface{}
		if err = xml.Unmarshal(data, &tmpResult); err != nil {
			return
		}
		// 判断是否成功
		var resultMap = tmpResult.(map[string]interface{})
		if _, has := resultMap[kFieldReturnCode]; !has {
			return ErrWxReturnCodeNotFound
		}
		if resultMap[kFieldReturnCode].(ReturnCode) != ReturnCodeSuccess {
			var pErr *PayError
			if err = json.Unmarshal(data, &pErr); err != nil {
				return
			}
			return pErr
		}
		// 校验签名
		if needVerifySign {
			params := make(url.Values)
			for key, value := range resultMap {
				strValue := fmt.Sprintf("%v", value)
				params.Add(key, strValue)
			}
			// 验证签名
			if err = c.VerifySign(params); err != nil {
				return
			}
		}
		// 参数绑定
		if err = xml.Unmarshal(data, result); err != nil {
			return
		}
	}
	return
}

// 返回内容
func (c *Client) OnReceivedData(fn func(method string, data []byte)) {
	c.onReceivedData = fn
}

// 验证签名
func (c *Client) VerifySign(values url.Values) (err error) {
	verifier := values.Get(kFieldSign)
	compareSign := c.sign(values)
	if strings.Compare(verifier, compareSign) != 0 {
		err = fmt.Errorf("验证签名失败，接口返回签名：%s，生成签名：%s", compareSign, verifier)
		return
	}
	return
}
