package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)
func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Server started123")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type netbankingRetailHdfcHandler struct {
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	testChromeDp()
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello world!"}`))
}

func testChromeDp2() {

	ctx2, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	fmt.Print("in testchromedp2 ========")
	timeoutContext, cancel := context.WithTimeout(ctx2, 10*time.Second)
	defer cancel()

	var example string
	err := chromedp.Run(timeoutContext,
		chromedp.Navigate(`https://pkg.go.dev/time`),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`body > footer`),
		// find and click "Example" link
		chromedp.Click(`#example-After`, chromedp.NodeVisible),
		// retrieve the text of the textarea
		chromedp.Value(`#example-After textarea`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}

func testChromeDp() {
	// create context

	ctx2, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	fmt.Print("reached here========")
	timeoutContext, cancel := context.WithTimeout(ctx2, 30*time.Second)
	defer cancel()

	// run task list
	h := netbankingRetailHdfcHandler{}
	authUrl := "https://api.razorpay.com/v1/payments/I5Bu6qpXWlRA4y/authenticate"
	var shot1, shot2, shot3, shot4, shot5, shot6, shot7, shot8 []byte
	err := chromedp.Run(timeoutContext, h.automateNetbankingRetailHdfcOTPSubmission(authUrl, &shot1, &shot2, &shot3, &shot4, &shot5, &shot6, &shot7, &shot8))
	if err != nil {
		fmt.Print("chromdp error", err)
	}

	if err1 := ioutil.WriteFile("shot1.png", shot1, 0o644); err1 != nil {
		fmt.Print("write error")
	}

	if err2 := ioutil.WriteFile("shot2.png", shot2, 0o644); err2 != nil {
		fmt.Print("write error")
	}

	if err3 := ioutil.WriteFile("shot3.png", shot3, 0o644); err3 != nil {
		fmt.Print("write error")
	}

	if err4 := ioutil.WriteFile("shot4.png", shot4, 0o644); err4 != nil {
		fmt.Print("write error")
	}

	if err5 := ioutil.WriteFile("shot5.png", shot5, 0o644); err5 != nil {
		fmt.Print("write error")
	}

	if err6 := ioutil.WriteFile("shot6.png", shot6, 0o644); err6 != nil {
		fmt.Print("write error")
	}

	if err7 := ioutil.WriteFile("shot7.png", shot7, 0o644); err7 != nil {
		fmt.Print("write error")
	}
	if err8 := ioutil.WriteFile("shot8.png", shot8, 0o644); err8 != nil {
		fmt.Print("write error")
	}
}

func (h netbankingRetailHdfcHandler) Do(ctx context.Context) error {

	otpstring := "123456"

	return chromedp.SendKeys(`//input[@name="fldOtpToken"]`, otpstring).Do(ctx)
}

func (h *netbankingRetailHdfcHandler) automateNetbankingRetailHdfcOTPSubmission(authUrl string, shot1 *[]byte, shot2 *[]byte, shot3 *[]byte, shot4 *[]byte, shot5 *[]byte, shot6 *[]byte, shot7 *[]byte, shot8 *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(authUrl),
		chromedp.FullScreenshot(shot1, 90),
		chromedp.WaitVisible(`//input[@name="fldLoginUserId"]`),
		chromedp.FullScreenshot(shot2, 90),
		chromedp.SendKeys(`//input[@name="fldLoginUserId"]`, "yourusername"),
		chromedp.Click(`//a[@onclick="return fLogon();"]`),
		chromedp.WaitVisible(`//input[@name="fldPassword"]`),
		chromedp.FullScreenshot(shot3, 90),
		chromedp.SendKeys(`//input[@name="fldPassword"]`, "yourpassword"),
		chromedp.FullScreenshot(shot4, 90),
		chromedp.Click(`//input[@name="chkrsastu"]`),
		chromedp.FullScreenshot(shot5, 90),
		chromedp.Click(`//a[@onclick="return fLogon();"]`),
		chromedp.WaitVisible(`//img[@alt="Continue"]`),
		chromedp.FullScreenshot(shot6, 90),
		chromedp.Click(`//img[@alt="Continue"]`),
		chromedp.WaitVisible(`//input[@name="fldOtpToken"]`),
		chromedp.FullScreenshot(shot7, 90),
		chromedp.Sleep(30 * time.Second),
		netbankingRetailHdfcHandler{},
		chromedp.FullScreenshot(shot8, 90),
		chromedp.Click(`//img[@alt="Submit"]`),
	}
}