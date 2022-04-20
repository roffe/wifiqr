package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	qrcode "github.com/skip2/go-qrcode"
	flag "github.com/spf13/pflag"
)

var (
	authenticationType, ssid, passwd string
	hidden                           bool
	size                             int
)

func init() {
	flag.StringVar(&authenticationType, "type", "WPA", "Authentication type")
	flag.StringVar(&ssid, "ssid", "test", "SSID")
	flag.StringVar(&passwd, "password", "123456", "Password")
	flag.BoolVar(&hidden, "hidden", false, "Hidden SSID")
	flag.IntVar(&size, "size", 1000, "QR-code size")
	flag.Parse()
}

func main() {
	b, err := genWiFiQR()
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("qr.png", b, 0755); err != nil {
		log.Fatal(err)
	}
}

func genWiFiQR() ([]byte, error) {
	var str strings.Builder

	str.WriteString(fmt.Sprintf("WIFI:T:%s;S:%s;P:%s;", strings.ToUpper(authenticationType), ssid, passwd))
	if hidden {
		str.WriteString("H:true;")
	}
	str.WriteString(";")

	q, err := qrcode.New(str.String(), qrcode.High)
	if err != nil {
		return nil, err
	}

	return q.PNG(size)
}
