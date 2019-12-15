package crypto

import (
	"testing"
)

func TestMD5(t *testing.T) {
	if MD5("Hello World") != "b10a8db164e0754105b7a99be72e3fe5" {
		t.Error("wrong MD5!")
	}
}

func TestSHA1(t *testing.T) {
	if SHA1("Hello World") != "0a4d55a8d778e5022fab701977c5d840bbc486d0" {
		t.Error("wrong SHA1!")
	}
}

func TestSHA256(t *testing.T) {
	if SHA256("Hello World") != "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e" {
		t.Error("wrong SHA2565!")
	}
}

func TestSHA512(t *testing.T) {
	if SHA512("Hello World") !=
		"2c74fd17edafd80e8447b0d46741ee243b7eb74dd2149a0ab1b9246fb30382f27e853d8585719e0e67cbda0daa8f51671064615d645ae27acb15bfb1447f459b" {
		t.Error("wrong SHA5125!")
	}
}
