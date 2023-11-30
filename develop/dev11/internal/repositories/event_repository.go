package repositories

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/heckphy/wb-l2/develop/dev11/internal/models"
)

type eventRepository struct {
	events map[string]models.Event
	mu     sync.RWMutex
}

func NewEventRepository() *eventRepository {
	return &eventRepository{
		events: make(map[string]models.Event),
	}
}

func (r *eventRepository) Add(ctx context.Context, event *models.Event) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.events[event.Uid] = *event
}

func (r *eventRepository) DeleteByUid(ctx context.Context, uid string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.events[uid]; !ok {
		return errors.New("event not found")
	}

	delete(r.events, uid)
	return nil
}

func (r *eventRepository) GetByUid(ctx context.Context, uid string) (*models.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if _, ok := r.events[uid]; !ok {
		return &models.Event{}, errors.New("event not found")
	}

	event := r.events[uid]

	return &event, nil
}

func (r *eventRepository) GetForDateRange(ctx context.Context, fromDate time.Time, toDate time.Time) []models.Event {
	r.mu.RLock()
	defer r.mu.RUnlock()

	fmt.Println("FROM DATE", fromDate)
	fmt.Println("TO DATE", toDate)

	res := []models.Event{}

	for _, event := range r.events {
		fmt.Println("EVENT", event)
		fmt.Println("EVENT.StartsAt", event.StartsAt)
		fmt.Println()

		if event.StartsAt.After(fromDate) && event.StartsAt.Before(toDate) {
			res = append(res, event)
		}
	}

	return res
}
