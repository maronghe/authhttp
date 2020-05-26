package main

import (
	"authhttp/log"
	"authhttp/router"
	"runtime/debug"
	"time"

	"github.com/astaxie/beego"
)

//func init() {
//	log.SetFormatter(&log.JSONFormatter{})
//	log.SetOutput(os.Stdout)
//
//	// Only log the warning severity or above.
//	log.SetLevel(log.WarnLevel)
//}

func main() {

	router.Init()

	log.Init()

	go f()

	beego.Run()

}

func f() {
	time.Sleep(5 * time.Second)
	defer func() {
		if err := recover(); nil != err {
			//debug.PrintStack()
			log.Sugar.Warnw("Panic Recover ",
				"err", err,
				"stack", string(debug.Stack()),
			)
		}
	}()

	log.Sugar.Info("Hello, I'm a goroutine")

	panic("Go down")

	log.Sugar.Info("Hello, I'm a goroutine")
}
