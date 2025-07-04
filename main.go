package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/miekg/dns"
)

const (
	domainName   = "ddd.uz."               // Domain to intercept
	serverIP     = "192.168.100.191"       // Change this to your real LAN IP
	staticFolder = "./dist"                // Path to static website
)

func startHTTP() {
	fs := http.FileServer(http.Dir(staticFolder))
	http.Handle("/", fs)
	log.Printf("[HTTP] Serving files from %s at http://%s", staticFolder, serverIP)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func forwardDNS(w dns.ResponseWriter, r *dns.Msg) {
	c := &dns.Client{}
	resp, _, err := c.Exchange(r, "8.8.8.8:53")
	if err != nil {
		log.Printf("[DNS] Forward error: %v", err)
		dns.HandleFailed(w, r)
		return
	}
	w.WriteMsg(resp)
}

func startDNS() {
	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		m := new(dns.Msg)
		m.SetReply(r)
		for _, q := range r.Question {
			if strings.EqualFold(q.Name, domainName) {
				rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, serverIP))
				if err == nil {
					m.Answer = append(m.Answer, rr)
					log.Printf("[DNS] %s -> %s", q.Name, serverIP)
				}
				w.WriteMsg(m)
				return
			}
		}
		forwardDNS(w, r)
	})
	go dns.ListenAndServe(":53", "udp", nil)
	go dns.ListenAndServe(":53", "tcp", nil)
	log.Printf("[DNS] Intercepting %s on port 53", domainName)
}

func main() {
	log.Println("[AIO] Starting up...")
	go startDNS()
	go startHTTP()

	// Wait for Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	log.Println("[AIO] Stopping")
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
