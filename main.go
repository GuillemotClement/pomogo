package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

func getUserInput() (int, error) {
	var durationStr string 
	fmt.Println("Saisir une durée de travail en minute: ")
	fmt.Scan(&durationStr)

	// conversion de la saisie 
	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		return 0, fmt.Errorf("la durée %d est invalide : %w", durationStr, err)
	}

	// check if valus is in the range 
	if duration < 1 || duration > 60 {
		return 0, fmt.Errorf("La durée en minute doit être entre 1 et 60: %d", duration)
	}

	return duration, nil
}

func startMinuterie(duration int){
	durationInSecond := duration * 60  

	for i := durationInSecond; i > 0; i-- {
		fmt.Printf("Temps restant: %d secondes\n", i)
		time.Sleep(1 * time.Second)
	}

	// send notification when period is end
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		fmt.Printf("Impossible de jour le super son de libération %v\n", err)
	}
	fmt.Println("Période de travail terminé")
}

func main(){
	fmt.Println("Welcome to Pomogo, let's go start a new work session !")
	duration, err := getUserInput()
	if err != nil {
		fmt.Printf("erreur de saisie: %v", err)
		return
	}

	startMinuterie(duration)
}



