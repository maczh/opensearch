package opensearch

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func md5Encode(content string) (md string) {
	h := md5.New()
	_, _ = io.WriteString(h, content)
	md = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func getRandomIntString(l int) string {
	str := "0123456789"
	return generateRandString(str, l)
}

func generateRandString(source string, l int) string {
	bytes := []byte(source)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func getSignedHeader(cf Config, method string, uri, body string) map[string]string {
	header := make(map[string]string)
	header["Content-MD5"] = md5Encode(body)
	header["Content-Type"] = "application/json"
	if method == "GET" {
		header["Content-MD5"] = ""
	}
	header["Date"] = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	header["X-Opensearch-Nonce"] = getRandomIntString(10)
	signSrc := method + "\n" +
		header["Content-MD5"] + "\n" +
		header["Content-Type"] + "\n" +
		header["Date"] + "\n" +
		"x-opensearch-nonce:" + header["X-Opensearch-Nonce"] + "\n" +
		uri
	//if method == "GET" {
	//	signSrc = method + "\n" +
	//		header["Date"] + "\n" +
	//		"x-opensearch-nonce:" + header["X-Opensearch-Nonce"] + "\n" +
	//		uri
	//}
	logger.Debug("签名明文:" + signSrc)
	mac := hmac.New(sha1.New, []byte(cf.OS_SECRET_KEY))
	mac.Write([]byte(signSrc))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	header["Authorization"] = "OPENSEARCH " + cf.OS_ACCESS_KEY + ":" + signature
	fmt.Println(header)
	return header
}
