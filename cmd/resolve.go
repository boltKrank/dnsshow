/*
Cheers SIMON ANDERSON <simon@boltkrank.com>
*/
package cmd

import (
	"fmt"
	"net"
	"strings"

	"github.com/miekg/dns"
	"github.com/spf13/cobra"
)

// resolveCmd represents the resolve command
var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Simplified output from a DNS resolution",
	Long:  `For educational purposes and more easy to understand explanation of DNS resolution`,
	Run: func(cmd *cobra.Command, args []string) {

		name := args[0]

		if !strings.HasSuffix(name, ".") {
			name = name + "."
		}
		fmt.Println("Result:", resolve(name))
	},
}

func resolve(name string) net.IP {

	fmt.Println("Reolving URL: ", name)

	fmt.Println("Using a root server from https://www.iana.org/domains/root/servers  (a.root-servers.net : 	198.41.0.4, 2001:503:ba3e::2:30 )")

	nameserver := net.ParseIP("198.41.0.4")

	for {

		// We always start with the root server ("represented by the . at the end of the address, or in this case - the IP")
		reply := dnsQuery(name, nameserver)

		// If we get a properly formed reply we don't need to go further (this reply might be from cache, or somewhere else - we don't care with the client)
		if ip := getAnswer(reply); ip != nil {
			// Best case: we get an answer to our query and we're done
			return ip
		} else if nsIP := getGlue(reply); nsIP != nil {
			// Second best: we get a "glue record" with the *IP address* of another nameserver to query
			nameserver = nsIP
		} else if domain := getNS(reply); domain != "" {
			// Third best: we get the *domain name* of another nameserver to query, which we can look up the IP for
			nameserver = resolve(domain)
		} else {
			// If there's no A record we just panic, this is not a very good
			// resolver :)
			panic("something went wrong")
		}
	}

}

func getAnswer(reply *dns.Msg) net.IP {
	for _, record := range reply.Answer {
		if record.Header().Rrtype == dns.TypeA {
			fmt.Println("  ", record)
			return record.(*dns.A).A
		}
	}

	// ANSWER PRINT:
	// fmt.Println("\n\n--- ANSWER   ---")
	// answer := message.Answer
	// fmt.Printf("\nanswer: %v\n", answer)

	// fmt.Printf("\n0                                            15")
	// fmt.Printf("\n+---------------------------------------------+")
	// fmt.Printf("\n|                    NAME                     |")
	// fmt.Printf("\n+---------------------------------------------|")
	// fmt.Printf("\n|                    TYPE                     |")
	// fmt.Printf("\n+---------------------------------------------|")
	// fmt.Printf("\n|                    CLASS                    |")
	// fmt.Printf("\n+---------------------------------------------|")
	// fmt.Printf("\n|                     TTL                     |")
	// fmt.Printf("\n+---------------------------------------------|")
	// fmt.Printf("\n|                   RDLENGTH                  |")
	// fmt.Printf("\n+---------------------------------------------|")
	// fmt.Printf("\n|                     RDATA                   |")
	// fmt.Printf("\n+---------------------------------------------+")
	// fmt.Printf("\n\n\n")

	return nil
}

func getGlue(reply *dns.Msg) net.IP {
	for _, record := range reply.Extra {
		if record.Header().Rrtype == dns.TypeA {
			fmt.Println("  ", record)
			return record.(*dns.A).A
		}
	}
	return nil
}

func getNS(reply *dns.Msg) string {
	for _, record := range reply.Ns {
		if record.Header().Rrtype == dns.TypeNS {
			fmt.Println("  ", record)
			return record.(*dns.NS).Ns
		}
	}
	return ""
}

func dnsQuery(name string, server net.IP) *dns.Msg {
	fmt.Printf("Checking %s %s\n", server.String(), name)
	msg := new(dns.Msg)
	fmt.Printf("\n msg is: %v", msg)
	msg.SetQuestion(name, dns.TypeA)
	fmt.Printf("\n msg is now: %v", msg)

	populateDiagram(msg)

	c := new(dns.Client)
	reply, _, _ := c.Exchange(msg, server.String()+":53")
	return reply
}

func boolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func populateDiagram(message *dns.Msg) {

	// fmt.Println("Message being genrated as: ")
	// fmt.Println("Populated, this is")
	// fmt.Println("            --------HEADER--------               ")
	// fmt.Println("                              1  1  1  1  1  1")
	// fmt.Println("0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5")
	// fmt.Println("+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+")
	// fmt.Println("|                      ID                       |")
	// fmt.Println("+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+")
	// fmt.Println("|QR|   OpCode  |AA|TC|RD|RA| Z|AD|CD|   RCODE   |")
	// fmt.Println("+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+")
	// fmt.Println("|                QDCOUNT/ZOCOUNT                |")
	// fmt.Println("+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+")
	// fmt.Println("|                ANCOUNT/PRCOUNT                |")
	// fmt.Println("+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+")
	// fmt.Println("|                NSCOUNT/UPCOUNT                |")
	// fmt.Println("+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+")
	// fmt.Println("|                    ARCOUNT                    |")
	// fmt.Println("+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+")

	fmt.Printf("\n msmessage g is now: %v", message)

	// Compressed ?
	compressed := boolToString(message.Compress)
	fmt.Printf("\ncompressed: %s\n", compressed)

	// ID
	id := message.MsgHdr.Id
	fmt.Printf("\nID: %d\n", id)

	// QR
	response := boolToString(message.MsgHdr.Response)
	fmt.Printf("\nQR: %s\n", response)

	// OpCode
	opCode := message.MsgHdr.Opcode
	fmt.Printf("\nopCode: %d\n", opCode)

	// AA
	authoritative := boolToString(message.MsgHdr.Authoritative)
	fmt.Printf("\nAuthoritative: %s\n", authoritative)

	// TC
	truncated := boolToString(message.MsgHdr.Truncated)
	fmt.Printf("\ntruncated: %s\n", truncated)

	// RD
	recursionDesired := boolToString(message.MsgHdr.RecursionDesired)
	fmt.Printf("\nrecursionDesired: %s\n", recursionDesired)

	// RA
	recursionAvailable := boolToString(message.MsgHdr.RecursionAvailable)
	fmt.Printf("\nrecursionAvailable: %s\n", recursionAvailable)

	// Z
	zero := boolToString(message.MsgHdr.Zero)
	fmt.Printf("\nzero: %s", zero)

	// AD
	authenticatedData := boolToString(message.MsgHdr.AuthenticatedData)
	fmt.Printf("\nauthenticatedData: %s\n", authenticatedData)

	// CD
	checkingDisabled := boolToString(message.MsgHdr.CheckingDisabled)
	fmt.Printf("\ncheckingDisabled: %s", checkingDisabled)

	// RCODE
	rcode := message.MsgHdr.Rcode
	fmt.Printf("\nrcode: %d\n", rcode)

	// QDCOUNT/ZOCOUNT

	// NSCOUNT/UPCOUNT
	// ARCOUNT

	fmt.Printf("\nQuestion:\n")
	fmt.Print(message.Question)
	fmt.Printf("\nAnswer: %v\n", message.Answer)
	fmt.Printf("Ns: %v\n", message.Ns)
	fmt.Printf("\nExtra: %v", message.Extra)
	fmt.Println("\n\n ")

	// QDCOUNT
	qdcount := uint16(len(message.Question))
	fmt.Printf("\nqdcount: %d\n", qdcount)

	ancount := uint16(len(message.Answer))
	fmt.Printf("\nqdcount: %d\n", ancount)

	nscount := uint16(len(message.Ns))
	fmt.Printf("\nqdcount: %d\n", nscount)

	arcount := uint16(len(message.Extra))
	fmt.Printf("\nqdcount: %d\n", arcount)

	fmt.Println("--- MESSAGE HEADER ---")

	fmt.Println("                              1  1  1  1  1  1")
	fmt.Println("0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5")
	fmt.Println("+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+")
	fmt.Printf("|                      %d                    |\n", id)
	fmt.Println("+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+")
	// fmt.Println("|QR|   OpCode  |AA|TC|RD|RA| Z|AD|CD|   RCODE     |")
	fmt.Printf("| %s |    %d      | %s| %s| %s| %s| %s| %s| %s|     %d    |\n",
		response, opCode, authoritative, truncated, recursionDesired, recursionAvailable, zero,
		authenticatedData, checkingDisabled, rcode)
	fmt.Println("+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+")
	fmt.Printf("|                      %d                        |\n", qdcount)
	fmt.Println("+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+")
	fmt.Printf("|                      %d                        |\n", ancount)
	fmt.Println("+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+")
	fmt.Printf("|                      %d                        |\n", nscount)
	fmt.Println("+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+")
	fmt.Printf("|                      %d                        |\n", arcount)
	fmt.Println("+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+")

	fmt.Println("--- QUESTION ---")
	//message.Question

	fmt.Println("+---------------------------------------------+")
	fmt.Println("| QNAME                                       |")
	fmt.Println("+---------------------------------------------|")
	fmt.Println("| QTYPE                                       |")
	fmt.Println("+---------------------------------------------|")
	fmt.Println("| QCLASS                                      |")
	fmt.Println("+---------------------------------------------+")

	question := message.Question
	fmt.Printf("\nquestion: %v\n", &question)
	fmt.Printf("\nquestion.length: %d", len(question))

	// // fmt.Printf("\nquestio[0]: %v\n", &question)
	// // fmt.Printf("\nquestio[0]: %v\n", &question)
	// // fmt.Printf("\nquestio[0]: %v\n", &question)

	// fmt.Println("---   NS     ---")
	// fmt.Printf("\n\n\n")
	// ns := message.Ns
	// fmt.Printf("\nns: %v\n", ns)
	// fmt.Println("---  EXTRA   ---")
	// fmt.Printf("\n\n\n")
	// extra := message.Extra
	// fmt.Printf("\nextra: %v\n", extra)
	// fmt.Printf("\n\n\n")

}

func init() {
	rootCmd.AddCommand(resolveCmd)

	// TODO
	// Add parameter flags here
	// i.e. change root server, or filter the outputs.

}
