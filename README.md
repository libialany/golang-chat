```
example of broadcast

pacakge main
import "log"
func main(){
b:=NewBroadcaster(100)
workerOne(b)
workerTwo(b)
for i:=0;i<5;i++{
log.Printf("Sending %v",i)
b.Submit(i)
}
b.Close()
}
func workerOne(b Broadcaster){
ch:= make(chan interface{})
b.Register(ch)
defer b.Unregister(ch)
go func(){
for v:=range ch {
log.Printf("workOne read %v",v)
}}}

func workerTwo(b Broadcaster){
ch:=make(chan interface{})
b.Register(ch)
defer b.Unregister(ch)
defer log.Printf("workerTwo is done\n")
go func(){
log.Printf("workerTwo read %v \n" <-ch)
}()
}
```

Why there is a broadcast example ?
Because its register the connection and this is part of Manager structure and other like New RomManager.

> main.go

1. **newroommanager**

2. **ginconnection**

3. set entrypoints

4. POST,DELETE,GET -> roomManger

5. static function roomid,listener,room.Manager,[clientgone] state (Stream) function for status of client.

> rooms.go

* Message { serId RoomIdTExt}
* Listener { RoomId  Chan}
* Manager { 2chan   broadcast}
* NewRoomManager  {  }