package mongodb

import (
	"fmt"
	"os"
)

func initFileFromBase64String(filename, value string) (*os.File, error) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("mongodb.initFileFromBase64String - err: ", err)
		return nil, err
	}
	b := base64DecodeToBytes(value)
	if _, err := f.Write(b); err != nil {
		fmt.Println("mongodb.initFileFromBase64String - write file err: ", err)
		return nil, err
	}
	f.Sync()
	return f, nil
}
