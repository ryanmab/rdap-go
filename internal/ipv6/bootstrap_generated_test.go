package ipv6

import (
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratedBootstrapHasAllServersEndingWithTailingSlash(t *testing.T) {
	for _, servers := range Bootstrap {
		assert.NotEmpty(t, servers, "Generated bootstrap has an empty server list for one of the IPs")

		for _, server := range servers {
			assert.Equal(t, "/", string(server[len(server)-1]), "Generated bootstrap server %q does not end with a trailing slash", server)
		}
	}
}

func TestAccessOrderAndBootstrapLineUp(t *testing.T) {
	t.Run("No non-extistant accesses", func(t *testing.T) {
		for ip := range Bootstrap {
			assert.NotEmpty(t, Bootstrap[ip], "Generated bootstrap has an orphaned entry for IP: %q", ip)
		}
	})

	t.Run("No superfolous bootstrap entries", func(t *testing.T) {
		// Because of the check above, we know that every IP in Bootstrap has at least one entry, and
		// therefore we can do a simple length check here
		assert.Equal(t, len(BootstrapAccessOrder), len(Bootstrap), "Generated bootstrap has %d entries, but access order has %d entries", len(Bootstrap), len(BootstrapAccessOrder))
	})

}

func TestGeneratedBootstrapIpV6sAreInDescendingOrderOfSubnextMask(t *testing.T) {
	var previousIP string

	r := regexp.MustCompile("([a-f0-9]{1,4}):([a-f0-9]{0,4}):([a-f0-9]{0,4})[:]{0,1}([a-f0-9]{0,4})/([0-9]+)")

	for _, ip := range BootstrapAccessOrder {
		if previousIP == "" {
			// First IP, nothing to compare to
			previousIP = ip
			continue
		}

		partsPrevious := r.FindStringSubmatch(previousIP)
		partsCurrent := r.FindStringSubmatch(ip)

		subnetPrevious := partsPrevious[5]
		subnetCurrent := partsCurrent[5]

		assert.GreaterOrEqual(t, subnetPrevious, subnetCurrent, "Generated bootstrap IPv6s are not in descending order. Subnet in %q is before %q", previousIP, ip)

		previousIP = ip
	}
}

func TestGeneratedBootstrapIpV6sAreInDescendingOrderOfHextets(t *testing.T) {
	var previousIP string

	r := regexp.MustCompile("([a-f0-9]{1,4}):([a-f0-9]{0,4}):([a-f0-9]{0,4})[:]{0,1}([a-f0-9]{0,4})/([0-9]+)")

	for _, ip := range BootstrapAccessOrder {
		if previousIP == "" {
			// First IP, nothing to compare to
			previousIP = ip
			continue
		}

		partsPrevious := r.FindStringSubmatch(previousIP)
		partsCurrent := r.FindStringSubmatch(ip)

		previousHextets := partsPrevious[1:5]
		currentHextets := partsCurrent[1:5]

		subnetPrevious := partsPrevious[5]
		subnetCurrent := partsCurrent[5]

		if subnetPrevious > subnetCurrent {
			// Different subnet, no need to compare hextets
			previousIP = ip
			continue
		}

		for i := 0; i < max(len(previousHextets), len(currentHextets)); i++ {
			if previousHextets[i] == "" && currentHextets[i] == "" {
				// Both hextets are empty, continue
				continue
			}

			if previousHextets[i] == "" && currentHextets[i] != "" {
				// Previous is empty, current is not, so current is bigger, we are good
				break
			}

			currentHextet, err := strconv.ParseInt(currentHextets[i], 16, 64)
			assert.NoError(t, err, "Error parsing hextet %q in IP %q", currentHextets[i], ip)

			previousHextet, err := strconv.ParseInt(previousHextets[i], 16, 64)
			assert.NoError(t, err, "Error parsing hextet %q in IP %q", previousHextets[i], previousIP)

			assert.True(t, previousHextet >= currentHextet, "Generated bootstrap IPv6s are not in descending order. Hextet %q (%d) in IP %q is less than hextet %q (%d) in IP %q", previousHextets[i], previousHextet, previousIP, currentHextets[i], currentHextet, ip)
		}

		previousIP = ip
	}
}
