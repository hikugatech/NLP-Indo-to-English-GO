package naivebayes

import (
	"io/ioutil"
	"nlp/src/dll"
	"strings"
)

type Posttag struct {
	isi [][]string
}

func New(files string) Posttag {
	var dataposttag [][]string
	var data Posttag
	txt, _ := ioutil.ReadFile(files)
	k := strings.Split(string(txt), "\n")

	for y := 0; y < len(k); y++ {
		a := strings.Split(k[y], ",")
		dataposttag = append(dataposttag, a)
	}
	data = Posttag{dataposttag}
	return data
}

func (data Posttag) X(match []string) string {
	// fmt.Println(data.isi[0])
	var temp []string
	var countjum []int
	for _, e := range data.isi {
		akhir := len(e) - 1
		if len(temp) == 0 {
			temp = append(temp, e[akhir])
			countjum = append(countjum, 1)
		} else if dll.Match_array(temp, e[akhir]) == false {
			temp = append(temp, e[akhir])
			countjum = append(countjum, 1)
		} else if dll.Match_array(temp, e[akhir]) == true {
			index := dll.Search_array(temp, e[akhir])
			countjum[index]++
		}
		// fmt.Println(e)
	}

	pXk := dll.Start_array_float32(len(temp))
	// fmt.Println(temp)
	for y, e := range match {
		p := data.p(e, y)
		for f, j := range p {
			if pXk[f] == -1 {
				pXk[f] = j
			} else {
				pXk[f] *= j
			}
		}
	}

	// fmt.Println(pXk)

	var hasil []float32
	for f, e := range pXk {
		o := e * (float32(countjum[f]) / float32(len(data.isi)))
		hasil = append(hasil, o)
	}

	var tmp float32
	tmp = -1
	for _, e := range hasil {
		if tmp == -1 {
			tmp = e
			continue
		}
		if tmp < e {
			tmp = e
		}
	}

	ruwet := dll.Search_array_float32(hasil, tmp)
	return temp[ruwet]
}

func (data Posttag) p(Xk string, Ci int) []float32 {
	var temp []string
	var countjum []int
	for _, e := range data.isi {
		akhir := len(e) - 1
		if len(temp) == 0 {
			temp = append(temp, e[akhir])
			countjum = append(countjum, 1)
		} else if dll.Match_array(temp, e[akhir]) == false {
			temp = append(temp, e[akhir])
			countjum = append(countjum, 1)
		} else if dll.Match_array(temp, e[akhir]) == true {
			index := dll.Search_array(temp, e[akhir])
			countjum[index]++
		}
	}

	countXk := dll.Start_array_int(len(countjum))
	for _, e := range data.isi {
		akhir := len(e) - 1
		// fmt.Println(e[2])
		if dll.Match_array(temp, e[akhir]) == true && e[Ci] == Xk {
			index := dll.Search_array(temp, e[akhir])
			countXk[index]++
		}
	}

	// fmt.Println(countXk)
	// fmt.Println(countjum)
	var jum []float32
	for y, e := range countjum {
		jum = append(jum, float32(countXk[y])/float32(e))
	}
	return jum
}
