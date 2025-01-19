package graphql_fuzzy_tester

type corpus struct {
	value        string
	numOfSuccess int
}

type corpusManager struct {
	limit int
}

func newCorpusManager(limit int) *corpusManager {
	return &corpusManager{limit: limit}
}

func (cm *corpusManager)
