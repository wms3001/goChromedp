package goChromedp

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/wms3001/goCommon"
	"github.com/wms3001/goTool"
)

type GoChromedp struct {
	Url     string `json:"url"`
	Quality int    `json:"quality"`
}

func (goChromedp *GoChromedp) Ctx() (context.Context, context.CancelFunc) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	return ctx, cancel
}

func (goChromedp *GoChromedp) CloseCtx(cancel context.CancelFunc) {
	cancel()
}

func (goChromedp *GoChromedp) Screenshot(ctx context.Context, cancel context.CancelFunc) *goCommon.Resp {
	var buf []byte
	resp := &goCommon.Resp{}
	tool := &goTool.GoTool{}
	tasks := chromedp.Tasks{
		chromedp.Navigate(goChromedp.Url),
		chromedp.FullScreenshot(&buf, goChromedp.Quality),
	}
	if err := chromedp.Run(ctx, tasks); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Data = tool.ByteToBase64(buf)
		resp.Code = 1
		resp.Message = "success"
	}
	goChromedp.CloseCtx(cancel)
	return resp
}

func (goChromedp *GoChromedp) PrintPdf(ctx context.Context, cancel context.CancelFunc) *goCommon.Resp {
	var buf []byte
	resp := &goCommon.Resp{}
	tool := &goTool.GoTool{}
	tasks := chromedp.Tasks{
		chromedp.Navigate(goChromedp.Url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, _ = page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			return nil
		}),
	}
	if err := chromedp.Run(ctx, tasks); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Data = tool.ByteToBase64(buf)
		resp.Code = 1
		resp.Message = "success"
	}
	goChromedp.CloseCtx(cancel)
	return resp
}
