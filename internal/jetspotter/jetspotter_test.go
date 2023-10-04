package jetspotter

import (
	"jetspotter/internal/aircraft"
	"jetspotter/internal/configuration"
	"reflect"
	"testing"

	"github.com/jftuga/geodist"
)

var (
	planes = []AircraftOutput{
		{
			Callsign:    "APEX11",
			Type:        aircraft.F16.Identifier,
			Description: aircraft.F16.Description,
		},
		{
			Callsign:    "APEX12",
			Type:        aircraft.F16.Identifier,
			Description: aircraft.F16.Description,
		},
		{
			Callsign:    "XSG123",
			Type:        aircraft.B77L.Identifier,
			Description: aircraft.B77L.Description,
		},
		{
			Callsign:    "GRZLY11",
			Type:        aircraft.A400.Identifier,
			Description: aircraft.A400.Description,
		},
	}

	locationMannekenPis = geodist.Coord{
		Lat: 50.844987343465924,
		Lon: 4.349981064923107,
	}

	locationElisabethPark = geodist.Coord{
		Lat: 50.86503662037458,
		Lon: 4.32399484006766,
	}

	locationChristRedeemer = geodist.Coord{
		Lat: -22.951907892908967,
		Lon: -43.21048377096087,
	}

	locationPyramidGiza = geodist.Coord{
		Lat: 29.979104641494533,
		Lon: 31.134157868680205,
	}
)

func TestFilterAircraftByTypeF16(t *testing.T) {
	expected := []AircraftOutput{
		{
			Callsign:    "APEX11",
			Type:        aircraft.F16.Identifier,
			Description: aircraft.F16.Description,
		},
		{
			Callsign:    "APEX12",
			Type:        aircraft.F16.Identifier,
			Description: aircraft.F16.Description,
		},
	}

	config := configuration.Config{
		AircraftTypes: []string{aircraft.F16.Identifier},
	}
	actual := filterAircraftByTypes(planes, config)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected '%v' to be the same as '%v'", expected, actual)
	}
}

func TestFilterAircraftByTypeALL(t *testing.T) {
	config := configuration.Config{
		AircraftTypes: []string{aircraft.ALL.Identifier},
	}
	expected := planes
	actual := filterAircraftByTypes(planes, config)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected '%v' to be the same as '%v'", expected, actual)
	}
}

func TestCalculateDistance(t *testing.T) {
	tbilisiAirportCoordinates := geodist.Coord{
		Lat: 41.4007,
		Lon: 44.5705,
	}

	kutaisiAirportCoordinates := geodist.Coord{
		Lat: 42.1033,
		Lon: 42.2830,
	}

	expected := 205
	actual := CalculateDistance(tbilisiAirportCoordinates, kutaisiAirportCoordinates)

	if expected != actual {
		t.Fatalf("expected '%v' to be the same as '%v'", expected, actual)
	}
}

func TestFilterAircraftByTypes(t *testing.T) {
	expected := []AircraftOutput{
		{
			Callsign:    "APEX11",
			Type:        aircraft.F16.Identifier,
			Description: aircraft.F16.Description,
		},
		{
			Callsign:    "APEX12",
			Type:        aircraft.F16.Identifier,
			Description: aircraft.F16.Description,
		},
		{
			Callsign:    "GRZLY11",
			Type:        aircraft.A400.Identifier,
			Description: aircraft.A400.Description,
		},
	}

	config := configuration.Config{
		AircraftTypes: []string{aircraft.F16.Identifier, aircraft.A400.Identifier},
	}
	actual := filterAircraftByTypes(planes, config)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected '%v' to be the same as '%v'", expected, actual)
	}
}

func TestConvertKnotsToKilometerPerHour(t *testing.T) {
	expected := 185
	actual := ConvertKnotsToKilometersPerHour(100)

	if expected != actual {
		t.Fatalf("expected '%v' to be the same as '%v'", expected, actual)
	}
}

func TestConvertFeetToMeters(t *testing.T) {
	expected := 30
	actual := ConvertFeetToMeters(100)

	if expected != actual {
		t.Fatalf("expected '%v' to be the same as '%v'", expected, actual)
	}
}

func TestSortAircraftByDistance(t *testing.T) {
	aircraft := []AircraftOutput{
		{
			Callsign: "APEX11",
			Distance: 120,
		},
		{
			Callsign: "APEX12",
			Distance: 60,
		},
		{
			Callsign: "APEX13",
			Distance: 10,
		},
	}

	sortedAircraft := SortByDistance(aircraft)

	if sortedAircraft[0].Callsign != "APEX13" || sortedAircraft[1].Callsign != "APEX12" || sortedAircraft[2].Callsign != "APEX11" {
		t.Fatal("List is not sorted by distance")
	}
}

func TestCalculateBearing1(t *testing.T) {

	expected := 320
	actual := int(CalculateBearing(locationMannekenPis, locationElisabethPark))

	if expected != actual {
		t.Fatalf("expected '%v' to be the same as '%v'", expected, actual)
	}
}

func TestCalculateBearing2(t *testing.T) {

	expected := 140
	actual := int(CalculateBearing(locationElisabethPark, locationMannekenPis))

	if expected != actual {
		t.Fatalf("expected '%v' to be the same as '%v'", expected, actual)
	}
}

func TestCalculateBearing3(t *testing.T) {

	expected := 56
	actual := int(CalculateBearing(locationChristRedeemer, locationPyramidGiza))

	if expected != actual {
		t.Fatalf("expected '%v' to be the same as '%v'", expected, actual)
	}
}

func TestCalculateBearing4(t *testing.T) {

	expected := 242
	actual := int(CalculateBearing(locationPyramidGiza, locationChristRedeemer))

	if expected != actual {
		t.Fatalf("expected '%v' to be the same as '%v'", expected, actual)
	}
}

func TestCalculateBearing5(t *testing.T) {
	source := geodist.Coord{
		Lat: 51.42676766088391,
		Lon: 4.623935349264089,
	}

	target := geodist.Coord{
		Lat: 51.426688015979074,
		Lon: 4.63915475148803,
	}

	expected := 91
	actual := int(CalculateBearing(source, target))

	if expected != actual {
		t.Fatalf("expected '%v' to be the same as '%v'", expected, actual)
	}
}
