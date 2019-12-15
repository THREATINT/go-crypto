package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"

	"golang.org/x/crypto/ripemd160"
)

func dohash(h hash.Hash, data []byte) string {
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func MD5(data string) string {
	return dohash(md5.New(), []byte(data))
}

func SHA1(data string) string {
	return dohash(sha1.New(), []byte(data))
}

func SHA256(data string) string {
	return dohash(sha256.New(), []byte(data))
}

func SHA384(data string) string {
	return dohash(sha512.New384(), []byte(data))
}

func SHA512(data string) string {
	return dohash(sha512.New(), []byte(data))
}

func RIPEMD160(data string) string {
	return dohash(ripemd160.New(), []byte(data))
}
