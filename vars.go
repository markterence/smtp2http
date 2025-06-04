package main

import (
	"flag"
	"log"
)

var (
	flagServerName     = flag.String("name", "smtp2http", "the server name")
	flagListenAddr     = flag.String("listen", ":smtp", "the smtp address to listen on")
	flagWebhook        = flag.String("webhook", "http://localhost:8080/my/webhook", "the webhook to send the data to")
	flagMaxMessageSize = flag.Int64("msglimit", 1024*1024*2, "maximum incoming message size")
	flagReadTimeout    = flag.Int("timeout.read", 5, "the read timeout in seconds")
	flagWriteTimeout   = flag.Int("timeout.write", 5, "the write timeout in seconds")
	flagAuthUSER       = flag.String("user", "", "user for smtp client")
	flagAuthPASS       = flag.String("pass", "", "pass for smtp client")
	flagDomain         = flag.String("domain", "", "domain for recieving mails")
	flagBase64HTML     = flag.Bool("base64html", false, "encode HTML body in base64 (default: false) for webhook")
	flagCompressBase64 = flag.Bool("compressbase64", false, "compress the base64 HTML body (default: false) for webhook")
)

func init() {
	flag.Parse()
	// Log the flags
	log.Printf("Server Name: %s", *flagServerName)
	log.Printf("Listen Address: %s", *flagListenAddr)
	log.Printf("Webhook URL: %s", *flagWebhook)
	log.Printf("Max Message Size: %d bytes", *flagMaxMessageSize)
	log.Printf("Read Timeout: %d seconds", *flagReadTimeout)
	log.Printf("Write Timeout: %d seconds", *flagWriteTimeout)
	log.Printf("Domain for Receiving Mails: %s", *flagDomain)
	log.Printf("Base64 HTML Encoding: %t", *flagBase64HTML)
	log.Printf("Compress Base64 HTML: %t", *flagCompressBase64)
	log.Println("Flags initialized successfully")
}
