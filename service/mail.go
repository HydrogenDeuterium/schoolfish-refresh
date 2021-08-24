package service

import "fmt"

func Sendmail(email, context string) error {
	//暂时不急着实现
	fmt.Println(email, context)
	return nil
}
