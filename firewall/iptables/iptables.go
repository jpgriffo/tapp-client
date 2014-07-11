package iptables

import "fmt"
import "github.com/jpgriffo/tapp-client/firewall/data"

func Apply(firewall data.Policy) {
	fmt.Println("iptables -A INPUT -i lo -j ACCEPT")
	fmt.Println("iptables -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT")

	for _, rule := range firewall.Rules {
		fmt.Printf("iptables -A INPUT -s %s -p %s --dport %d:%d -j ACCEPT\n", rule.Cidr, rule.Protocol, rule.MinPort, rule.MaxPort)
	}
	fmt.Println("iptables -P INPUT ACCEPT")
	fmt.Println("iptables -F INPUT")
}

func Drop() {
	fmt.Println("iptables -F INPUT")
	fmt.Println("iptables -P INPUT DROP")
}
