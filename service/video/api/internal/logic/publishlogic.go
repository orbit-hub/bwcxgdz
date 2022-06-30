package logic

import (
	"bwcxgdz/dousheng/common/oss"
	"bwcxgdz/dousheng/service/video/api/internal/svc"
	"bwcxgdz/dousheng/service/video/api/internal/types"
	"bwcxgdz/dousheng/service/video/rpc/video"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"

	"github.com/gofrs/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	r      *http.Request
	svcCtx *svc.ServiceContext
}

func NewPublishLogic(ctx context.Context, r *http.Request, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		r:      r,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *types.PublishRequest) (resp *types.PublishResponse, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	err = l.r.ParseMultipartForm(oss.MaxFileSize)
	if err != nil {
		return nil, err
	}
	file, handler, err := l.r.FormFile("data")
	if err != nil {
		return nil, err
	}
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	u2, err := uuid.NewV4()
	fileName := u2.String() + "." + "mp4"
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bucket, err := oss.OssClient.Bucket(oss.BucketName)
	if err != nil {
		return nil, err
	}
	err = bucket.PutObject("dousheng/video/"+fileName, io.Reader(file))
	if err != nil {
		return nil, err
	}
	url := "https://dousheng-bwcxgdz.oss-cn-hangzhou.aliyuncs.com/dousheng/video/" + fileName
	jpegBytes, err := readFrameAsJpeg(url)
	if err != nil {
		return nil, err
	}
	u3, err := uuid.NewV4()
	coverName := u3.String() + "." + "jpg"
	if err != nil {
		return nil, err
	}
	err = bucket.PutObject("dousheng/img/"+coverName, bytes.NewReader(jpegBytes))
	_, err = l.svcCtx.VideoRpc.Publish(l.ctx, &video.VideoInfoRequest{
		UserId:   uid,
		PlayUrl:  url,
		CoverUrl: "http://dousheng-bwcxgdz.oss-cn-hangzhou.aliyuncs.com/dousheng/img/" + coverName,
		Title:    req.Title,
	})
	if err != nil {
		return nil, err
	}
	return &types.PublishResponse{}, nil
}

// ReadFrameAsJpeg
// 从视频流中截取一帧并返回 需要在本地环境中安装ffmpeg并将bin添加到环境变量
func readFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), err
}
