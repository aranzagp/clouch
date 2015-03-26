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
