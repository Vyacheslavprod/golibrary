//Интерфейсы, утверждение типа
package main

import "github.com/vyacheslavprod/golibrary/gadget"

type Player interface{
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

//Утверждение типа
func TryOut(player Player)  {
	player.Play("Test sound")
	player.Stop()
	recorder := player.(gadget.TapeRecorder)
	recorder.Record()
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip up", "9 to 5"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
	
	//Утверждение типа
	TryOut(gadget.TapeRecorder{})
	
	//Утверждение типа
	var tapRecord gadget.TapeRecorder = player.(gadget.TapeRecorder)
	tapRecord.Record()
	
}