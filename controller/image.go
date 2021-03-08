package controller

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
	"github.com/yakarim/website-walid/cfg"
	"github.com/yakarim/website-walid/database"
	"github.com/yakarim/website-walid/model"
)

// Img type
type Img struct {
	cfg.Cfg
	img model.Img
}

// GetImages image
func (i Img) GetImages(ctx *atreugo.RequestCtx) error {
	imgs, _ := i.img.QueryAll()
	for i, _ := range imgs {
		imgs[i].Body = nil
	}
	return ctx.JSONResponse(imgs, 200)
}

// ImageGet controller
func (i Img) ImageGet(ctx *atreugo.RequestCtx) error {
	key := ctx.UserValue("key")
	b, err := i.img.QueryOne(key.(string))
	if err != nil {
		return ctx.HTTPResponse("", 404)
	}
	n := int64(b.Size)
	Length := strconv.FormatInt(n, 16)
	ctx.Response.Header.Set("Content-Type", b.Type)
	ctx.Response.Header.Set("Content-Length", Length)
	return ctx.HTTPResponseBytes(b.Body, 200)
}

// ImageSave data
func (i Img) ImageSave(ctx *atreugo.RequestCtx) error {
	var img database.Images
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	nameimg, body, types, size := i.Image(file, 500, 0)
	uuid, _ := uuid.NewUUID()

	img.UID = uuid.String()
	img.Body = body
	img.Name = nameimg
	img.Type = types
	img.Size = size
	err = i.img.Create(img)
	if err != nil {
		return i.JSON(ctx, cfg.H{"msg": "cant upload"}, fasthttp.StatusBadRequest)
	}
	return i.JSON(ctx, cfg.H{"url": "media-get-" + img.UID, "id": img.UID}, 200)
}
