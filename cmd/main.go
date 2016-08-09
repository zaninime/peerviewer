package main

import (
	"fmt"
	"time"

	"github.com/ziutek/gst"
)

func main() {
	fmt.Println("Hello")
	src := gst.ElementFactoryMake("videotestsrc", "Test source")
	src.SetProperty("do-timestamp", true)
	src.SetProperty("pattern", 18)
	sink := gst.ElementFactoryMake("xvimagesink", "sink")
	//http.ListenAndServe(":8080", initHTTP())
	pipe := gst.NewPipeline("mypipe")
	pipe.Add(src, sink)
	pipe.SetState(gst.STATE_PLAYING)
	src.Link(sink)
	time.Sleep(10 * time.Minute)
}
