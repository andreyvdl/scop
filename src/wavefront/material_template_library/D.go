package materialtemplatelibrary

type DStyle int

type dOption func(d *D)

const (
	DEFAULT_STYLE DStyle = iota
	HALO_STYLE
)

type D struct {
	Factor float32
	Style  DStyle
}

func NewD(value float32, opts ...dOption) *D {
	d := &D{value, DEFAULT_STYLE}

	for _, o := range opts {
		o(d)
	}

	return d
}

func WithStyle(style DStyle) dOption {
	return func(d *D) {
		d.Style = style
	}
}
