package testmod

import "fmt"


//go mod init github.com/tallongsun/testmod

//git init
//git add .
//git commit -m "first commit"
//git remote add origin https://github.com/tallongsun/testmod.git
//git push -u origin master

//git tag v1.0.0
//git push --tags

func Hi(name string) string {
	return fmt.Sprintf("Hi, %s!",name)
}
