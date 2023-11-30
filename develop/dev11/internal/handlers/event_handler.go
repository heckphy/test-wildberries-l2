package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/heckphy/wb-l2/develop/dev11/internal/dtos"
	"github.com/heckphy/wb-l2/develop/dev11/internal/models"
)

type eventService interface {
	Create(context.Context, dtos.EventCreateDTO) (*models.Event, error)
	Update(context.Context, dtos.EventUpdateDTO) (*models.Event, error)
	DeleteByUid(context.Context, dtos.EventDeleteDTO) error
	GetForDay(context.Context, time.Time) []models.Event
	GetForWeek(context.Context, time.Time) []models.Event
	GetForMonth(context.Context, time.Time) []models.Event
}

type eventHandler struct {
	eventService eventService
}

func NewEventHandler(eventService eventService) *eventHandler {
	return &eventHandler{
		eventService: eventService,
	}
}

func (h *eventHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: "Method not allowed",
		})
		return
	}

	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: "Wrong Content-Type header",
		})
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: "Failed to parse input",
		})
		return
	}

	startsAt, err := time.Parse("2006-01-02", r.FormValue("startsAt"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: err.Error(),
		})
		return
	}

	endsAt, err := time.Parse("2006-01-02", r.FormValue("endsAt"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: err.Error(),
		})
		return
	}

	createDTO := dtos.EventCreateDTO{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		StartsAt:    startsAt,
		EndsAt:      endsAt,
	}

	event, err := h.eventService.Create(r.Context(), createDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: "Internal server error",
		})
		return
	}

	getDTO := dtos.EventGetDTO{
		Uid:         event.Uid,
		Title:       event.Title,
		Description: event.Description,
		StartsAt:    event.StartsAt,
		EndsAt:      event.EndsAt,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dtos.ResultDTO{
		Result: getDTO,
	})
}

func (h *eventHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: "Method not allowed",
		})
		return
	}

	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: "Wrong Content-Type header",
		})
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: "Failed to parse input",
		})
		return
	}

	startsAt, err := time.Parse("2006-01-02", r.FormValue("startsAt"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: err.Error(),
		})
		return
	}

	endsAt, err := time.Parse("2006-01-02", r.FormValue("endsAt"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: err.Error(),
		})
		return
	}

	updateDTO := dtos.EventUpdateDTO{
		Uid:         r.FormValue("uid"),
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		StartsAt:    startsAt,
		EndsAt:      endsAt,
	}

	event, err := h.eventService.Update(r.Context(), updateDTO)
	getDTO := dtos.EventGetDTO{
		Uid:         event.Uid,
		Title:       event.Title,
		Description: event.Description,
		StartsAt:    event.StartsAt,
		EndsAt:      event.EndsAt,
	}

	json.NewEncoder(w).Encode(dtos.ResultDTO{
		Result: getDTO,
	})
}

func (h *eventHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: "Method not allowed",
		})
		return
	}

	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: "Wrong Content-Type header",
		})
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: "Failed to parse input",
		})
		return
	}

	deleteDTO := dtos.EventDeleteDTO{
		Uid: r.FormValue("uid"),
	}

	h.eventService.DeleteByUid(r.Context(), deleteDTO)

	w.WriteHeader(http.StatusNoContent)
}

func (h *eventHandler) GetFor(
	w http.ResponseWriter,
	r *http.Request,
	getfor func(ctx context.Context, date time.Time) []models.Event,
) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: "Method not allowed",
		})
		return
	}

	rawDate := r.URL.Query().Get("date")
	parsedDate, err := time.Parse("2006-01-02", rawDate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dtos.ErrorDTO{
			Error: "Invalid date provided",
		})
		return
	}

	events := getfor(r.Context(), parsedDate)

	json.NewEncoder(w).Encode(dtos.ResultDTO{
		Result: events,
	})
}

func (h *eventHandler) GetForDay(w http.ResponseWriter, r *http.Request) {
	h.GetFor(w, r, h.eventService.GetForDay)
}

func (h *eventHandler) GetForWeek(w http.ResponseWriter, r *http.Request) {
	h.GetFor(w, r, h.eventService.GetForWeek)
}

func (h *eventHandler) GetForMonth(w http.ResponseWriter, r *http.Request) {
	h.GetFor(w, r, h.eventService.GetForMonth)
}
