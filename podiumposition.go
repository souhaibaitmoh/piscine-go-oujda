package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i <= len(podium)-1; i++ {
		for j := i + 1; j <= len(podium)-1; j++ {
			podium[i][0], podium[j][0] = podium[j][0], podium[i][0]
		}
	}
	return podium
}
