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

// MD5 (data)
// return MD5-Hash of given string
func MD5(data string) string {
	return dohash(md5.New(), []byte(data))
}

// SHA1 (data)
// return SHA1-Hash of given string
func SHA1(data string) string {
	return dohash(sha1.New(), []byte(data))
}

// SHA256 (data)
// return SHA256-Hash of given string
func SHA256(data string) string {
	return dohash(sha256.New(), []byte(data))
}

// SHA384 (data)
// return SHA386-Hash of given string
func SHA384(data string) string {
	return dohash(sha512.New384(), []byte(data))
}

// SHA512 (data)
// return SHA512-Hash of given string
func SHA512(data string) string {
	return dohash(sha512.New(), []byte(data))
}

// RIPEMD160 (data)
// return RIPEMD160-Hash of given string
func RIPEMD160(data string) string {
	return dohash(ripemd160.New(), []byte(data))
}
