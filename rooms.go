package  main
import (
	"github.com/dustin/go-broadcvast"
)
type Message struct(
	 UserId string
	 RoomId string
	 Text string
)

type Listener struct{
	RoomId string
	Chan chan interface{}
}

type Manager struct {
	roomChannels map[string]broadcast.Broadcaster
	open chan*Listener
	close chan *Listener
	delete chan string
	messagues chan *Message
}

func NewRoomManager() *Manager{
	manager:= &Manager{
		roomChannels: make(map[string]broadcast.Broadcaster),
		open: make( *Listener ,100)
		close: make(*Listener ,100)
		delete :make()
	}
}