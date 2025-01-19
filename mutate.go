package graphql_fuzzy_tester

type mutator struct {
	keywords []string
}

func newMutator() *mutator {
	m := &mutator{
		keywords: make([]string, len(keywords)),
	}

	for _, keyword := range keywords {
		m.keywords = append(m.keywords, keyword)
	}

	return m
}

func (m *mutator) mutate(value string) string {
	return ""
}

func (m *mutator) mutateKeyword(value string) string {
	return ""
}
