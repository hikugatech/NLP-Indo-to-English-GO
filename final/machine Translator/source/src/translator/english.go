package translator

import (
	"nlp/src/dll"
	"nlp/src/naivebayes"
	"strconv"
	"strings"
)

func (tl WordsTL) Translate_EN(words []string) []string {

	var temp []string

	tenses, _ := tl.cek_tenses(words)
	// fmt.Println(tenses)

	for _, e := range words {
		temp = append(temp, tl.search_translate(e))
	}

	// fmt.Println(tenses)

	co := 0
	for x, _ := range temp {
		matcher := []string{"imbuhan", "Kdepan", "partikel", "penjelas"}
		if tl.tags[x] == "predikat" {
			break
		}
		// fmt.Println(tl.tags[x])

		if dll.Match_array(matcher, tl.tags[x]) && co < 1 && tl.tags[x] == "penjelas" {
			temp[x] = tl.to_be(temp[x-1])
			co++

			if tl.tags[x+1] == "predikat" {
				var tmb string
				if tenses == "simplepresent" {
					tmb = "-s"
				} else if tenses == "simplepast" {
					tmb = "-ed"
				} else if tenses == "presentcontinous" {
					tmb = "-ing"
				}

				temp = append(temp, "0")
				copy(temp[x+3:], temp[x+2:])
				temp[x+2] = tmb
				// temp[x] = temp[x+1]
				// temp[x+1] = tmb
			}
			continue
		}

		if len(tenses) > 0 && (dll.Match_array(matcher, tl.tags[x]) || tl.tags[x+1] == "predikat") {
			var tmb string
			if tenses == "simplepresent" {
				tmbs := "-s"
				temp = append(temp, "0")
				copy(temp[x+3:], temp[x+2:])
				temp[x+2] = tmbs
			} else if tenses == "simplepast" {
				tmb = "-ed"
				temp[x] = temp[x+1]
				temp[x+1] = tmb
			} else if tenses == "presentcontinous" {
				tmb = "-ing"
				temp[x] = temp[x+1]
				temp[x+1] = tmb
			}

		}
	}

	return temp

}

func (tl WordsTL) cek_tenses(words []string) (string, []string) {

	matcher := []string{"imbuhan", "Kdepan", "partikel", "penjelas"}
	nn := dll.Start_array_string(2)
	t := dll.Start_array_string(2)
	for x, e := range tl.tags {
		if e == "predikat" {
			break
		}
		if dll.Match_array(matcher, e) {
			if tl.grams[x] == "KD" {
				nn[1] = "KD"
				t[0] = words[x]
			} else if tl.grams[x] == "KT" {
				nn[0] = "KT"
				t[1] = words[x]
			}
		}
	}

	var match []string
	if nn[0] == "" && nn[1] == "" {
		match = []string{strconv.Itoa(0), strconv.Itoa(0), strconv.Itoa(0)}
	} else if nn[0] != "" && nn[1] != "" {
		match = []string{nn[0], nn[1], t[0]}
	} else if nn[1] != "" {
		match = []string{strconv.Itoa(0), nn[1], t[0]}
	} else {
		match = []string{nn[0], strconv.Itoa(0), strconv.Itoa(0)}
	}

	// fmt.Println(match)

	bayes := naivebayes.New("dataset/psttlerEN.csv")
	simpul := bayes.X(match)
	return simpul, t
}

func (tl WordsTL) search_translate(kata string) string {
	index := dll.Search_array(tl.id, kata)
	// fmt.Println(tl.id)
	if index == -1 {
		return "\"" + kata + "\""
	}
	return tl.en[index]
}

func (tl WordsTL) to_be(word string) string {
	am := []string{"i"}
	are := []string{"you", "they", "we"}

	word = strings.ToLower(word)

	if dll.Match_array(am, word) {
		return "am"
	} else if dll.Match_array(are, word) {
		return "are"
	} else {
		return "is"
	}
}
