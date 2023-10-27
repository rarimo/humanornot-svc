package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/requests"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/responses"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/core"
)

var TickerDuration = time.Second

func GetVerifyStatus(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		Log(r).WithError(err).Error("failed to upgrade connection")
		return
	}
	defer c.Close()

	req, err := requests.NewVerifyStatusRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse request")
		if err := c.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseInternalServerErr, fmt.Sprintf("bad request: %s", err.Error())),
		); err != nil {
			Log(r).WithError(err).Error("failed to write to websocket")
		}
		return
	}

	ticker := time.NewTicker(TickerDuration)

	for range ticker.C {
		user, err := KYCService(r).GetVerifyStatus(req)
		switch {
		case errors.Is(err, core.ErrUserNotFound):
			continue
		case err != nil:
			Log(r).WithError(err).Error("failed to get user")
			if err := c.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseInternalServerErr, "failed to write to websocket"),
			); err != nil {
				Log(r).WithError(err).Error("failed to write to websocket")
			}
			return
		}

		if user.ClaimID.String() != "" {
			if err := c.WriteJSON(responses.NewVerifyStatus(user)); err != nil {
				Log(r).WithError(err).Error("failed to write to websocket")
			}
			return
		}
	}
}
