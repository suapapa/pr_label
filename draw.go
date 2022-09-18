package main

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/pkg/errors"
	"github.com/suapapa/print_address_label/draw"
	"golang.org/x/image/font"
)

const (
	// font size
	fsFromAddr = 36
	fsFromName = 48
	fsToAddr   = 100
	fsToName   = 160
)

func drawAddressFrom(addr *Addr) (image.Image, error) {
	mw := ql800MaxWidth
	var addrF font.Face
	var err error
	if addrF, err = draw.GetFont(fsFromAddr); err != nil {
		return nil, errors.Wrap(err, "fail to draw from")
	}

	var addrLines []string
	addrLines = append(addrLines, draw.FitToLines(addrF, mw, addr.Line1)...)
	addrLines = append(addrLines, draw.FitToLines(addrF, mw, addr.Line2)...)

	img, err := drawAddress(addrLines, addr.Name, addr.PhoneNumber, fsFromAddr, fsFromName, mw)
	if err != nil {
		return nil, errors.Wrap(err, "fail to draw from")
	}
	return img, nil
}

func drawAddressTo(addr *Addr) (image.Image, error) {
	w := (ql800MaxWidth * 3) / 2
	var addrF font.Face
	var err error
	if addrF, err = draw.GetFont(fsToAddr); err != nil {
		return nil, errors.Wrap(err, "fail to draw from")
	}
	addrLines := []string{addr.Line1, addr.Line2}
	for _, al := range addrLines {
		lw, _ := draw.MeasureTxt(addrF, al)
		if lw > w {
			w = lw
		}
	}

	img, err := drawAddress(addrLines, addr.Name, addr.PhoneNumber, fsToAddr, fsToName, w+5)
	if err != nil {
		return nil, errors.Wrap(err, "fail to draw from")
	}

	// TODO: rotate 90
	return img, nil
}

func drawAddress(addrLines []string, name, pn string, addrFSize, nameFSize float64, width int) (image.Image, error) {
	addrF, err := draw.GetFont(addrFSize)
	if err != nil {
		return nil, err
	}
	nameF, err := draw.GetFont(nameFSize)
	if err != nil {
		return nil, err
	}
	phoneF, err := draw.GetFont(addrFSize / 2)
	if err != nil {
		return nil, err
	}

	var y float64
	dc := gg.NewContext(width, int(addrFSize+5)*len(addrLines)+int(nameFSize+nameFSize*0.2+10))
	dc.SetColor(color.White)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.SetFontFace(addrF)
	for _, line := range addrLines {
		y += (addrFSize + 5)
		dc.DrawStringAnchored(line, 5, y, 0, 0)
	}
	y += (nameFSize + 5)
	dc.SetFontFace(nameF)
	dc.DrawStringAnchored(name, float64(width)-5, y, 1, 0)
	dc.SetFontFace(phoneF)
	dc.DrawStringAnchored(pn, 5, y, 0, 0)

	return dc.Image(), nil
}
