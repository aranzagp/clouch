package mapify

import (
	"fmt"
	"strings"
)

const (
	ignore    = "-"
	omitempty = ",omitempty"
)

type tag struct {
	name   string
	option string
}

func (t tag) Ignore() bool {
	return t.option == ignore
}

func (t tag) OmitEmpty() bool {
	return t.option == omitempty
}

func getTag(tags string) *tag {
	//fmt.Println(tags)
	fields := strings.Split(tags, ",")
	fmt.Println(fields)
	fmt.Println(len(fields))

	tg := tag{}

	if len(fields) > 1 {
		tg.option = fields[1]
	}

	tg.name = fields[0]

	return &tg
}
