package encrypt

import (
	"crypto/sha256"
	"github.com/JQZhangC/framework/pkg/crypto"
	"io/ioutil"
	"os"
)

func EncryptConfig(key, inputFilePath, outputFilePath string) error {
	hash := sha256.New()
	hash.Write([]byte(key))
	secret := hash.Sum(nil)

	conf := inputFilePath

	bs, err := ioutil.ReadFile(conf)

	if err != nil {
		return err
	}

	c, err := crypto.NewAES256(string(secret))

	if err != nil {
		return err
	}

	ps := c.Encode(bs)
	ps = append(secret, ps...)

	err = ioutil.WriteFile(outputFilePath, ps, 0644)
	if err != nil {
		return err
	}

	os.Remove(conf)
	return nil
}
