package encrypt

import "fmt"

func EncryptConfig(key, inputFilePath, outputFilePath string) error {
	fmt.Println(inputFilePath, outputFilePath, key)
	//hash := sha256.New()
	//hash.Write([]byte(key))
	//secret := hash.Sum(nil)
	//
	//conf := inputFilePath
	//
	//bs, err := ioutil.ReadFile(conf)
	//
	//if err != nil {
	//	return err
	//}
	//
	//c, err := crypto.NewAES256(string(secret))
	//
	//if err != nil {
	//	return err
	//}
	//
	//ps := c.Encode(bs)
	//ps = append(secret, ps...)
	//
	//err = ioutil.WriteFile(outputFilePath, ps, 0644)
	//if err != nil {
	//	return err
	//}
	//
	//os.Remove(conf)
	return nil
}
