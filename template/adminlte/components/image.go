package components

import (
	"html/template"
	"github.com/chenhg5/go-admin/template/types"
)

type ImgAttribute struct {
	Name   string
	Witdh  string
	Height string
	Src    string
}

func (*AdminlteStruct) Image() types.ImgAttribute {
	return &ImgAttribute{
		"image",
		"50",
		"50",
		"",
	}
}

func (compo *ImgAttribute) SetWidth(value string) types.ImgAttribute {
	(*compo).Witdh = value
	return compo
}

func (compo *ImgAttribute) SetHeight(value string) types.ImgAttribute {
	(*compo).Height = value
	return compo
}

func (compo *ImgAttribute) SetSrc(value string) types.ImgAttribute {
	(*compo).Src = value
	return compo
}

func (compo *ImgAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "image")
}