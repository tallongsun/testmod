package testmod

import (
	"errors"
	"fmt"
)


//go mod init github.com/tallongsun/testmod

//git init
//git add .
//git commit -m "first commit"
//git remote add origin https://github.com/tallongsun/testmod.git
//git push -u origin master

//git tag v1.0.0
//git push --tags

//git checkout -b v1
//git push -u origin v1
//git commit -m "Emphasize our friendliness" testmod.go
//git tag v1.0.1
//git push --tags origin v1


func Hi(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "pt":
		return fmt.Sprintf("Oi, %s!", name), nil
	case "es":
		return fmt.Sprintf("¡Hola, %s!", name), nil
	case "fr":
		return fmt.Sprintf("Bonjour, %s!", name), nil
	default:
		return "", errors.New("unknown language")
	}
}
