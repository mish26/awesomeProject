package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./log.go")
	if err != nil {
		log.Fatalln("fileOpenError!")
	}
	defer file.Close()
	data := make([]byte,100)
	// countがイニシャライズされているから、エラーにならない
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("fileReadError!")
	}
	fmt.Println(count, string(data))

	// :=(イニシャライズだとエラーになる)
	err = os.Chdir("test")
	if err != nil {
		log.Fatalln("chDirError!")
	}

	// 上記の省略形
	// 一つ師からerrを返さない場合によく使う書き方
	if 	err = os.Chdir("test"); err != nil {
		log.Fatalln("chDirError!")
	}

}

