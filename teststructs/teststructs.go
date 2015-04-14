package teststructs

type TestA struct {
	ID    string
	Revs  []string
	Hello string
}

type TestB struct {
	ID    string
	Revs  []string
	Hello string
	Foo   *string
	Num   *int
}

type TestC struct {
	ID    string
	Revs  []string
	TestA TestA
}

type TestD struct {
	ID    string
	Revs  []string
	TestA *TestA
}

type TestE struct {
	ID    string
	Revs  []string
	TestD TestD
}

type TestF struct {
	DocID    string   `clouch:"_id"`
	Revision []string `clouch:"_revs"`
	Hello    string   `clouch:"hello"`
}

type TestG struct {
	ID    string
	Revs  []string
	Hello string  `clouch:",omitempty"`
	Num   int     `clouch:",omitempty"`
	Num2  int     `clouch:",omitempty"`
	Float float64 `clouch:",omitempty"`
	Foo   string  `clouch:",omitempty"`
}

type TestH struct {
	ID    string
	Revs  []string
	Bool  bool     `clouch:",omitempty"`
	Slice []string `clouch:",omitempty"`
}

type TestI struct {
	ID    string
	Revs  []string
	Hello string   `clouch:",omitempty"`
	Num   int      `clouch:",omitempty"`
	Float float64  `clouch:",omitempty"`
	Bool  bool     `clouch:",omitempty"`
	Slice []string `clouch:",omitempty"`
	TestB *TestB   `clouch:",omitempty"`
}

type TestJ struct {
	ID    string
	Revs  []string
	Hello string `clouch:"-"`
}
