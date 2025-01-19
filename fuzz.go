package graphql_fuzzy_tester

// Opts is a struct that holds the options for the fuzzing
type Opts struct {
}

type QueryFunc func(string) error

func Fuzz(opts Opts, fn QueryFunc) {

}
