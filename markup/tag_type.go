package markup

type TagType int

const (
	TagTypeOpen = TagType(iota)
	TagTypeClose
	TagTypeSelfClosing
	TagTypeCloseAll
)
