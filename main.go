package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Vote struct {
	candidates  map[string]int
	departement string
	nbVote      int
}

func main() {
	// Ouverture du fichier
	file, err := os.Open("data.csv")

	if err != nil {
		fmt.Println("Le fichier n'a pas pu être ouvert:", err)
	}
	fileScan := bufio.NewScanner(file)

	fileScan.Split(bufio.ScanLines)

	totalOfVote := 0
	isFirstLine := true

	numberOfVoteByCandidate := make(map[string]int)
	numberVoteByCandidateByDepartement := make(map[string]int)
	rankingByDepartement := make(map[string]int)

	for fileScan.Scan() {

		// On saute la première ligne
		if isFirstLine {
			isFirstLine = false
			continue
		}

		vote := createEntry(fileScan.Text())

		for key, value := range vote.candidates {
			candidateDepartementKey := key + "_" + vote.departement
			numberVoteByCandidateByDepartement[candidateDepartementKey] += value
			numberOfVoteByCandidate[key] += value
		}
		rankingByDepartement[vote.departement] += vote.nbVote
		totalOfVote += vote.nbVote
	}

	fmt.Print("Nombre de votes : ", totalOfVote)
	fmt.Print("\n\n\n ------------------------------------------------------------------------------------------------------------------------------------------------ ")
	fmt.Print("\n Nombre de votes par candidats \n")

	// Affichage du nombres de votes
	for key, value := range numberOfVoteByCandidate {
		fmt.Print("\n Candidat: ", key, " Nombre de votes : ", value)
	}

	fmt.Print("\n\n\n ------------------------------------------------------------------------------------------------------------------------------------------------ ")
	fmt.Print("\n Nombre de votes par candidats par département \n")
	// Affichage des vote par candidat et département
	for key, value := range numberVoteByCandidateByDepartement {
		splitedKey := strings.Split(key, "_")
		candidat := splitedKey[0]
		departement := splitedKey[1]
		fmt.Print("\n Département: ", departement, " Candidat : ", candidat, " Number of vote : ", value)
	}

	fmt.Print("\n\n\n ------------------------------------------------------------------------------------------------------------------------------------------------ ")
	fmt.Print("\n Classement par cirsconcription \n")

	displayRankingByCirconscription(rankingByDepartement)
}

func displayRankingByCirconscription(rankingByDepartement map[string]int) {

	ranking := make([]string, 0, len(rankingByDepartement))

	for key := range rankingByDepartement {
		ranking = append(ranking, key)
	}

	sort.SliceStable(ranking, func(index, j int) bool {
		return rankingByDepartement[ranking[index]] > rankingByDepartement[ranking[j]]
	})

	for index := 0; index < len(ranking); index++ {
		fmt.Print("\n#", index+1, ": ", ranking[index])
	}
}

func createEntry(data string) Vote {
	// On split les données
	dataSplit := strings.Split(data, ";")

	vote := &Vote{}
	vote.candidates = getDataOfCandidate(dataSplit)
	vote.departement = dataSplit[1]
	intVar, err := strconv.Atoi(dataSplit[10])

	if err != nil {
		fmt.Print(err.Error())
	}

	vote.nbVote = intVar
	return *vote
}

func getDataOfCandidate(dataSplit []string) map[string]int {
	candidat := make(map[string]int)

	for index := 23; index < len(dataSplit); index += 7 {
		if index > len(dataSplit)-1 {
			break
		}

		name := dataSplit[index]
		voteStr := dataSplit[index+2]

		vote, _ := strconv.Atoi(voteStr)

		candidat[name] = vote
	}
	return candidat
}
