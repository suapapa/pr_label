package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/pkg/errors"
)

var (
	from = &Addr{
		Line1:       "경기 성남시 분당구 판교역로 235 (에이치스퀘어 엔동)",
		Line2:       "7층",
		Name:        "카카오 엔터프라이즈",
		PhoneNumber: "010-1234-5678",
	}
	to = &Addr{
		Line1:       "경기도 성남시 분당구 판교역로 166",
		Line2:       "",
		Name:        "판교 아지트",
		PhoneNumber: "010-7656-0329",
	}
)

func TestPrintOrder(t *testing.T) {
	ord := Order{
		ID:   "1234567890",
		From: from,
		To:   to,
	}
	json.NewEncoder(os.Stdout).Encode(&ord)
}

func TestDrawAddressXXX(t *testing.T) {
	img, err := drawAddressFrom(from)
	if err != nil {
		t.Error(errors.Wrap(err, "fail to draw address"))
	}
	if err := saveImg2Png(img, "addr_from.png"); err != nil {
		t.Error(err)
	}
	img, err = drawAddressTo(to)
	if err != nil {
		t.Error(errors.Wrap(err, "fail to draw address"))
	}
	if err := saveImg2Png(img, "addr_to.png"); err != nil {
		t.Error(err)
	}
}

func TestDrawAddress(t *testing.T) {
	img, err := drawAddress(
		[]string{from.Line1, from.Line2},
		from.Name,
		from.PhoneNumber,
		fsFromAddr, fsFromName,
		ql800MaxWidth,
	)
	if err != nil {
		t.Error(errors.Wrap(err, "fail to draw address"))
	}
	if err := saveImg2Png(img, "draw_from.png"); err != nil {
		t.Error(err)
	}

	img, err = drawAddress(
		[]string{to.Line1, to.Line2},
		to.Name,
		to.PhoneNumber,
		fsToAddr, fsToName,
		(ql800MaxWidth*3)/2,
	)
	if err != nil {
		t.Error(errors.Wrap(err, "fail to draw address"))
	}
	if err := saveImg2Png(img, "draw_to.png"); err != nil {
		t.Error(err)
	}
}
