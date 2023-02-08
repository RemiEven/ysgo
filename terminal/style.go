package terminal

import (
	"sort"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/remieven/ysgo/markup"
)

func toStyledText(line *markup.ParseResult) string {
	if len(line.Attributes) == 0 {
		return line.Text
	}

	type insertPoint struct {
		position int
		start    bool
		color    string
		flag     string
	}

	points := make([]*insertPoint, 0, 2*len(line.Attributes))
	for _, attribute := range line.Attributes {
		var p1, p2 *insertPoint
		switch attribute.Name {
		case "color":
			color := attribute.Properties["color"].StringValue
			p1 = &insertPoint{
				position: attribute.Position,
				start:    true,
				color:    color,
			}
			p2 = &insertPoint{
				position: attribute.Position + attribute.Length,
				color:    color,
			}
		case "i", "u":
			p1 = &insertPoint{
				position: attribute.Position,
				start:    true,
				flag:     attribute.Name,
			}
			p2 = &insertPoint{
				position: attribute.Position + attribute.Length,
				flag:     attribute.Name,
			}
		default:
			continue
		}
		points = append(points, p1, p2)
	}

	slices.SortStableFunc(points, func(a, b *insertPoint) bool {
		return a.position < b.position
	})

	builder := strings.Builder{}
	lastPosition := 0
	flags := map[string]struct{}{}
	allFlags := func() string {
		r := make([]string, 0, len(flags))
		for f := range flags {
			r = append(r, f)
		}
		sort.Strings(r)
		return strings.Join(r, "")
	}

	for _, p := range points {
		if p.position > lastPosition {
			builder.WriteString(line.Text[lastPosition:p.position])
			lastPosition = p.position
		}
		switch {
		case p.color != "" && p.start:
			builder.WriteString("[" + p.color + "]")
		case p.color != "" && !p.start:
			builder.WriteString("[-]")
		case p.flag != "" && p.start:
			flags[p.flag] = struct{}{}
			builder.WriteString("[::" + allFlags() + "]")
		case p.flag != "" && !p.start:
			delete(flags, p.flag)
			if len(flags) == 0 {
				builder.WriteString("[::-]")
			} else {
				builder.WriteString("[::" + allFlags() + "]")
			}
		}
	}
	if lastPosition < len(line.Text) {
		builder.WriteString(line.Text[lastPosition:])
	}

	return builder.String()
}
