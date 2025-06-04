package main

import (
	"net/mail"

	"github.com/alash3al/go-smtpsrv"
	"github.com/markterence/smtp2http/plugins"
)

func extractEmails(addr []*mail.Address, _ ...error) []string {
	ret := []string{}

	for _, e := range addr {
		ret = append(ret, e.Address)
	}

	return ret
}

func transformStdAddressToEmailAddress(addr []*mail.Address) []*EmailAddress {
	ret := []*EmailAddress{}

	for _, e := range addr {
		ret = append(ret, &EmailAddress{
			Address: e.Address,
			Name:    e.Name,
		})
	}

	return ret
}

func getHTMLBody(app_context *AppContext, msg *smtpsrv.Email) string {
	flags := app_context.flags

	// When base64 and compressed HTML is requested
	if flags.Base64HTML && flags.CompressBase64 {
		return plugins.CompressHTMLBody(string(msg.HTMLBody))
	} else if flags.Base64HTML && !flags.CompressBase64 {
		// When only base64 is requested
		return plugins.StringToBase64(string(msg.HTMLBody))
	} else {
		// just ignore that case when base64 is false and compression is true.
		// Return the HTML body as is.
		return string(msg.HTMLBody)
	}
}

// func smtpsrvMesssage2EmailMessage(msg *smtpsrv.Context)
