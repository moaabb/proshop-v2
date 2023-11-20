package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moaabb/proshop-v2/asset_svc/infra/config"

	"go.uber.org/zap"
)

type UploadHandler struct {
	l   *zap.Logger
	cfg *config.Config
}

func NewHandler(z *zap.Logger, cfg *config.Config) *UploadHandler {
	return &UploadHandler{
		cfg: cfg,
		l:   z,
	}

}

func (uh *UploadHandler) UploadImage(c *gin.Context) {
	err := c.Request.ParseMultipartForm(2 << 20)
	if err != nil {
		uh.l.Error("could not parse image data from multipart", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "could not upload file",
		})
		return
	}

	file, _ := c.FormFile("image")
	name := strings.Split(file.Filename, ".")[0]
	extension := strings.Split(file.Filename, ".")[1]

	savePath := fmt.Sprintf("%s/images/%s-%s.%s", uh.cfg.SaveFileBasePath, name, uuid.New(), extension)
	imageDbPath := strings.Split(savePath, uh.cfg.SaveFileBasePath)[1]
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		uh.l.Error("could not parse image data from multipart", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "could not upload file",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"image": imageDbPath,
	})
}
