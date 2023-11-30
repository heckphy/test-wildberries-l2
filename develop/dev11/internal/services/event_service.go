package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/heckphy/wb-l2/develop/dev11/internal/dtos"
	"github.com/heckphy/wb-l2/develop/dev11/internal/models"
)

type eventRepository interface {
	Add(context.Context, *models.Event)
	DeleteByUid(context.Context, string) error
	GetByUid(context.Context, string) (*models.Event, error)
	GetForDateRange(context.Context, time.Time, time.Time) []models.Event
}

type eventService struct {
	eventRepository eventRepository
}

func NewEventService(eventRepository eventRepository) *eventService {
	return &eventService{
		eventRepository: eventRepository,
	}
}

func (s *eventService) Create(ctx context.Context, dto dtos.EventCreateDTO) (*models.Event, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return &models.Event{}, err
	}

	event := &models.Event{
		Uid:         uid.String(),
		Title:       dto.Title,
		Description: dto.Description,
		StartsAt:    dto.StartsAt,
		EndsAt:      dto.EndsAt,
	}

	s.eventRepository.Add(ctx, event)

	return event, nil
}

func (s *eventService) Update(ctx context.Context, dto dtos.EventUpdateDTO) (*models.Event, error) {
	event, err := s.eventRepository.GetByUid(ctx, dto.Uid)
	if err != nil {
		return &models.Event{}, err
	}

	event.Title = dto.Title
	event.Description = dto.Description
	event.StartsAt = dto.StartsAt
	event.EndsAt = dto.EndsAt

	s.eventRepository.Add(ctx, event)

	return event, nil
}

func (s *eventService) DeleteByUid(ctx context.Context, dto dtos.EventDeleteDTO) error {
	return s.eventRepository.DeleteByUid(ctx, dto.Uid)
}

func (s *eventService) GetForDay(ctx context.Context, fromDate time.Time) []models.Event {
	toDate := fromDate.Add(24 * time.Hour)

	return s.eventRepository.GetForDateRange(ctx, fromDate, toDate)
}

func (s *eventService) GetForWeek(ctx context.Context, fromDate time.Time) []models.Event {
	toDate := fromDate.Add(7 * 24 * time.Hour)

	return s.eventRepository.GetForDateRange(ctx, fromDate, toDate)
}

func (s *eventService) GetForMonth(ctx context.Context, fromDate time.Time) []models.Event {
	toDate := fromDate.Add(30 * 24 * time.Hour)

	return s.eventRepository.GetForDateRange(ctx, fromDate, toDate)
}
