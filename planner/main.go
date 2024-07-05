package main

import (
	"github.com/vyacheslavprod/golibrary/dates"
	"fmt"
)

func main() {
	days := 3
	fmt.Println("Your appoitment is in", days, "days")
	fmt.Println("with a follow-up in", days+dates.DaysInWeek, "days")
}
