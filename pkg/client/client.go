package client

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/ryanmab/rdap-go/internal/cache"
	"github.com/ryanmab/rdap-go/internal/query"
	"github.com/ryanmab/rdap-go/internal/registry"
	"github.com/ryanmab/rdap-go/pkg/client/response/asn"
	"github.com/ryanmab/rdap-go/pkg/client/response/dns"
	"github.com/ryanmab/rdap-go/pkg/client/response/ipv4"
	"github.com/ryanmab/rdap-go/pkg/client/response/ipv6"
)

// Client is an RDAP client for performing lookups on domains, IPv4, and IPv6 addresses.
type Client struct {
	httpClient *http.Client
	cache      *cache.Cache
}

// New creates a new RDAP client instance with default settings.
func New() *Client {
	return &Client{
		httpClient: &http.Client{},
		cache:      cache.New(),
	}
}

// LookupDomain looks up a domain, using RDAP and retrieves its Domain registration data.
func (client *Client) LookupDomain(domain string) (*dns.Response, error) {
	domain = strings.ToLower(strings.TrimSpace(domain))

	if !strings.HasPrefix(domain, "http://") && !strings.HasPrefix(domain, "https://") {
		domain = "https://" + domain
	}
	url, err := url.Parse(domain)

	if err != nil {
		return nil, err
	}

	tld := url.Hostname()[strings.LastIndex(url.Hostname(), ".")+1:]

	slog.Info("Parsed domain to TLD for lookup", "domain", domain, "tld", tld)

	servers, err := registry.GetServers(query.DomainQuery, tld)

	if err != nil {
		slog.Error("failed to get RDAP servers for TLD", "tld", tld, "error", err)
		return nil, err
	}

	response, err := client.request(servers, query.DomainQuery, url.Hostname())

	if response, ok := response.(dns.Response); ok {
		return &response, err
	}

	return nil, fmt.Errorf("unexpected response type returned from RDAP server call (expected dns.Response), type was: %T", response)
}

// LookupIPv4 looks up an IPv4 address, using RDAP and retrieves its IP registration data.
func (client *Client) LookupIPv4(ip string) (*ipv4.Response, error) {
	ip = strings.TrimSpace(ip)
	servers, err := registry.GetServers(query.IPv4Query, ip)

	if err != nil {
		slog.Error("failed to get RDAP servers for IPv4 address", "ipv4", ip, "error", err)
		return nil, err
	}

	response, err := client.request(servers, query.IPv4Query, ip)

	if response, ok := response.(ipv4.Response); ok {
		return &response, err
	}

	return nil, fmt.Errorf("unexpected response type returned from RDAP server call (expected ipv4.Response), type was: %T", response)
}

// LookupIPv6 looks up an IPv6 address, using RDAP and retrieves its IP registration data.
func (client *Client) LookupIPv6(ip string) (*ipv6.Response, error) {
	ip = strings.TrimSpace(ip)
	servers, err := registry.GetServers(query.IPv6Query, ip)

	if err != nil {
		slog.Error("failed to get RDAP servers for IPv6 address", "ipv6", ip, "error", err)
		return nil, err
	}

	response, err := client.request(servers, query.IPv6Query, ip)

	if response, ok := response.(ipv6.Response); ok {
		return &response, err
	}

	return nil, fmt.Errorf("unexpected response type returned from RDAP server call (expected ipv6.Response), type was: %T", response)
}

// LookupASN looks up a given Autnum, using RDAP and retrieves its registration data.
func (client *Client) LookupASN(autnum uint32) (*asn.Response, error) {
	autnumAsString := strconv.FormatUint(uint64(autnum), 10)

	servers, err := registry.GetServers(query.AsnQuery, autnumAsString)

	if err != nil {
		slog.Error("failed to get RDAP servers for Autnum", "asn", autnum, "error", err)
		return nil, err
	}

	response, err := client.request(servers, query.AsnQuery, autnumAsString)

	if response, ok := response.(asn.Response); ok {
		return &response, err
	}

	return nil, fmt.Errorf("unexpected response type returned from RDAP server call (expected asn.Response), type was: %T", response)

}

// ClearCache empties the cache of any responses previously recorded by the Client.
func (client *Client) ClearCache() {
	client.cache.Clear()
}

// WithHTTPClient sets a custom HTTP client for the RDAP client to use for requests.
func (client *Client) WithHTTPClient(httpClient *http.Client) {
	client.httpClient = httpClient
}

// WithCache sets a custom cache for the RDAP client to use for caching responses.
func (client *Client) WithCache(cache *cache.Cache) {
	client.cache = cache
}

// Request performs an RDAP request to the provided servers for the given query type and identifier.
func (client *Client) request(servers []string, queryType query.RdapQuery, identifier string) (any, error) {
	if output := client.cache.Get(queryType, identifier); output != nil {
		slog.Info("Response cache hit. Using cached response instead of performing RDAP request", "identifier", identifier, "query", queryType)

		return *output, nil
	}

	for _, server := range servers {
		serverResponse, err := client.httpClient.Get(server + queryType.String() + "/" + identifier)

		if err != nil {
			slog.Warn("RDAP server request failed. Using another server if available.", "server", server, "error", err)
			continue
		}

		if serverResponse.StatusCode != http.StatusOK {
			slog.Warn("RDAP server request returned non-200 status code. Using another server if available.", "server", server, "code", serverResponse.StatusCode)
			continue
		}

		defer serverResponse.Body.Close()

		response, err := parseResponse(queryType, serverResponse)

		if err != nil {
			slog.Warn("Failed to parse RDAP server response. Using another server if available.", "server", server, "error", err)
			return response, err
		}

		slog.Info("RDAP server request successful", "server", server, "identifier", identifier, "query", queryType)

		client.cache.Set(queryType, identifier, response)

		return response, err
	}

	return nil, fmt.Errorf("all RDAP servers failed for query type %s and identifier %s", queryType.String(), identifier)
}

// Parse the RDAP server response based on the query type into a validated response
// struct.
func parseResponse(queryType query.RdapQuery, response *http.Response) (any, error) {
	decoder := json.NewDecoder(response.Body)
	validate := validator.New(validator.WithRequiredStructEnabled())

	switch queryType {
	case query.DomainQuery:
		var output dns.Response
		if err := decoder.Decode(&output); err != nil {
			return nil, err
		}

		if err := validate.Struct(output); err != nil {
			return nil, err
		}

		return output, nil
	case query.IPv4Query:
		var output ipv4.Response

		if err := decoder.Decode(&output); err != nil {
			return nil, err
		}

		if err := validate.Struct(output); err != nil {
			return nil, err
		}

		return output, nil
	case query.IPv6Query:
		var output ipv6.Response

		if err := decoder.Decode(&output); err != nil {
			return nil, err
		}

		if err := validate.Struct(output); err != nil {
			return nil, err
		}

		return output, nil

	case query.AsnQuery:
		var output asn.Response

		if err := decoder.Decode(&output); err != nil {
			return nil, err
		}

		if err := validate.Struct(output); err != nil {
			return nil, err
		}

		return output, nil
	default:
		return nil, fmt.Errorf("unsupported query type: %s", queryType.String())
	}
}
