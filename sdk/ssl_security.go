package sdk

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	gmx509 "github.com/tjfoc/gmsm/x509"
)

type SSLSecurity struct {
	RootCerts   *x509.CertPool
	ClientCerts []tls.Certificate
}

func NewSSLSecurity(rootCertFile string, clientCertFile, clientKeyFile string) (*SSLSecurity, error) {
	rootCerts := x509.NewCertPool()
	rootPemData, err := ioutil.ReadFile(rootCertFile)
	if err != nil {
		return nil, err
	}
	rootCerts.AppendCertsFromPEM(rootPemData)
	clientCert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		return nil, err
	}
	return &SSLSecurity{
		RootCerts:   rootCerts,
		ClientCerts: []tls.Certificate{clientCert},
	}, nil
}

// 国密TLS 暂时仅支持忽略和单向认证
type GMSSLSecurity struct {
	RootCerts *gmx509.CertPool
}

func NewGMSSLSecurity(rootCertFile string) (*GMSSLSecurity, error) {
	rootPemData, err := ioutil.ReadFile(rootCertFile)
	if err != nil {
		return nil, err
	}
	rootCerts := gmx509.NewCertPool()
	rootCerts.AppendCertsFromPEM(rootPemData)

	return &GMSSLSecurity{
		RootCerts: rootCerts,
	}, nil
}
