package view


type Position interface {
	Pos() (int, int)
}

type FullView struct {
	status   *StatusLine
	textArea *TextArea
}

func NewFullView(name string, c Position) *FullView {
	return &FullView{
		status:   NewStatusLine(name, c),
		textArea: NewTextArea(),
	}
}

func (v *FullView) Render(a *ViewAtts, g GridWriter) error {
	v.status.Render(a, g)
	v.textArea.Render(a, g)
	return nil
}
