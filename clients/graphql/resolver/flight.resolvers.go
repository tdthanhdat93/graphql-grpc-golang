package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"FlightBookingAPI/clients/graphql/models"
	dto "FlightBookingAPI/clients/services/models"
	"context"
)

// CreateFlight is the resolver for the createFlight field.
func (r *mutationResolver) CreateFlight(ctx context.Context, input models.NewFlight) (*models.Flight, error) {
	// panic(fmt.Errorf("not implemented: CreateFlight - createFlight"))
	flight := dto.FromGraphQLModel(&input)
	newFlight, err := r.flightService.CreateFlight(flight)
	if err != nil {
		return nil, err
	}

	pRes := newFlight.ToGraphQLModel()
	return pRes, nil
}

// UpdateFlight is the resolver for the updateFlight field.
func (r *mutationResolver) UpdateFlight(ctx context.Context, input models.UpdateFlight) (*models.Flight, error) {
	// panic(fmt.Errorf("not implemented: UpdateFlight - updateFlight"))
	flight := &dto.Flight{
		Id: input.ID,
	}

	if input.PlaneID != nil {
		flight.PlaneId = *input.PlaneID
	}
	if input.Departure != nil {
		flight.Departure = *input.Departure
	}
	if input.TimeDepart != nil {
		flight.TimeDepart = dto.FromGraphQLDateTime(input.TimeDepart)
	}
	if input.TimeArrive != nil {
		flight.TimeArrive = dto.FromGraphQLDateTime(input.TimeArrive)
	}
	if input.AvailableSeat != nil {
		flight.AvailableSeat = int32(*input.AvailableSeat)
	}

	updatedFlight, err := r.flightService.UpdateFlight(flight)
	if err != nil {
		return nil, err
	}

	pRes := updatedFlight.ToGraphQLModel()
	return pRes, nil
}

// BookFlight is the resolver for the bookFlight field.
func (r *mutationResolver) BookFlight(ctx context.Context, input models.Booking) (*models.Flight, error) {
	// panic(fmt.Errorf("not implemented: BookFlight - bookFlight"))
	bookedFlight, err := r.flightService.BookFlight(input.FlightID)
	if err != nil {
		return nil, err
	}

	pRes := bookedFlight.ToGraphQLModel()
	return pRes, nil
}

// DeleteFlight is the resolver for the deleteFlight field.
func (r *mutationResolver) DeleteFlight(ctx context.Context, input models.DeleteFlight) (*models.DeletedFlight, error) {
	// panic(fmt.Errorf("not implemented: DeleteFlight - deleteFlight"))
	err := r.flightService.DeleteFlight(input.ID)
	if err != nil {
		return nil, err
	}
	return &models.DeletedFlight{ID: input.ID}, nil
}

// Flights is the resolver for the flights field.
func (r *queryResolver) Flights(ctx context.Context) ([]*models.Flight, error) {
	// panic(fmt.Errorf("not implemented: Flights - flights"))
	flights, err := r.flightService.ListFlights()
	if err != nil {
		return nil, err
	}

	listRes := make([]*models.Flight, 0)
	for _, flight := range flights {
		listRes = append(listRes, flight.ToGraphQLModel())
	}
	return listRes, err
}

// FindFlights is the resolver for the findFlights field.
func (r *queryResolver) FindFlights(ctx context.Context, param models.FindFlightParams) ([]*models.Flight, error) {
	// panic(fmt.Errorf("not implemented: FindFlights - findFlights"))
	req := &dto.FindFlightParams{}
	if param.Arrival != nil {
		req.Arrival = *param.Arrival
	}
	if param.Departure != nil {
		req.Departure = *param.Departure
	}
	if param.StartTimeRange == nil {
		param.StartTimeRange = &models.InputDateTime{
			Year:  1900,
			Month: 1,
			Day:   1,
		}
	}
	req.StartTimeRange = dto.FromGraphQLDateTime(param.StartTimeRange)
	if param.EndTimeRange == nil {
		param.EndTimeRange = &models.InputDateTime{
			Year:  2100,
			Month: 12,
			Day:   31,
		}
	}
	req.EndTimeRange = dto.FromGraphQLDateTime(param.EndTimeRange)

	flights, err := r.flightService.FindFlights(req)
	if err != nil {
		return nil, err
	}

	listRes := make([]*models.Flight, 0)
	for _, flight := range flights {
		listRes = append(listRes, flight.ToGraphQLModel())
	}
	return listRes, err
}