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
	fsToName   = 120
)

func drawAddressFrom(addr *Addr) (image.Image, error) {
	mw := ql800MaxPix
	var addrF font.Face
	var err error
	if addrF, err = draw.GetFont(fsFromAddr); err != nil {
		return nil, errors.Wrap(err, "fail to draw from")
	}

	var addrLines []string
	addrLines = append(addrLines, draw.FitToLines(addrF, mw, addr.Line1)...)
	addrLines = append(addrLines, draw.FitToLines(addrF, mw, addr.Line2)...)

	img, err := drawAddress(addrLines, addr.Name, addr.PhoneNumber, fsFromAddr, fsFromName, mw, -1)
	if err != nil {
		return nil, errors.Wrap(err, "fail to draw from")
	}
	return img, nil
}

func drawAddressTo(addr *Addr) (image.Image, error) {
	mw := (ql800MaxPix * 3) / 2
	var addrF font.Face
	var err error
	if addrF, err = draw.GetFont(fsToAddr); err != nil {
		return nil, errors.Wrap(err, "fail to draw from")
	}

	var addrLines []string
	addrLines = append(addrLines, draw.FitToLines(addrF, mw, addr.Line1)...)
	addrLines = append(addrLines, draw.FitToLines(addrF, mw, addr.Line2)...)

	img, err := drawAddress(addrLines, addr.Name, addr.PhoneNumber, fsToAddr, fsToName, mw, ql800MaxPix)
	if err != nil {
		return nil, errors.Wrap(err, "fail to draw from")
	}

	// TODO: rotate 90
	return img, nil
}

func drawAddress(addrLines []string, name, pn string, addrFSize, nameFSize float64, width int, height int) (image.Image, error) {
	addrF, err := draw.GetFont(addrFSize)
	if err != nil {
		return nil, err
	}
	nameF, err := draw.GetFont(nameFSize)
	if err != nil {
		return nil, err
	}
	phoneFSize := addrFSize / 2
	phoneF, err := draw.GetFont(phoneFSize)
	if err != nil {
		return nil, err
	}

	var y float64
	var varHeight bool
	if height < 0 {
		varHeight = true
	}

	if varHeight {
		height = int(addrFSize+5)*len(addrLines) + int(nameFSize+nameFSize*0.2+10)
	}
	dc := gg.NewContext(width, height)
	dc.SetColor(color.White)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.SetFontFace(addrF)
	for _, line := range addrLines {
		y += (addrFSize + 5)
		dc.DrawStringAnchored(line, 5, y, 0, 0)
	}
	if varHeight {
		y += (nameFSize + 5)
		dc.SetFontFace(nameF)
		dc.DrawStringAnchored(name, float64(width)-5, y, 1, 0)
		dc.SetFontFace(phoneF)
		dc.DrawStringAnchored(pn, 5, y, 0, 0)
	} else {
		y = float64(height - 5)
		dc.SetFontFace(phoneF)
		// dc.DrawStringAnchored(pn, 5, y, 0, -1)
		dc.DrawStringAnchored(pn, float64(width)-5, y, 1, -1)
		y -= (phoneFSize + 5)
		y = float64(height - 5)
		dc.SetFontFace(nameF)
		dc.DrawStringAnchored(name, float64(width)-5, y, 1, -1)
	}

	return dc.Image(), nil
}
