package packageSample

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// pwd出力
	currentPath, err := os.Getwd()
	fmt.Println(currentPath,err)

	// ls出力
    files, err1 := ioutil.ReadDir(".")
	if err1 != nil {
		log.Fatal("err1")
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}

	// osopenでもfileを開ける
	// ネットワーク関係のioに使われる
	// パケット読み込み→パケット書き出し
	content, err2 := ioutil.ReadFile("README.md")
	if err2 != nil {
		log.Fatal("err2", err2)
	}
	fmt.Println(string(content))

	if err := ioutil.WriteFile("README2.md",content, 0666); err != nil {
		log.Fatalln(err)
	}

	r := bytes.NewBuffer([]byte("abc"))
	content2, _ := ioutil.ReadAll(r)
	fmt.Println(string(content2))

}


