package delivery

import (
	"fmt"
	"ikuzports/features/event"
	"ikuzports/features/participant"
	"ikuzports/features/user"
	"ikuzports/middlewares"
	"ikuzports/utils/helper"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type ParticipantDelivery struct {
	participantService participant.ServiceInterface
	userService        user.ServiceInterface
	eventService       event.ServiceInterface
	oauth              *oauth2.Config
}

func New(service participant.ServiceInterface, e *echo.Echo, userService user.ServiceInterface, eventService event.ServiceInterface, oauthConfig *oauth2.Config) {
	handler := &ParticipantDelivery{
		participantService: service,
		userService:        userService,
		eventService:       eventService,
		oauth:              oauthConfig,
	}
	e.POST("/participants", handler.Create, middlewares.JWTMiddleware())
}

func (delivery *ParticipantDelivery) Create(c echo.Context) error {
	participantInput := ParticipantRequest{}
	errBind := c.Bind(&participantInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(participantInput)
	dataCore.UserID = middlewares.ExtractTokenUserId(c)

	err := delivery.participantService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data "+err.Error()))
	}

	if participantInput.Token != "" {
		dataUser, errUsr := delivery.userService.GetById(dataCore.UserID)
		log.Printf("This is your email: %v", dataUser.Email)
		if errUsr != nil {
			log.Printf("Unable to retrieve user: %v", errUsr)
		}
		dataEvent, errEvnt := delivery.eventService.GetByID(participantInput.EventID)
		log.Printf("This is your event: %v", dataEvent.Name)
		if errEvnt != nil {
			log.Printf("Unable to retrieve user: %v", errEvnt)
		}

		startString := dataEvent.StartDate.Format("2006-01-02")
		endString := dataEvent.EndDate.Format("2006-01-02")

		startdateTime := fmt.Sprintf("%sT00:00:00+07:00", startString)
		enddateTime := fmt.Sprintf("%sT00:00:00+07:00", endString)
		log.Printf("This is your startdate: %v", startdateTime)
		location := fmt.Sprintf("%s, %s", dataEvent.Address, dataEvent.City)
		events := &calendar.Event{
			Summary:     dataEvent.Name,
			Description: dataEvent.Description,
			Start: &calendar.EventDateTime{
				DateTime: startdateTime,
				TimeZone: "Asia/Jakarta",
			},
			End: &calendar.EventDateTime{
				DateTime: enddateTime,
				TimeZone: "Asia/Jakarta",
			},
			Attendees: []*calendar.EventAttendee{
				{Email: dataUser.Email},
			},
			Location: location,
		}

		tokenOauth := &oauth2.Token{AccessToken: participantInput.Token}
		log.Printf("This is your tokenOauth: %v", tokenOauth)
		client := delivery.oauth.Client(c.Request().Context(), tokenOauth)
		srv, err := calendar.NewService(c.Request().Context(), option.WithHTTPClient(client))
		if err != nil {
			log.Printf("Unable to retrieve Calendar client: %v", err)
		}

		_, errCr := srv.Events.Insert("primary", events).SendUpdates("all").Do()
		if errCr != nil {
			log.Printf("Unable to create event. %v\n", err)
		}

	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Succes create data"))
}
