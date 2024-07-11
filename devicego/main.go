// Интерфейсы, утверждение типа
package main

import (
	"fmt"

	"github.com/vyacheslavprod/golibrary/gadget"
)

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

//Функция для TapePlayer и TapeRecorder с проверкой безопастности вызова метода
func TryOut(player Player)  {
	player.Play("Test sound")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip up", "9 to 5"}
	var player Player = gadget.TapePlayer{}
	//Проверка, можно ли вызвать метод
	record, ok := player.(gadget.TapeRecorder)
	if ok {
		record.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
	fmt.Println()
	//Утверждение типа
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
	//Утверждение типа
	var tapRecord gadget.TapeRecorder = player.(gadget.TapeRecorder)
	fmt.Println()
	tapRecord.Record()
	
}