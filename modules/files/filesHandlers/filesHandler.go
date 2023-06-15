package filesHandlers

import (
	"github.com/winai-pgm-itsystem/all-day-shop/config"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/files/filesUsecases"
)

type IFilesHandler interface{}

type filesHandler struct {
	cfg           config.IConfig
	filesUsecases filesUsecases.IFilesUsecase
}

func FilesHandler(cfg config.IConfig, filesUsecases filesUsecases.IFilesUsecase) IFilesHandler {

	return &filesHandler{
		cfg:           cfg,
		filesUsecases: filesUsecases,
	}
}
