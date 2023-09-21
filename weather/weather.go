// Package weather shows how to build an API client.
//
// # Request data
//
// We may need to send request data as a JSON-encoded body, rather than embedded
// in the URL. In this case it’s probably a good idea to define some APIRequest
// adapter struct, as we did with OWMResponse.  To make sure we’re marshalling
// the request data correctly, the httptest handler can check it by
// unmarshalling it and comparing it with the original value. For APIs with
// multiple endpoints with different types of requests and responses, we may
// need multiple adapter structs.
//
// # CRUD methods
//
// One common pattern for APIs that manage some kind of external resource is the
// set of methods known as CRUD: Create, Read, Update, and Delete. It makes
// sense to map each of these to a corresponding Go method on the client object.
// There’s usually some unique ID associated with the resource, so that we can
// specify the one we want. Usually the API assigns an ID, so Create should
// return it, while Read, Update, and Delete should take it as a parameter. You
// can often write a single test for all the CRUD methods at once. It should
// create a resource, read it to check it was created properly, update it, and
// read it again “to make sure the update worked. The test can then delete the
// resource, and make sure that a final read fails because it no longer exists.
//
// Level: advanced
// Topics: api, net/http/httptest, tpg-tools
package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Conditions struct {
	Summary string
	Temperature
}

type owmResponse struct {
	Weather []struct {
		Main string
	}
	Main struct {
		Temp float64
	}
}

func ParseResponse(data []byte) (Conditions, error) {
	var owmResponse owmResponse
	if err := json.Unmarshal(data, &owmResponse); err != nil {
		return Conditions{}, fmt.Errorf(
			"invalid API response %s: %w", data, err)
	}
	if len(owmResponse.Weather) < 1 {
		return Conditions{}, fmt.Errorf("invalid API response %s: want at least one weather element", data)
	}
	conditions := Conditions{
		Summary:     owmResponse.Weather[0].Main,
		Temperature: Temperature(owmResponse.Main.Temp),
	}
	return conditions, nil
}

type Client struct {
	apiKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(key string) Client {
	return Client{
		apiKey:     key,
		BaseURL:    "https://api.openweathermap.org",
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c Client) FormatURL(location string) string {
	URL := fmt.Sprintf("%s/data/2.5/weather?q=%s&APPID=%s", c.BaseURL, location, c.apiKey)
	return URL
}

func (c Client) GetWeather(location string) (Conditions, error) {
	URL := c.FormatURL(location)
	resp, err := c.HTTPClient.Get(URL)
	if err != nil {
		return Conditions{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Conditions{}, fmt.Errorf(
			"unexpected response status: %q", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Conditions{}, err
	}
	return ParseResponse(data)
}

func Get(location, key string) (Conditions, error) {
	c := NewClient(key)
	conditions, err := c.GetWeather(location)
	if err != nil {
		return Conditions{}, err
	}
	return conditions, nil
}

const usage = `Usage: weather LOCATION

Example: weather London,UK`

func Main() int {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		return 1
	}
	location := os.Args[1]

	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		log.Fatal("Please set the environment variable OPENWEATHERMAP_API_KEY.")
	}

	conditions, err := Get(location, key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Printf("%s %.1f°C\n", conditions.Summary, conditions.Temperature.Celsius())
	return 0
}

type Temperature float64

func (t Temperature) Celsius() float64 {
	return float64(t) - 273.15
}
