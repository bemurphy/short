package short

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"
)

func Shorten(url string) (string, error) {
	h := md5.New()
	h.Write([]byte(url))
	digest := hex.EncodeToString(h.Sum(nil))
	bi := big.NewInt(0)
	bi, ok := bi.SetString(digest, 16)
	if !ok {
		return "", errors.New("Big int weirdness")
	}

	buf := bytes.Buffer{}
	err := binary.Write(&buf, binary.BigEndian, uint32(bi.Uint64()))
	if err != nil {
		return "", err
	}

	s := base64.URLEncoding.EncodeToString(buf.Bytes())

	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "==", "", -2)

	return s, nil
}
