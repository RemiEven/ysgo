package markup

type tagType int8

const (
	tagTypeOpen = tagType(iota)
	tagTypeClose
	tagTypeSelfClosing
	tagTypeCloseAll
)
