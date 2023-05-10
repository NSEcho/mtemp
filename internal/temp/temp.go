package temp

type Fetcher interface {
	Fetch() (string, error)
}

type Checker interface {
	Check() ([]Content, error)
}

type Reader interface {
	Read(val, mail any) error
}

type Temper interface {
	Fetcher
	Checker
}

type Content struct {
	ID      any
	Subject string
	From    string
	Body    string
}
