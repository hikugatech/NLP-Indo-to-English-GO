package translator

type WordsTL struct {
	id    []string
	en    []string
	grams []string
	tags  []string
}

func New(id []string, en []string, grams []string, tags []string) WordsTL {
	tl := WordsTL{id: id, en: en, grams: grams, tags: tags}
	return tl
}
