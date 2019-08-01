package command

import (
	"text/template"
	"bytes"
	"fmt"
	"wikux-api/common"
)

func CommandSendEmailPembayaran(){
	type EmailData struct {
		Name string
		Email string
		Price string
		RequestDate string
	}

	var emailTemplateStr = `
	<!DOCTYPE html>
	<html>
	
	<head> </head>
	
	<body>
		<p style="margin: 0in 0in 8pt; line-height: 107%; font-size: 11pt; font-family: Calibri, sans-serif;"><strong><em><span style="font-size: 14.0pt; color: black;">Mohon segera selesaikan pembayaran anda sebelum Kamis, 02 Agustus 2019 23:59.</span></em></strong></p>
		<p style="line-height: 11.75pt; margin: 0in 0in 8pt; font-size: 11pt; font-family: Calibri, sans-serif;"><strong><em><span style="color: black;">&nbsp;</span></em></strong></p>
		<p style="margin: 0in 0in 8pt; line-height: 11.75pt; font-size: 12pt; font-family: 'Times New Roman', serif;"><span style="font-size: 11.0pt; font-family: Calibri, sans-serif; color: black;">Rincian Pesanan&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;: Tiket Early Bird Single Package</span></p>
		<p style="margin: 0in 0in 8pt; line-height: 11.75pt; font-size: 12pt; font-family: 'Times New Roman', serif;"><span style="font-size: 11.0pt; font-family: Calibri, sans-serif; color: black;">Jumlah&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;: 1</span></p>
		<p style="line-height: 11.75pt; margin: 0in 0in 8pt; font-size: 11pt; font-family: Calibri, sans-serif;"><span style="color: black;">&nbsp;</span></p>
		<p style="line-height: 11.75pt; margin: 0in 0in 8pt; font-size: 11pt; font-family: Calibri, sans-serif;"><span style="color: black;">Hi, </span><span style="font-size: 14.0pt; color: black;"> {{ .Name }} </span></p>
		<p style="line-height: 106%; margin: 0in 0in 8pt; font-size: 11pt; font-family: Calibri, sans-serif;"><span style="font-size: 12.0pt; line-height: 106%; color: black;">Silakan lakukan pembayaran ke rekening berikut:</span></p>
		<p style="line-height: 106%; margin: 0in 0in 8pt; font-size: 11pt; font-family: Calibri, sans-serif;"><span style="font-size: 12.0pt; line-height: 106%; color: black;">Nomor Rekening&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; : <strong>07-0000-680-7106 (Bank Mandiri)</strong></span></p>
		<p style="line-height: 106%; margin: 0in 0in 8pt; font-size: 11pt; font-family: Calibri, sans-serif;"><span style="font-size: 12.0pt; line-height: 106%; color: black;">Nama Pemilik &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; : <strong>Yayasan Alumni Wikusama SMK Telkom Malang.</strong></span></p>
		<p style="line-height: 11.75pt; margin: 0in 0in 8pt; font-size: 11pt; font-family: Calibri, sans-serif;"><span style="color: black;">Total yang harus dibayar&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span style="font-size: 14.0pt; color: black;">Rp {{ .Price }} </span></p>
		<p style="margin: 0in 0in 8pt 0.5in; line-height: 11.75pt; font-size: 11pt; font-family: Calibri, sans-serif;"><span style="font-size: 10.0pt; color: black;">(mohon ditransfer beserta kode uniknya, untuk memudahkan kami mengidentifikasi pembayaran)</span></p>
		<p style="margin: 0in 0in 8pt 0.5in; line-height: 11.75pt; font-size: 11pt; font-family: Calibri, sans-serif;">&nbsp;</p>
		<p style="text-align: justify; line-height: 11.75pt; margin: 0in 0in 8pt; font-size: 11pt; font-family: Calibri, sans-serif;"><span style="color: black;">Jika ada pertanyaan, silakan email ke publikasi@wikusama.org atau dapat menghubungi (via whatsapp chat) ke +62 896-1349-0734.</span></p>
	</body>
	</html>
	`
	var to = []EmailData{
		EmailData{"", "", "", "",},
	}
	
	tpl, _ := template.New("email").Parse(emailTemplateStr)

	for _, t := range to {
		var tplBuffer bytes.Buffer

		if err := tpl.Execute(&tplBuffer, t); err != nil {
			return
		}

		fmt.Println(t)

		common.CommonSendEmail(common.CommonSingleSendEmail{t.Email, 
			[]string{},
			[]string{},
			"Konfirmasi Pesanan", 
			tplBuffer.String()})
	}
}