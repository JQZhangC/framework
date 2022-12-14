package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"math/rand"
	"time"
)

type Crypto struct {
	key   []byte
	block cipher.Block
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func NewAES256(key string) (*Crypto, error) {

	ks := []byte(key)[:32]

	block, err := aes.NewCipher(ks)
	return &Crypto{block: block, key: ks}, err
}

func (c *Crypto) IVs() []byte {

	n := c.block.BlockSize()

	k := make([]byte, n)

	nano := time.Now().UnixNano()
	seed := rand.NewSource(nano)
	rnd := rand.New(seed)

	for i := 0; i < n; i++ {
		k[i] = byte(rnd.Intn(256))
	}

	return k
}

func (c *Crypto) Encode(src []byte) []byte {

	iv := c.IVs()
	src = append(iv, src...)

	k := PKCS5Padding(src, c.block.BlockSize())

	dst := make([]byte, len(k))

	enc := cipher.NewCBCEncrypter(c.block, iv)
	enc.CryptBlocks(dst, k)

	return dst
}
