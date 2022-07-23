/*

Exercise: Stringers
Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

*/

package implementingStringer

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	var str string

	for i, part := range ipAddr {
		if i == len(ipAddr)-1 {
			str = str + strconv.Itoa(int(part))
		} else {
			str = str + strconv.Itoa(int(part)) + "."
		}
	}

	return str
}

func Test() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
