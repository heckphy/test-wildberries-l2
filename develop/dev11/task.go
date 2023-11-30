package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/heckphy/wb-l2/develop/dev11/config"
	"github.com/heckphy/wb-l2/develop/dev11/internal/handlers"
	"github.com/heckphy/wb-l2/develop/dev11/internal/middleware"
	"github.com/heckphy/wb-l2/develop/dev11/internal/models"
	"github.com/heckphy/wb-l2/develop/dev11/internal/repositories"
	"github.com/heckphy/wb-l2/develop/dev11/internal/services"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

func main() {
	// TODO: app entity
	cfg := config.LoadConfig()
	ctx := context.Background()

	eventRepository := repositories.NewEventRepository()

	eventRepository.Add(ctx, &models.Event{
		Uid:         "11111",
		Title:       "first event",
		Description: "descr1",
		StartsAt:    time.Date(2023, 11, 23, 0, 0, 0, 0, time.UTC),
		EndsAt:      time.Date(2023, 11, 29, 0, 0, 0, 0, time.UTC),
	})

	eventRepository.Add(ctx, &models.Event{
		Uid:         "22222",
		Title:       "second event",
		Description: "descr2",
		StartsAt:    time.Date(2023, 12, 01, 0, 0, 0, 0, time.UTC),
		EndsAt:      time.Date(2023, 12, 06, 0, 0, 0, 0, time.UTC),
	})

	eventRepository.Add(ctx, &models.Event{
		Uid:         "33333",
		Title:       "third event",
		Description: "descr3",
		StartsAt:    time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
		EndsAt:      time.Date(2023, 12, 27, 0, 0, 0, 0, time.UTC),
	})

	eventService := services.NewEventService(eventRepository)
	eventHandler := handlers.NewEventHandler(eventService)

	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", eventHandler.Create)
	mux.HandleFunc("/update_event", eventHandler.Update)
	mux.HandleFunc("/delete_event", eventHandler.Delete)
	mux.HandleFunc("/events_for_day", eventHandler.GetForDay)
	mux.HandleFunc("/events_for_week", eventHandler.GetForWeek)
	mux.HandleFunc("/events_for_month", eventHandler.GetForMonth)

	handler := middleware.NewLogger(mux)

	err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), handler)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
