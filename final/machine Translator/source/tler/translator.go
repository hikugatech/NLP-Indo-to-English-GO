package tler

import (
	"io/ioutil"
	"nlp/dataset"
	"nlp/src/dll"
	"nlp/src/posttagger"
	"nlp/src/translator"
	"regexp"
	"strconv"
	"strings"
)

func English(text string) ([]string, []string, []string, []string) {
	words := tokenizer_EN(text)
	_, _, id, en, kataen, gramen, _ := datafiles()
	var grams []string

	for _, e := range words {
		// fmt.Println(e)
		if dll.Match_array(kataen, e) {
			index := dll.Search_array_all(kataen, e)
			// fmt.Println(index)
			var in []string
			for _, y := range index {
				if !dll.Match_array(in, gramen[y]) {
					in = append(in, gramen[y])
				}
			}
			grams = append(grams, strings.Join(in, ","))
		} else {
			grams = append(grams, strconv.Itoa(0))
		}
	}

	// fmt.Println(words)
	// fmt.Println(grams)

	tags := posttagger.New(grams)
	tags = tags.English(words)
	tler := translator.New(id, en, grams, tags.Tags)
	tler_ID := tler.Translate_ID(words)
	return words, tler_ID, grams, tags.Tags
}

func Indonesia(text string) ([]string, []string, []string, []string) {
	words := tokenizer(strings.ToLower(text))
	kataid, gramid, en, id, _, _, _ := datafiles()
	var grams []string
	for _, e := range words {
		if dll.Match_array(kataid, e) {
			index := dll.Search_array_all(kataid, e)
			var in []string
			for _, y := range index {
				in = append(in, gramid[y])
			}
			grams = append(grams, strings.Join(in, ","))
		} else {
			grams = append(grams, strconv.Itoa(0))
		}
	}

	tags := posttagger.New(grams)
	tags = tags.Indonesian(words)
	tler := translator.New(id, en, grams, tags.Tags)
	tler_EN := tler.Translate_EN(words)
	// fmt.Println(words)
	// fmt.Println(tler_EN)
	// // fmt.Println(tler_ID)
	// fmt.Println(grams)
	// fmt.Println(tags.Tags)

	return words, tler_EN, grams, tags.Tags
}

// id, en := openfile("dataset/ind-eng.csv", "", -1)
// kataid, gramid := openfile("dataset/lexup.csv", "\t", -1)
// kataen, dasaren, gramen := openfiles("dataset/lexEN.csv")

// ens := strings.Join(en, "|")
// ids := strings.Join(id, "|")
// kataids := strings.Join(kataid, "|")
// gramids := strings.Join(gramid, "|")
// kataens := strings.Join(kataen, "|")
// dasarens := strings.Join(dasaren, "|")
// gramens := strings.Join(gramen, "|")
// writes := []byte("package dataset \nvar Datakataid string = \"" + kataids + "\"\nvar Datagramid string = \"" + gramids + "\"\nvar Datakataen string = \"" + kataens + "\"\nvar Datadasaren string = \"" + dasarens + "\"\nvar Datagramen string = \"" + gramens + "\"\nvar Dataens string = \"" + ens + "\"\nvar DataID string = \"" + ids + "\"")
// ioutil.WriteFile("dataset/dataset.go", writes, 0644)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func tokenizer(text string) (words []string) {
	split := []string{"di", "me", "mem"}
	for _, e := range split {
		regex, _ := regexp.MatchString(" "+e+".", text)
		if regex {
			text = strings.Replace(text, " "+e+" ", " "+e+"= ", -1)
			text = strings.Replace(text, " "+e, " "+e+"- ", -1)
			text = strings.Replace(text, " "+e+"- = ", " "+e+" ", -1)
		}
	}
	words = strings.Fields(text)
	return
}

func tokenizer_EN(text string) (words []string) {
	// split := []string{"es", "s", "ing", "ed"}
	// for _, e := range split {
	// 	regex, _ := regexp.MatchString("!?"+e+"", text)
	// 	if regex {
	// 		text = strings.Replace(text, ""+e+"", " -"+e+" ", -1)
	// 	}
	// }
	words = strings.Fields(text)
	return
}

func datafiles() ([]string, []string, []string, []string, []string, []string, []string) {
	gramid := strings.Split(string(dataset.Datagramid), "|")
	kataid := strings.Split(string(dataset.Datakataid), "|")
	en := strings.Split(string(dataset.Dataens), "|")
	id := strings.Split(string(dataset.DataID), "|")
	gramen := strings.Split(string(dataset.Datagramen), "|")
	dasaren := strings.Split(string(dataset.Datadasaren), "|")
	kataen := strings.Split(string(dataset.Datakataen), "|")

	return kataid, gramid, en, id, kataen, gramen, dasaren
}

func openfile(files string, ciri string, jalur int) (id []string, en []string) {
	// var id []string
	// var en []string
	txt, _ := ioutil.ReadFile(files)
	k := strings.Split(string(txt), "\n")
	for _, y := range k {
		var word []string
		if ciri != "" {
			word = strings.Split(y, "\t")
			// fmt.Println(word)
		} else {
			word = strings.Split(y, ",")
		}

		var ss string
		if jalur > -1 {
			// fmt.Println(word[len(word)-1])
			ss = word[len(word)-1]
		} else {
			ss = word[1]
		}
		// fmt.Println(word)
		eng := strings.Split(ss, "|")

		for _, x := range eng {
			en = append(en, x)
			id = append(id, strings.ToLower(word[0]))
		}
	}

	return id, en
}
func openfiles(files string) (id []string, ens []string, en []string) {
	// var id []string
	// var en []string
	txt, _ := ioutil.ReadFile(files)
	k := strings.Split(string(txt), "\n")
	for _, y := range k {
		var word []string

		word = strings.Split(y, "\t")
		// fmt.Println(word)

		var ss string

		ss = word[2]

		// fmt.Println(word)
		eng := strings.Split(ss, "|")

		for _, x := range eng {
			en = append(en, x)
			id = append(id, strings.ToLower(word[0]))
			ens = append(ens, strings.ToLower(word[1]))
			// fmt.Println(word)
		}
	}

	return
}
