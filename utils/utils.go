package utils

import "golang.org/x/crypto/bcrypt"

func Hashtool(key string) (string, error) {
	//工具的话还是返回错误比较好
	var err error
	hashedkey, err := bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
	//这里用现成的加密包bcrypt
	if err != nil {
		return "", err
	}
	return string(hashedkey), nil
}

func ComparePassword(dbPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
