package main

import "fmt"

func connectDB() {
	panic("Unable to connect to database!")
}

func save() {
	defer func() {
		// panicのリカバリー
		// panicの引数を返す
		log := recover()
		fmt.Println(log)
	}()
	connectDB()
	/* deferじゃないとリカバリできない
	log := recover()
	fmt.Println(log)
	 */
}

func main() {
	save()
	fmt.Println("OK")
}

