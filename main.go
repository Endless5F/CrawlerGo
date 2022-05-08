package main

import (
	"GoCodeProject/crawler/engine"
	"GoCodeProject/crawler/scheduler"
	"GoCodeProject/crawler/zhenai/parser"
)

func main() {
	// engine.SimpleEngine{}.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })

	e := engine.ConcurrentEngine{
		Shceduler:   &scheduler.SimpleSheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
