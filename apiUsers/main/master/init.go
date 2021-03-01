package master

import (
	"database/sql"

	"github.com/disebud/api-users-test/main/master/controllers"
	"github.com/disebud/api-users-test/main/master/repositories"
	"github.com/disebud/api-users-test/main/master/usecases"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB) {
	UserRepo := repositories.InitUserRepoImpl(db)
	userUseCase := usecases.InitUserUsecase(UserRepo)
	controllers.UserController(r, userUseCase)

}
