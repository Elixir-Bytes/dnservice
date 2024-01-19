package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lixiangzhong/dnsutil"
	"github.com/miekg/dns"
)

var host = flag.String("h", "", "Host to query")

func main() {
	flag.Parse()

	if *host == "" {
		log.Fatal("Can't run without -h flag with host")
	}

	var dig dnsutil.Dig
	dig.SetDNS("1.1.1.1")
	a, err := dig.A(*host)
	fmt.Println(a, err)
}

// type A struct {
// 	Hdr RR_Header
// 	A   net.IP `dns:"a"`
// }

// type RR_Header struct {
// 	Name     string `dns:"cdomain-name"`
// 	Rrtype   uint16
// 	Class    uint16
// 	Ttl      uint32
// 	Rdlength uint16 // Length of data after header.
// }

type DNSRecord struct {
	Error  error
	IP     string
	Header DNSHeader
}

type DNSHeader struct {
	Name     string `json:"name"`
	RrType   uint16 `json:"rr_type"`
	Class    uint16 `json:"class"`
	Ttl      uint32 `json:"ttl"`
	RdLength uint16 `json:"rd_length"`
}

func convertHeader(header dns.RR_Header) DNSHeader {
	return DNSHeader{
		header.Name,
		header.Rrtype,
		header.Class,
		header.Ttl,
		header.Rdlength,
	}
}

func foo(dig dnsutil.Dig, host string) []DNSRecord {
	resChan := make(chan []DNSRecord)

	go func(results chan []DNSRecord) {
		res, err := dig.A(host)
		if err != nil {
			record := DNSRecord{Error: err}
			resChan <- []DNSRecord{record}
			return
		}
		records := []DNSRecord{}
		for _, record := range res {
			ip := string(record.A)
			header := convertHeader(record.Hdr)
			rec := DNSRecord{IP: ip, Header: header}
			records = append(records, rec)
		}
		resChan <- records
	}(resChan)

	go func(results chan []DNSRecord) {
		res, err := dig.SOA(host)
		if err != nil {
			record := DNSRecord{Error: err}
			resChan <- []DNSRecord{record}
			return
		}
		records := []DNSRecord{}
		for _, record := range res {
			ip := string(record.SOA)
			header := convertHeader(record.Hdr)
			rec := DNSRecord{IP: ip, Header: header}
			records = append(records, rec)
		}
		resChan <- records
	}(resChan)

	go func(results chan []DNSRecord) {
		res, err := dig.NS(host)
		if err != nil {
			record := DNSRecord{Error: err}
			resChan <- []DNSRecord{record}
			return
		}
		records := []DNSRecord{}
		for _, record := range res {
			ip := string(record.NS)
			header := convertHeader(record.Hdr)
			rec := DNSRecord{IP: ip, Header: header}
			records = append(records, rec)
		}
		resChan <- records
	}(resChan)

	go func(results chan []DNSRecord) {
		res, err := dig.CNAME(host)
		if err != nil {
			record := DNSRecord{Error: err}
			resChan <- []DNSRecord{record}
			return
		}
		records := []DNSRecord{}
		for _, record := range res {
			ip := string(record.CNAME)
			header := convertHeader(record.Hdr)
			rec := DNSRecord{IP: ip, Header: header}
			records = append(records, rec)
		}
		resChan <- records
	}(resChan)

	go func(results chan []DNSRecord) {
		res, err := dig.A(host)
		if err != nil {
			record := DNSRecord{Error: err}
			resChan <- []DNSRecord{record}
			return
		}
		records := []DNSRecord{}
		for _, record := range res {
			ip := string(record.A)
			header := convertHeader(record.Hdr)
			rec := DNSRecord{IP: ip, Header: header}
			records = append(records, rec)
		}
		resChan <- records
	}(resChan)

	// a, err := dig.PTR(host)
	// a, err := dig.TXT(host)

	go func(results chan []DNSRecord) {
		res, err := dig.A(host)
		if err != nil {
			record := DNSRecord{Error: err}
			resChan <- []DNSRecord{record}
			return
		}
		records := []DNSRecord{}
		for _, record := range res {
			ip := string(record.A)
			header := convertHeader(record.Hdr)
			rec := DNSRecord{IP: ip, Header: header}
			records = append(records, rec)
		}
		resChan <- records
	}(resChan)

	go func(results chan []DNSRecord) {
		res, err := dig.A(host)
		if err != nil {
			record := DNSRecord{Error: err}
			resChan <- []DNSRecord{record}
			return
		}
		records := []DNSRecord{}
		for _, record := range res {
			ip := string(record.A)
			header := convertHeader(record.Hdr)
			rec := DNSRecord{IP: ip, Header: header}
			records = append(records, rec)
		}
		resChan <- records
	}(resChan)
	// a, err := dig.AAAA(host)
	// a, err := dig.MX(host)
	// a, err := dig.SRV(host)
	// a, err := dig.CAA(host)
	// a, err := dig.SPF(host)

	allRecords := []DNSRecord{}
	for i := 0; i < 11; i++ {
		records := <-resChan
		for _, record := range records {
			allRecords = append(allRecords, record)
		}
	}

	close(resChan)

	return allRecords
}
