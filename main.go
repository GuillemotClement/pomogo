package main

import (
	"flag"
	"fmt"
	"strconv"

	//	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

// func getUserInput() (int, error) {
//	var durationStr string
//	fmt.Println("Saisir une durée de travail en minute: ")
//	fmt.Scan(&durationStr)

// conversion de la saisie
//	duration, err := strconv.Atoi(durationStr)
//	if err != nil {
//		return 0, fmt.Errorf("la durée %d est invalide : %w", durationStr, err)
//	}

// check if valus is in the range
//	if duration < 1 || duration > 60 {
//		return 0, fmt.Errorf("La durée en minute doit être entre 1 et 60: %d", duration)
//	}

//	return duration, nil
//}

func startMinuterie(duration int) {
	fmt.Println("Début de la session de travail. Have Fun !")

	durationinsecond := duration * 60

	for i := durationinsecond; i > 0; i-- {
		fmt.Printf("temps restant: %d secondes\n", i)
		time.Sleep(1 * time.Second)
	}

	// send notification when period is end
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		fmt.Printf("impossible de jour le super son de libération %v\n", err)
	}
	fmt.Println("Période de travail terminé")
}

func main() {
	// déclation des flags
	// durationPtr := flag.String("duration", "", "Temps pour la session de travail")
	// parse les arguments lance en ligne de commande au démarrage de l'outil
	flag.Parse()

	duration, err := strconv.Atoi(flag.Arg(0)) // on recupere le premier argument sans le flag
	if err != nil {
		fmt.Printf("saisie incorrect %q: %s\n", flag.Arg(0), err)
		return
	}
	// durationStr := *durationPtr

	//duration, err := strconv.Atoi(durationStr)
	//if err != nil {
	//fmt.Printf("saisie incorrect %q: %s", durationStr, err)
	//return
	//}
	startMinuterie(duration)
}
