package test

import (
	"github.com/blockchain-jd-com/framework-go/crypto"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
)

/*
 * Author: imuge
 * Date: 2020/5/29 下午3:21
 */

var (
	// 网关服务IP
	GATEWAY_HOST = "127.0.0.1"
	// 网关服务端口
	GATEWAY_PORT = 8080
	// 是否建立安全连接
	SECURE = true
	// 是否是国密
	GM_SECURE = true
	// SSL配置
	SSL_ROOT_CERT        = "/home/imuge/jd/nodes/peer0/config/certs/root.crt"
	SSL_CLIENT_CERT      = ""
	SSL_CLIENT_KEY       = ""
	SSL_CLIENT_ENC_KEY   = "/home/imuge/jd/nodes/peer0/config/keys/gw0.enc.key"
	SSL_CLIENT_ENC_CERT  = "/home/imuge/jd/nodes/peer0/config/certs/tls/gw0.enc.crt"
	SSL_CLIENT_SIGN_KEY  = "/home/imuge/jd/nodes/peer0/config/keys/gw0.sign.key"
	SSL_CLIENT_SIGN_CERT = "/home/imuge/jd/nodes/peer0/config/certs/tls/gw0.sign.crt"

	// KEYPAIR身份模式，区块链上已存在的有操作权限的用户公私钥信息
	NODE_PRIVITE_KEY = crypto.DecodePrivKey(
		"177gjyiJjdZNfEu4kgR97BftoUtJ54ixiQZS9uktxqtDue6bGwwBLz4hEXw7Gu5fgGpGceG",
		base58.MustDecode("AXhhKihAa2LaRwY5mftnngSPKDF4N9JignnQ4skynY8y"))
	NODE_PUBLIC_KEY = crypto.DecodePubKey("7VeRKo8hU8mYfvGUGbNuCrfDqSY6PqgBt1dWWJwP6ofJnzkR")

	// CA身份模式，区块链上已存在的有操作权限的用户证书和私钥信息
	// 	cert, _             = ca.RetrieveCertificate("-----BEGIN CERTIFICATE-----\nMIIB4DCCAYagAwIBAgIENhE1ZTAKBggqhkjOPQQDAjBwMQwwCgYDVQQKDANKRFQxDTALBgNVBAsM\nBFJPT1QxCzAJBgNVBAYTAkNOMQswCQYDVQQIDAJCSjELMAkGA1UEBwwCQkoxDTALBgNVBAMMBHJv\nb3QxGzAZBgkqhkiG9w0BCQEWDGltdWdlQGpkLmNvbTAeFw0yMTEwMjcwODQ3MDdaFw0zMTEwMjUw\nODQ3MDdaMHExDDAKBgNVBAoMA0pEVDENMAsGA1UECwwEUEVFUjELMAkGA1UEBhMCQ04xCzAJBgNV\nBAgMAkJKMQswCQYDVQQHDAJCSjEOMAwGA1UEAwwFcGVlcjAxGzAZBgkqhkiG9w0BCQEWDGltdWdl\nQGpkLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLFhLigz1Rpd1rahUmlLiatzhYgnQtVP\nyZApmn42oWiEFMa68xaQb5jV6YLrikLK1EzyZDHLZBEoD9iS6ad7KqqjDTALMAkGA1UdEwQCMAAw\nCgYIKoZIzj0EAwIDSAAwRQIgBllErLVMu5qG6kpEyvY1rWmeVn+4SzhrH3w8+dPHlqQCIQC2Cf86\nBl/6zHUzsOZdbbXOjv6cuUh6VwO60HeKgAHQeg==\n-----END CERTIFICATE-----")
	// 	NODE_PUBLIC_KEY     = ca.RetrievePubKey(cert)
	// 	NODE_PRIVITE_KEY, _ = ca.RetrievePrivKey(cert.Algorithm, "-----BEGIN EC PARAMETERS-----\nBggqhkjOPQMBBw==\n-----END EC PARAMETERS-----\n-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIONhm+hn9bN6drxV2b6xeIBfFoBS95AKSPn0y8v7ryuYoAoGCCqGSM49\nAwEHoUQDQgAEsWEuKDPVGl3WtqFSaUuJq3OFiCdC1U/JkCmafjahaIQUxrrzFpBv\nmNXpguuKQsrUTPJkMctkESgP2JLpp3sqqg==\n-----END EC PRIVATE KEY-----")

	NODE_KEY = ledger_model.NewBlockchainKeypair(NODE_PUBLIC_KEY, NODE_PRIVITE_KEY)
)
