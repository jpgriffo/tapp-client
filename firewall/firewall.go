package firewall

import "runtime"
import "fmt"

import "github.com/jpgriffo/tapp-client/firewall/data"
import "github.com/jpgriffo/tapp-client/firewall/iptables"

func Apply(firewall data.Policy) () {
	fmt.Println(runtime.GOOS)
	iptables.Apply(firewall)
}

func Drop() () {
	iptables.Drop()
}