package sm

import (
	"errors"
	"github.com/ZZMarquis/gm/sm3"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
)

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:46 下午
 */

var (
	SM3_DIGEST_BYTES  = 256 / 8
	SM3_DIGEST_LENGTH = framework.ALGORYTHM_CODE_SIZE + SM3_DIGEST_BYTES
)

var _ framework.HashFunction = (*SM3HashFunction)(nil)

type SM3HashFunction struct {
}

func (S SM3HashFunction) GetAlgorithm() framework.CryptoAlgorithm {
	return SM3_ALGORITHM
}

func (S SM3HashFunction) Hash(data []byte) *framework.HashDigest {
	digestBytes := sm3.Sum(data)
	return framework.NewHashDigest(S.GetAlgorithm(), digestBytes[:])
}

func (S SM3HashFunction) Verify(digest *framework.HashDigest, data []byte) bool {
	hashDigest := S.Hash(data)
	return bytes.Equals(hashDigest.ToBytes(), digest.ToBytes())
}

func (S SM3HashFunction) SupportHashDigest(digestBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+摘要长度，以及算法标识；
	return S.GetAlgorithm().Match(digestBytes, 0) && SM3_DIGEST_LENGTH == len(digestBytes)
}

func (S SM3HashFunction) ParseHashDigest(digestBytes []byte) (*framework.HashDigest, error) {
	if S.SupportHashDigest(digestBytes) {
		return framework.ParseHashDigest(digestBytes)
	} else {
		return nil, errors.New("invalid digestBytes!")
	}
}
