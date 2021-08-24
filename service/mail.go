package service

import (
	"fmt"
	"os"
)

func Sendmail(email, context string) error {
	//暂时不急着实现
	if os.Getenv("test") != "True" {
		fmt.Println(email, context)
	}
	return nil
}
