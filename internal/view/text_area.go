package view

type TextArea struct {
}

func NewTextArea() *TextArea {
	return &TextArea{}
}

func (v *TextArea) Render(a *ViewAtts, g GridWriter) error {
	return nil
}
