package daemon

import (
	"fmt"
	"net/http"

	"github.com/coutzzzzz/mb-go-test/internal/controller"
	"github.com/coutzzzzz/mb-go-test/internal/migration"
	"github.com/coutzzzzz/mb-go-test/internal/repository"
	"github.com/coutzzzzz/mb-go-test/internal/service"
	"github.com/coutzzzzz/mb-go-test/pkg/config"
	"github.com/coutzzzzz/mb-go-test/pkg/database"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func ServeHTTP(router *mux.Router) {
	load := config.Load()

	db := database.NewGorm(load.Database.DSN())
	port := fmt.Sprintf(":%s", load.Service.Port)

	migration.Migrate(db)

	httpServer := BuildHttpServer(router, db)

	http.ListenAndServe(port, httpServer)
}

func BuildHttpServer(router *mux.Router, db *gorm.DB) *mux.Router {
	mmsRepository := repository.NewMMSRepository(db)
	mmsService := service.NewMMSService(mmsRepository)
	mmsHandlers := controller.NewMMSController(mmsService)

	candleService := service.NewCandleService(mmsRepository)
	candleService.Run("BTC", "BRL")
	candleService.Run("ETH", "BRL")

	MapRoutes(router, mmsHandlers)
	return router
}

func MapRoutes(router *mux.Router, handlers *controller.MMSController) {
	router.HandleFunc("/{pair}/mms", handlers.GetMMS).Methods("GET")
}
