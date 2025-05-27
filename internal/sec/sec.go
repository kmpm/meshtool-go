package sec

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"errors"
)

var (
	deecryptionKeys = []string{
		"1PG7OiApB1nwvP+rz05pAQ==", // AQ== is an alias for this AES128 key
	}
)

// createNonce generates a nonce string based on the provided id and from values.
func createNonce(id, from uint32) []byte {
	packetID64 := uint64(id)
	blockCounter := uint32(0)
	// Create a buffer for the nonce
	nonce := make([]byte, 16)
	// Write packetId, fromNode, and block counter to the buffer
	binary.LittleEndian.PutUint64(nonce[0:8], packetID64)
	binary.LittleEndian.PutUint32(nonce[8:12], from)
	binary.LittleEndian.PutUint32(nonce[12:16], blockCounter)
	return nonce
}

func Decrypt(encrypted []byte, id, from uint32) ([]byte, error) {

	for _, keyBase := range deecryptionKeys {
		// decode base64 key into a byte slice
		key, err := base64.StdEncoding.DecodeString(keyBase)
		if err != nil {
			return nil, errors.Join(errors.New("error decoding key"), err)
		}

		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}
		// decrypt the encrypted string using the key and nonce created from id and from
		nonce := createNonce(id, from)
		stream := cipher.NewCTR(block, nonce)
		plaintext := make([]byte, len(encrypted))
		stream.XORKeyStream(plaintext, []byte(encrypted))
		return plaintext, nil
	}

	return nil, errors.New("no decryption keys available")
}

type StringDecoder interface {
	DecodeString(s string) ([]byte, error)
}

func DecryptString(dec StringDecoder, encryptedBase64 string, id, from uint32) ([]byte, error) {
	// decode base64 string into a byte slice
	encryptedBytes, err := dec.DecodeString(encryptedBase64)
	if err != nil {
		return nil, errors.Join(errors.New("error decoding encrypted string"), err)
	}
	return Decrypt(encryptedBytes, id, from)
}

func DecryptRawStd(encryptedBase64 string, id, from uint32) ([]byte, error) {
	// decode base64 string into a byte slice
	return DecryptString(base64.RawStdEncoding, encryptedBase64, id, from)
}

func DecryptStd(encryptedBase64 string, id, from uint32) ([]byte, error) {
	return DecryptString(base64.StdEncoding, encryptedBase64, id, from)
}

func mustBase64DecodeStd(s string) []byte {
	// decode base64 string into a byte slice
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return decoded
}

func mustBase64DecodeRawStd(s string) []byte {
	// decode base64 string into a byte slice
	decoded, err := base64.RawStdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return decoded
}
