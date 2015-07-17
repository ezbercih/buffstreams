package main

import (
	"strconv"
	"time"

	"github.com/StabbyCutyou/buffstreams"
	"github.com/StabbyCutyou/buffstreams/test/message"
	"github.com/golang/protobuf/proto"
)

// TestCallback is a simple server for test purposes. It has a single callback,
// which is to unmarshall some data and log it.
func TestCallback(bts []byte) error {
	msg := &message.Note{}
	err := proto.Unmarshal(bts, msg)
	return err
}

func main() {
	cfg := buffstreams.BuffTCPListenerConfig{
		MaxMessageSize: 2048,
		EnableLogging:  true,
		Address:        buffstreams.FormatAddress("", strconv.Itoa(5031)),
		Callback:       TestCallback,
	}

	btl := buffstreams.NewBuffTCPListener(cfg)
	btl.StartListeningAsync()
	// Need to block until ctrl+c, but having trouble getting signal trapping to work on OSX...
	time.Sleep(time.Minute * 10)
}
