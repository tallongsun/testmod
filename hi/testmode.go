package testmod

import "fmt"


//go mod init github.com/tallongsun/testmod

func Hi(name string) string {
	return fmt.Sprintf("Hi, %s",name)
}
