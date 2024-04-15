package data

import (
	"benchmark/proto/flightsearch"
	"github.com/goccy/go-json"
	"log"
)

func MockDataAsString() []byte {
	dataByte, err := json.Marshal(MockData())
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err)
	}

	return dataByte
}

func MockData() *flightsearch.SearchFlightsResponse {
	return &flightsearch.SearchFlightsResponse{
		Flights: []*flightsearch.FlightInfo{
			{
				Airline:       "THY",
				FlightNumber:  "TK123",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T00:00:00Z",
				ArrivalTime:   "2024-04-15T01:00:00Z",
				Price:         350.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T00:00:00Z",
						ArrivalTime:    "2024-04-15T01:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "SunExpress",
				FlightNumber:  "XQ127",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T01:00:00Z",
				ArrivalTime:   "2024-04-15T02:00:00Z",
				Price:         360.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T01:00:00Z",
						ArrivalTime:    "2024-04-15T02:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "SunExpress",
				FlightNumber:  "XQ126",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T02:00:00Z",
				ArrivalTime:   "2024-04-15T03:00:00Z",
				Price:         340.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T02:00:00Z",
						ArrivalTime:    "2024-04-15T03:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "SunExpress",
				FlightNumber:  "XQ125",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T03:00:00Z",
				ArrivalTime:   "2024-04-15T04:00:00Z",
				Price:         330.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T03:00:00Z",
						ArrivalTime:    "2024-04-15T04:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "SunExpress",
				FlightNumber:  "XQ124",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T04:00:00Z",
				ArrivalTime:   "2024-04-15T05:00:00Z",
				Price:         320.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T04:00:00Z",
						ArrivalTime:    "2024-04-15T05:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "SunExpress",
				FlightNumber:  "XQ123",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T05:00:00Z",
				ArrivalTime:   "2024-04-15T06:00:00Z",
				Price:         310.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T05:00:00Z",
						ArrivalTime:    "2024-04-15T06:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "Pegasus",
				FlightNumber:  "PC227",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T06:00:00Z",
				ArrivalTime:   "2024-04-15T07:00:00Z",
				Price:         600.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T06:00:00Z",
						ArrivalTime:    "2024-04-15T07:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "Pegasus",
				FlightNumber:  "PC226",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T07:00:00Z",
				ArrivalTime:   "2024-04-15T08:00:00Z",
				Price:         350.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T07:00:00Z",
						ArrivalTime:    "2024-04-15T08:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "Pegasus",
				FlightNumber:  "PC225",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T08:00:00Z",
				ArrivalTime:   "2024-04-15T09:00:00Z",
				Price:         550.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T08:00:00Z",
						ArrivalTime:    "2024-04-15T09:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "Pegasus",
				FlightNumber:  "PC224",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T09:00:00Z",
				ArrivalTime:   "2024-04-15T10:00:00Z",
				Price:         530.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T09:00:00Z",
						ArrivalTime:    "2024-04-15T10:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "Pegasus",
				FlightNumber:  "PC223",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T10:00:00Z",
				ArrivalTime:   "2024-04-15T11:00:00Z",
				Price:         540.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T10:00:00Z",
						ArrivalTime:    "2024-04-15T11:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "THY",
				FlightNumber:  "TK129",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T11:00:00Z",
				ArrivalTime:   "2024-04-15T12:00:00Z",
				Price:         420.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T11:00:00Z",
						ArrivalTime:    "2024-04-15T12:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "THY",
				FlightNumber:  "TK128",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T12:00:00Z",
				ArrivalTime:   "2024-04-15T13:00:00Z",
				Price:         440.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T12:00:00Z",
						ArrivalTime:    "2024-04-15T13:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "THY",
				FlightNumber:  "TK127",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T13:00:00Z",
				ArrivalTime:   "2024-04-15T14:00:00Z",
				Price:         420.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T13:00:00Z",
						ArrivalTime:    "2024-04-15T14:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "Delta Airlines",
				FlightNumber:  "DL123",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T14:00:00Z",
				ArrivalTime:   "2024-04-15T15:00:00Z",
				Price:         700.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T14:00:00Z",
						ArrivalTime:    "2024-04-15T15:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "THY",
				FlightNumber:  "TK126",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T15:00:00Z",
				ArrivalTime:   "2024-04-15T16:00:00Z",
				Price:         600.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T15:00:00Z",
						ArrivalTime:    "2024-04-15T16:00:00Z",
						LayoverMinutes: 60,
					},
				},
			}, {
				Airline:       "THY",
				FlightNumber:  "TK124",
				Origin:        "JFK",
				Destination:   "LAX",
				DepartureTime: "2024-04-15T16:00:00Z",
				ArrivalTime:   "2024-04-15T17:00:00Z",
				Price:         650.00,
				IsDirect:      true,
				Connections: []*flightsearch.ConnectionInfo{
					{
						Airline:        "United Airlines",
						FlightNumber:   "UA789",
						Origin:         "ORD",
						Destination:    "LAX",
						DepartureTime:  "2024-04-15T16:00:00Z",
						ArrivalTime:    "2024-04-15T17:00:00Z",
						LayoverMinutes: 60,
					},
				},
			},
		},
	}
}
