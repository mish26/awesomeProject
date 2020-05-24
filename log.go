package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	loggingSettings("sample.log")
	log.Println("logging!")
	log.Printf("%T %v","testlog","testllll")

	textFile := "hogefuga.txt"
	_,err := os.Open(textFile)
	if err != nil{
		log.Fatalln(textFile,"is nothing: ", err)
	}

	// Fatalを使うとプログラムの実行が終了する。sucessは出力されない。
	log.Fatalln("error!!")
	fmt.Println("success!!!")
}

func loggingSettings(logFile string) {
	// どういったモードでlogファイルを開くか
	logFileSetting, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// ターミナルに出力、logファイルへの出力
    multiLogFile := io.MultiWriter(os.Stdout, logFileSetting)
	// multiLogFile := io.MultiWriter(logFileSetting)
	// logの出力設定
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
