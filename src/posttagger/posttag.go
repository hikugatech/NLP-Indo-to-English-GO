package posttagger

type Kalimat struct {
	// words []string //Word = Kata
	gram []string //grammar = jenis kata
	Tags []string //Tags = Jenis kalimat
}

func New(grams []string) Kalimat {
	var tags []string
	// memasukkan param kedalam struct
	kalimat := Kalimat{grams, tags}

	//melakukan proses posttagger untuk menghasilkan jenis kalimat
	// kalimat = kalimat.loop_posttag()
	return kalimat
}
