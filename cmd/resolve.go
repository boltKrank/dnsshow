/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
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
		reply := dnsQuery(name, nameserver)

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
	fmt.Printf("dig -r @%s %s\n", server.String(), name)
	msg := new(dns.Msg)
	msg.SetQuestion(name, dns.TypeA)
	c := new(dns.Client)
	reply, _, _ := c.Exchange(msg, server.String()+":53")
	return reply
}

func init() {
	rootCmd.AddCommand(resolveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resolveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resolveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
