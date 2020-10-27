package test

import (
	"design-pattern/template"
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	smsOTP := &template.Sms{}
	o := template.Otp{
		IOtp: smsOTP,
	}
	o.GenAndSendOTP(4)

	fmt.Println()
	emailOTP := &template.Email{}

	o = template.Otp{
		IOtp: emailOTP,
	}
	o.GenAndSendOTP(4)
}
