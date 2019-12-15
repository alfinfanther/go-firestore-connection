package products

import (
	"gofirestore/config"
	"gofirestore/globals"
	"gofirestore/models"
	"log"
	"net/http"
)

func ProductAll(w http.ResponseWriter, r *http.Request) {
	ctx, app := config.Connection()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	if err != nil {
		globals.RespondWithJSON(w, http.StatusBadRequest, models.Result{
			Status:  http.StatusBadRequest,
			Message: "Database not found",
		})
		return
	}

	globals.RespondWithJSON(w, http.StatusOK, models.Result{
		Status:  http.StatusOK,
		Message: "Database found",
	})
	return
}
