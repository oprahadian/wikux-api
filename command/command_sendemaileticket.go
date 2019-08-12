package command

import (
	"text/template"
	"bytes"
	"fmt"
	"wikux-api/common"
)

func CommandSendEmailEticket(){
	type EmailData struct {
		Name string
		Email string
		Eticket string
	}

	var emailTemplateStr = `
	<!DOCTYPE html>
<html>

<head></head>

<body>
    <p>Hi {{ .Name }},</p>
    <p>Terima kasih telah melakukan pemesanan tiket Early Bird untuk acara Wikusama Festival Connect 2019.
		<br />Pembayaran tiket kamu telah terkonfirmasi dan kami terima. Jangan lupa datang tepat waktu saat registrasi dibuka pada pukul 9.00 WIB dan sesi utama akan berlangsung pada pukul 10.00 WIB.</p>
		<p><a href="https://wikufest.org/wikufest_connect_2019_ticket/ticket.html?d={{ .Eticket }}">E - Ticket</a></p>
    <p>Jika kamu punya pertanyaan seputar acara ini, silahkan kirim email ke <a href="mailto:publikasi@wikusama.org">publikasi@wikusama.org</a> atau whatsapp/telegram ke <strong>+62 896-1349-0734</strong></p>
    <table>
        <tbody>
            <tr>
                <td class="m_9167785011762792042inner-container" style="padding: 13px 40px 26px 40px;" align="center" valign="top">
                    <table style="max-width: 600px; width: 100%;" border="0" width="600" cellspacing="0" cellpadding="0" align="center">
                        <tbody>
                            <tr>
                                <td style="border: solid 1px #d7d7d7;">
                                    <table border="0" cellspacing="0" cellpadding="0">
                                        <tbody>
                                            <tr>
                                                <td id="m_9167785011762792042map-image-mobile28a328670-0720-449a-b6c4-9815d9118adb" class="m_9167785011762792042map-image-mobile m_9167785011762792042mktoText" style="display: none;">
                                                    <a href="https://go.google-mkto.com/dc/Mrv_hxtRECjDkd0Ykq7qcDW5aTM9-Bkx2wSLTSBBXLfas87fxVtes_C_JKGdjQpKGMTt3EHNGcRalEnt9aPf4uRt8LfjVQlk_WXXo_ee5QvFML-Q91RWSFRZupmYDulBU2atzRrRAU2rAP3vlANAESlrxLAD3-EC_SbAOmyApGDg3wUeYUtH_j2Man1b77nK3vGCWg2cD3ursjkkII5r5sfs52R7A0wMj3TEQckvo2S6aSLa-iDSUJS5Q0DpuD4XNh75U4_4wHW0B1PXzFsNHw==/n0TI0n101tAC0CP0tB3b2KS" target="blank" rel="noopener" data-saferedirecturl="https://www.google.com/url?q=https://go.google-mkto.com/dc/Mrv_hxtRECjDkd0Ykq7qcDW5aTM9-Bkx2wSLTSBBXLfas87fxVtes_C_JKGdjQpKGMTt3EHNGcRalEnt9aPf4uRt8LfjVQlk_WXXo_ee5QvFML-Q91RWSFRZupmYDulBU2atzRrRAU2rAP3vlANAESlrxLAD3-EC_SbAOmyApGDg3wUeYUtH_j2Man1b77nK3vGCWg2cD3ursjkkII5r5sfs52R7A0wMj3TEQckvo2S6aSLa-iDSUJS5Q0DpuD4XNh75U4_4wHW0B1PXzFsNHw%3D%3D/n0TI0n101tAC0CP0tB3b2KS&amp;source=gmail&amp;ust=1565319046824000&amp;usg=AFQjCNF3aJpD53VjDKfUEsnEZagpf54rsg"> <img class="m_9167785011762792042no-arrow CToWUd" style="font-family: 'Google Sans','Roboto',Helvetica,Arial,sans-serif; color: #1a73e8;" src="https://ci5.googleusercontent.com/proxy/fr0VHi9j2nOV-H2BQkyUGjdKrR6VWGFZm5KqgSQ7yAGarEmmguvC4DE-bA6Hm3q9E8vmGukS5dIVlQVn5CiCiMEkWqJfVKQkVijlTqxXkTXyo7p-kYLOESDwqUjPp02ahxZN=s0-d-e1-ft#https://lp.google-mkto.com/rs/248-TPC-286/images/Ballroom+Kempinski+-+Mobile.png" alt="Map Image" /> </a>
                                                </td>
                                            </tr>
                                            <tr>
                                                <td id="m_9167785011762792042map-image-desktop28a328670-0720-449a-b6c4-9815d9118adb" class="m_9167785011762792042hom m_9167785011762792042mktoText">
                                                    <a href="https://goo.gl/maps/ukFi56xxwzfPADCz7" target="_blank" rel="noopener" data-saferedirecturl="https://goo.gl/maps/ukFi56xxwzfPADCz7"> <img class="m_9167785011762792042no-arrow CToWUd" style="width: 518px; font-family: 'Google Sans','Roboto',Helvetica,Arial,sans-serif; color: #1a73e8; display: block; margin: 0; border: none;" src="https://wikufest.org/assets/images/usmar-ismail-hall-maps.png" alt="Map Image" width="518" /> </a>
                                                </td>
                                            </tr>
                                            <tr>
                                                <td class="m_9167785011762792042location-details" style="padding: 24px 18px;">
                                                    <table border="0" cellspacing="0" cellpadding="0">
                                                        <tbody>
                                                            <tr>
                                                                <td id="m_9167785011762792042map-date-time28a328670-0720-449a-b6c4-9815d9118adb" class="m_9167785011762792042width100 m_9167785011762792042td-paddit m_9167785011762792042mktoText" style="width: 236px; vertical-align: top;" width="236">
                                                                    <table border="0" cellspacing="0" cellpadding="0">
                                                                        <tbody>
                                                                            <tr>
                                                                                <td style="padding: 0; vertical-align: top; max-width: 24px;" width="24"><img class="m_9167785011762792042no-arrow m_9167785011762792042no-squeeze m_9167785011762792042twentyfour CToWUd" src="https://ci4.googleusercontent.com/proxy/isTaZEz6WYvLerCUTAp6bpKMOpNhq6wkw1NEucFl-p0JtKZ3zsZi48mouJXo_Fm3B5Ia_PKi-tMIkkEXz8IYex7OnqE9F5B0p9baHGacYfc4hm0htbwJtvEOHqY=s0-d-e1-ft#https://lp.google-mkto.com/rs/248-TPC-286/images/calendar-icon-3-gray.png" alt="" width="24" height="24" /></td>
                                                                                <td style="padding: 0; width: 7px; min-width: 7px; max-width: 7px;" width="7">&nbsp;</td>
                                                                                <td class="m_9167785011762792042pt-bit" style="padding-top: 4px; word-break: break-word;"><span class="m_9167785011762792042date" style="color: #757575; font-family: 'Google Sans','Roboto',Helvetica,Arial,sans-serif; font-size: 13px; font-weight: 500; line-height: 20px;">Saturday, 31 August, 2019</span>
                                                                                    <br /><span class="m_9167785011762792042time" style="color: #757575; font-family: 'Google Sans','Roboto',Helvetica,Arial,sans-serif; font-size: 13px; font-weight: 500; line-height: 20px;">9:00 AM&ndash;4:00 PM </span></td>
                                                                            </tr>
                                                                            <tr>
                                                                                <td style="line-height: 10px; font-size: 10px;" colspan="3">&nbsp;</td>
                                                                            </tr>
                                                                        </tbody>
                                                                    </table>
                                                                    <table border="0" cellspacing="0" cellpadding="0">
                                                                        <tbody>
                                                                            <tr>
                                                                                <td style="padding: 0; vertical-align: top; max-width: 24px;" width="24">&nbsp;</td>
                                                                                <td style="padding: 0; width: 7px; min-width: 7px; max-width: 7px;" width="7">&nbsp;</td>
                                                                                <td class="m_9167785011762792042pt-bit" style="padding-top: 4px; word-break: break-word;">
                                                                                    <table border="0" cellspacing="0" cellpadding="0">
                                                                                        <tbody>
                                                                                            <tr>
                                                                                                <td><img class="m_9167785011762792042no-arrow m_9167785011762792042no-squeeze m_9167785011762792042twentyfour CToWUd" src="https://ci4.googleusercontent.com/proxy/UQMg__cQppUy9jjHS8Ylm8ozvGGdnlXASGeMBWJLune9RDJrFyZ8QpsSI0tEj221P8uZbhwryX_6u1pE6byIIOioIiXGJjP0XzOLTR6Hi4U=s0-d-e1-ft#https://lp.google-mkto.com/rs/248-TPC-286/images/cs3-plus.png_cQppUy9jjHS8Ylm8ozvGGdnlXASGeMBWJLune9RDJrFyZ8QpsSI0tEj221P8uZbhwryX_6u1pE6byIIOioIiXGJjP0XzOLTR6Hi4U=s0-d-e1-ft#https://lp.google-mkto.com/rs/248-TPC-286/images/cs3-plus.png_cQppUy9jjHS8Ylm8ozvGGdnlXASGeMBWJLune9RDJrFyZ8QpsSI0tEj221P8uZbhwryX_6u1pE6byIIOioIiXGJjP0XzOLTR6Hi4U=s0-d-e1-ft#https://lp.google-mkto.com/rs/248-TPC-286/images/cs3-plus.png_cQppUy9jjHS8Ylm8ozvGGdnlXASGeMBWJLune9RDJrFyZ8QpsSI0tEj221P8uZbhwryX_6u1pE6byIIOioIiXGJjP0XzOLTR6Hi4U=s0-d-e1-ft#https://lp.google-mkto.com/rs/248-TPC-286/images/cs3-plus.png" alt="" width="20" height="20" /></td>
                                                                                                <td style="padding: 0; width: 7px; min-width: 7px; max-width: 7px;" width="7">&nbsp;</td>
                                                                                                <td><a style="color: #1a73e8; text-decoration: none; font-family: 'Google Sans','Roboto',Helvetica,Arial,sans-serif; font-size: 13px; font-weight: 500; line-height: 20px;" href="https://go.google-mkto.com/dc/MVIYjepvLGdc_Bq05Z3HDvCFHoFM9jFSTtlurC3TYA9bw_XWXDYn90exnZS1J0krNRYPOTnQZXWrLVFI_Yf39bJlOm-Zn3VnhyznIn8COKDZcbRvJj9rldlcNqZAUNwIBao2kJUJViZeUYPr7u1LC2XjhbY_HD-uqETTj9Rd6qztCzbzkrxdd8DxQPfh6CFoaNOoN6u-ajX6L0fnTwzRNo7DvdF7yBw7QZe3JBXjnGxw58bRoHc9fDDW1O3-jjwzaoG3peCIqTHJkOUZXWTOM5Ct3RsGTzfh0eKBucgLCiHvUfiRwS6WmByepUJU9kHXyplyWLfk2r9tSJuzPR-eiIqiSnCbF1JjTwiiHqY6uZwZg8RBR_zyjUCJJVpnBqtnTJ-az6pAhCWtNA_ARYeeur67OwHZw-1jmyBs2q7vgbFqQv6umzcHPI5GV0hlenXhfBXh5DKfGBFMUqbhL9hNIwx9DxxphsbmFaQ76cS1CVPQIAyD29A5ueQuJ9xhlRWN-xaKDlrDhLRjN3NKhzPiw45bT2r0XsuFJPuFD9HDj6V2Ke-lloQvqEV0dCFIgM8Mko5ogiciqVvZuR6mdlQcTmQIfD_pR-SbDkkHMrybcdru_0mkL3j82xYK5B5aPL6RfrzfYE7-eHTOdjJLc_ekJ0VA1dFNY0k5VRxWsQ_VCa45J2gW_iNLpbazwQ2CO357Q_OHHJ5PYSBqf8UeuaGofhWY5iNXjxLumxGbmjLOM0CtbaB9F8e9zHHx47FO1I_e/n0TI0n101tAC0CP0tB3b2KS" target="_blank" rel="noopener" data-saferedirecturl="https://www.google.com/url?q=https://go.google-mkto.com/dc/MVIYjepvLGdc_Bq05Z3HDvCFHoFM9jFSTtlurC3TYA9bw_XWXDYn90exnZS1J0krNRYPOTnQZXWrLVFI_Yf39bJlOm-Zn3VnhyznIn8COKDZcbRvJj9rldlcNqZAUNwIBao2kJUJViZeUYPr7u1LC2XjhbY_HD-uqETTj9Rd6qztCzbzkrxdd8DxQPfh6CFoaNOoN6u-ajX6L0fnTwzRNo7DvdF7yBw7QZe3JBXjnGxw58bRoHc9fDDW1O3-jjwzaoG3peCIqTHJkOUZXWTOM5Ct3RsGTzfh0eKBucgLCiHvUfiRwS6WmByepUJU9kHXyplyWLfk2r9tSJuzPR-eiIqiSnCbF1JjTwiiHqY6uZwZg8RBR_zyjUCJJVpnBqtnTJ-az6pAhCWtNA_ARYeeur67OwHZw-1jmyBs2q7vgbFqQv6umzcHPI5GV0hlenXhfBXh5DKfGBFMUqbhL9hNIwx9DxxphsbmFaQ76cS1CVPQIAyD29A5ueQuJ9xhlRWN-xaKDlrDhLRjN3NKhzPiw45bT2r0XsuFJPuFD9HDj6V2Ke-lloQvqEV0dCFIgM8Mko5ogiciqVvZuR6mdlQcTmQIfD_pR-SbDkkHMrybcdru_0mkL3j82xYK5B5aPL6RfrzfYE7-eHTOdjJLc_ekJ0VA1dFNY0k5VRxWsQ_VCa45J2gW_iNLpbazwQ2CO357Q_OHHJ5PYSBqf8UeuaGofhWY5iNXjxLumxGbmjLOM0CtbaB9F8e9zHHx47FO1I_e/n0TI0n101tAC0CP0tB3b2KS&amp;source=gmail&amp;ust=1565319046825000&amp;usg=AFQjCNEGrj-bVEtEBWDTQEB2hXoFlvCxCg">Google Calendar</a></td>
                                                                                            </tr>
                                                                                            <tr>
                                                                                                <td style="line-height: 10px; font-size: 10px;" colspan="3">&nbsp;</td>
                                                                                            </tr>
                                                                                            <tr>
                                                                                                <td><img class="m_9167785011762792042no-arrow m_9167785011762792042no-squeeze m_9167785011762792042twentyfour CToWUd" src="https://ci4.googleusercontent.com/proxy/UQMg__cQppUy9jjHS8Ylm8ozvGGdnlXASGeMBWJLune9RDJrFyZ8QpsSI0tEj221P8uZbhwryX_6u1pE6byIIOioIiXGJjP0XzOLTR6Hi4U=s0-d-e1-ft#https://lp.google-mkto.com/rs/248-TPC-286/images/cs3-plus.png" alt="" width="20" height="20" /></td>
                                                                                                <td style="padding: 0; width: 7px; min-width: 7px; max-width: 7px;" width="7">&nbsp;</td>
                                                                                                <td><a style="color: #1a73e8; text-decoration: none; font-family: 'Google Sans','Roboto',Helvetica,Arial,sans-serif; font-size: 13px; font-weight: 500; line-height: 20px;" href="https://calendar.google.com/calendar/ical/publikasi%40wikusama.org/public/basic.ics" target="_blank" rel="noopener" data-saferedirecturl="https://calendar.google.com/calendar/ical/publikasi%40wikusama.org/public/basic.ics">Other Calendar</a></td>
                                                                                            </tr>
                                                                                        </tbody>
                                                                                    </table>
                                                                                </td>
                                                                            </tr>
                                                                        </tbody>
                                                                    </table>
                                                                </td>
                                                                <td class="m_9167785011762792042hom" style="width: 10px;" width="10">&nbsp;</td>
                                                                <td id="m_9167785011762792042map-location28a328670-0720-449a-b6c4-9815d9118adb" class="m_9167785011762792042width100 m_9167785011762792042mktoText" style="width: 236px; vertical-align: top;" width="236">
                                                                    <table border="0" cellspacing="0" cellpadding="0">
                                                                        <tbody>
                                                                            <tr>
                                                                                <td style="padding: 0; vertical-align: top; max-width: 24px;" width="24"><img class="m_9167785011762792042no-arrow m_9167785011762792042no-squeeze m_9167785011762792042twentyfour CToWUd" src="https://ci3.googleusercontent.com/proxy/mkdQ_4GVpNdA77U0Xpdse7KVKwGGQH-F0g_ldRDpxqVrBSUMnFgbI2mbU_h34Ni-t5f7pXrCf8INga37jN9sYMHDeLPh9kSWbO7j-PofACq7uWKa4y0Mw8M=s0-d-e1-ft#https://lp.google-mkto.com/rs/248-TPC-286/images/location-marker-2.png" alt="" width="24" height="24" /></td>
                                                                                <td style="padding: 0; width: 7px; min-width: 7px; max-width: 7px;" width="7">&nbsp;</td>
                                                                                <td class="m_9167785011762792042pt-bit" style="padding-top: 4px; word-break: break-word;"><span class="m_9167785011762792042location" style="color: #1a73e8; font-family: 'Google Sans','Roboto',Helvetica,Arial,sans-serif; font-size: 13px; font-weight: 500; line-height: 20px;"> Usmar Ismail Hall<br />(beside Plaza Festival)<br />Jl. H. R. Rasuna Said, <br />South Jakarta<br />DKI Jakarta 12940, Indonesia </span></td>
                                                                            </tr>
                                                                        </tbody>
                                                                    </table>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </td>
            </tr>
        </tbody>
    </table>
    <p>Sampai jumpa di Wikusama Festival Connect 2019</p>
</body>

</html>
	`
	var to = []EmailData{
		EmailData{"", "", "",},}
	
	tpl, _ := template.New("email").Parse(emailTemplateStr)

	for _, t := range to {
		var tplBuffer bytes.Buffer

		if err := tpl.Execute(&tplBuffer, t); err != nil {
			return
		}

		fmt.Println(t)

		common.CommonSendEmailMailgun(common.CommonSingleSendEmail{t.Email, 
			[]string{},
			[]string{"", "", "", ""},
			"Pembayaran Berhasil - Early Bird Ticket Wikusama Festival Connect 2019", 
			tplBuffer.String(),
			""})
	}
}