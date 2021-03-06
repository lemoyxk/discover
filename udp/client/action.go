/**
* @program: discover
*
* @description:
*
* @author: lemo
*
* @create: 2021-02-26 11:57
**/

package client

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/lemoyxk/console"
	"github.com/lemoyxk/kitty/socket"
	"github.com/lemoyxk/kitty/socket/udp/client"

	"discover/app"
	"discover/message"
)

func SendWhoIsMaster() {
	for i := 0; i < 10; i++ {
		console.AssertError(
			app.Node.Client.ProtoBufEmit(socket.ProtoBufPack{
				Event: "/WhoIsMaster",
				Data: &message.WhoIsMaster{
					Addr:      app.Node.Addr,
					Timestamp: app.Node.StartTime.UnixNano(),
					IsMaster:  app.Node.IsMaster(),
				},
			}),
		)

		time.Sleep(time.Millisecond * 100)
	}
}

func WhoIsMaster(c *client.Client, stream *socket.Stream) error {
	var data message.WhoIsMaster
	var err = proto.Unmarshal(stream.Data, &data)
	if err != nil {
		return err
	}

	app.Node.ServerMap.Set(data.Addr.Addr, &data)

	return nil
}
