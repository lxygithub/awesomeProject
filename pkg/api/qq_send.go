package api

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/Logiase/gomirai/bot"
	"github.com/Logiase/gomirai/message"
)

func Send(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var qq uint
	var temp, err = strconv.ParseUint(r.URL.Query().Get("qq"), 10, 64)
	if err == nil {
		qq = uint(temp)
	} else {
		fmt.Fprintf(w, "qq is error")
		return
	}

	msg := r.URL.Query().Get("msg")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c := bot.NewClient("default", "http://124.70.68.221:8001", "12345678")
	c.Logger.Level = logrus.TraceLevel
	key, err := c.Auth()
	if err != nil {
		c.Logger.Fatal(err)
	}
	b, err := c.Verify(qq, key)
	if err != nil {
		c.Logger.Fatal(err)
	}
	//defer c.Release(qq)

	go func() {
		err = b.FetchMessages()
		if err != nil {
			c.Logger.Fatal(err)
		}
	}()

	for {
		select {
		case e := <-b.Chan:
			switch e.Type {
			case message.EventReceiveGroupMessage:
				_, err = b.SendGroupMessage(e.Sender.Group.Id, 0, message.PlainMessage(msg))
				if err != nil {
					fmt.Println(err)
				}
			}
		case <-interrupt:
			fmt.Println("######")
			fmt.Println("interrupt")
			fmt.Println("######")
			//c.Release(qq)
			c.Release(qq)
			return
		}

	}
}
