package main

type sumPair struct {
	valA int
	valB int
}

func (s *sumPair) sum() int {
	return s.valA + s.valB
}

func main() {
	inFile, err := ioutil.ReadFile("../input")
	//numMap := make(map[int][]*sumPair)

	//if err != nil {
	//	log.Fatal(err)
	//}

	//scanner := bufio.NewScanner(inFile)

	//for scanner.Scan() {
	//	s := scanner.Text()
	//	val, _ := strconv.Atoi(s)
	//	key, _ := strconv.Atoi(s[len(s)-1:])

	//numMap[10-key] = append(numMap[10-key], val)

	//}

	//if err := scanner.Err(); err != nil {
	//	log.Fatal(err)
	//}

	//for key, sumList := range numMap {
	//if key >= 5 {
	//for _, valI := range sumList {
	//for _, valJ := range numMap[10-key] {
	//if valI+valJ == 2020 {
	//	fmt.Println(valI * valJ)
	//}
	//			}
	//		}
	//	}
	//}
}
