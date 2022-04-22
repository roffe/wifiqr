package cmd

import (
	"fmt"
	"strings"

	"github.com/skip2/go-qrcode"
)

func genWiFiQR(authenticationType, ssid, passwd string, hidden bool, size int) ([]byte, error) {
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
