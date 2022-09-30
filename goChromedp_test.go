package goChromedp

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGoChromedp_Screenshot(t *testing.T) {
	goChr := &GoChromedp{}
	goChr.Url = "https://www.oschina.net"
	goChr.Quality = 100
	ctx, cancel := goChr.Ctx()
	resp := goChr.Screenshot(ctx, cancel)
	res, _ := json.Marshal(resp)
	log.Println(string(res))
}

func TestGoChromedp_PrintPdf(t *testing.T) {
	goChr := &GoChromedp{}
	goChr.Url = "https://www.oschina.net"
	ctx, cancel := goChr.Ctx()
	resp := goChr.PrintPdf(ctx, cancel)
	res, _ := json.Marshal(resp)
	log.Println(string(res))
}
