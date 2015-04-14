package utils

import "strings"

type Tag struct {
	Name   string
	Option string
}

func (t Tag) Ignore() bool {
	return t.Option == Ignore
}

func (t Tag) OmitEmpty() bool {
	return t.Option == OmitEmpty
}

func GetTag(tags string) *Tag {

	fields := strings.Split(tags, ",")

	tg := Tag{}

	if len(fields) > 1 {
		s := []string{",", fields[1]}
		fields[1] = strings.Join(s, "")
		tg.Option = fields[1]
	}

	tg.Name = fields[0]

	return &tg
}
