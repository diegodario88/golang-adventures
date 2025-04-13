package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	chromeHost := os.Getenv("CHROME_HOST")
	if chromeHost == "" {
		chromeHost = "localhost"
	}

	devToolsURL := fmt.Sprintf("http://%s:9222", chromeHost)
	fmt.Printf("Connecting to Chrome at %s\n", devToolsURL)

	allocCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), devToolsURL)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var buf []byte

	URL := "https://www.bcb.gov.br/estatisticas/detalhamentoGrafico/graficoshome/selic"
	if err := chromedp.Run(ctx, printToPDF(URL, &buf)); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./assets/comprovante.pdf", buf, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("wrote comprovante.pdf")
}

func printToPDF(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(".grafico-destaque", chromedp.ByQuery),
		chromedp.Sleep(3 * time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {

			pdfParams := page.PrintToPDF().
				WithPrintBackground(true).
				WithDisplayHeaderFooter(true).
				WithPreferCSSPageSize(false)

			buf, _, err := pdfParams.Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
