package weather_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jreisinger/gokatas/weather"
)

func TestParseResponse_CorrectlyParsesJSONData(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}
	want := weather.Conditions{
		Summary:     "Clouds",
		Temperature: 291.39,
	}
	got, err := weather.ParseResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseResponse_ReturnsErrorGivenEmptyData(t *testing.T) {
	t.Parallel()
	_, err := weather.ParseResponse([]byte{})
	if err == nil {
		t.Fatal("want error parsing empty response, got nil")
	}
}

func TestParseResponse_ReturnsErrorGivenInvalidData(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/weather_invalid.json")
	if err != nil {
		t.Fatal(err)
	}
	_, err = weather.ParseResponse(data)
	if err == nil {
		t.Fatal("want error parsin invalid response, got nil")
	}
}

func TestFormatURL_ReturnsCorrectURL(t *testing.T) {
	t.Parallel()
	key := "abc123"
	c := weather.NewClient(key)
	location := "Paris,FR"
	want := "https://api.openweathermap.org/data/2.5/weather?q=Paris,FR&APPID=abc123"
	got := c.FormatURL(location)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetWeather_ReturnsCorrectWeatherConditions(t *testing.T) {
	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "testdata/weather.json")
		}))
	defer ts.Close()
	c := weather.NewClient("dummyKey")
	c.HTTPClient = ts.Client()
	c.BaseURL = ts.URL
	location := "Paris,FR"
	got, err := c.GetWeather(location)
	if err != nil {
		t.Fatal(err)
	}
	want := weather.Conditions{
		Summary:     "Clouds",
		Temperature: 291.39,
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCelsius_CorrectlyConvertsKelvinToCelsius(t *testing.T) {
	t.Parallel()
	input := weather.Temperature(274.15)
	want := 1.0
	got := input.Celsius()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
