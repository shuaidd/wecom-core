package media

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"

	"github.com/shuaidd/wecom-core/internal/client"
	"github.com/shuaidd/wecom-core/types/media"
)

// UploadImage 上传图片
// 上传图片得到图片URL，该URL永久有效
// 文档: https://developer.work.weixin.qq.com/document/path/90256
func (s *Service) UploadImage(ctx context.Context, imagePath string) (*media.UploadImageResponse, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %w", err)
	}
	defer file.Close()

	return s.UploadImageFromReader(ctx, file, filepath.Base(imagePath))
}

// UploadImageFromReader 从 io.Reader 上传图片
func (s *Service) UploadImageFromReader(ctx context.Context, reader io.Reader, filename string) (*media.UploadImageResponse, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("media", filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err := io.Copy(part, reader); err != nil {
		return nil, fmt.Errorf("failed to copy file content: %w", err)
	}

	contentType := writer.FormDataContentType()
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	return client.PostMultipartAndUnmarshal[media.UploadImageResponse](
		s.client, ctx, "/cgi-bin/media/uploadimg", body.Bytes(), contentType,
	)
}

// UploadMedia 上传临时素材
// 素材上传得到media_id，该media_id仅三天内有效
// 文档: https://developer.work.weixin.qq.com/document/path/90253
func (s *Service) UploadMedia(ctx context.Context, mediaType media.MediaType, mediaPath string) (*media.UploadMediaResponse, error) {
	file, err := os.Open(mediaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open media file: %w", err)
	}
	defer file.Close()

	return s.UploadMediaFromReader(ctx, mediaType, file, filepath.Base(mediaPath))
}

// UploadMediaFromReader 从 io.Reader 上传临时素材
func (s *Service) UploadMediaFromReader(ctx context.Context, mediaType media.MediaType, reader io.Reader, filename string) (*media.UploadMediaResponse, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("media", filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err := io.Copy(part, reader); err != nil {
		return nil, fmt.Errorf("failed to copy file content: %w", err)
	}

	contentType := writer.FormDataContentType()
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	query := url.Values{}
	query.Set("type", string(mediaType))

	return client.PostMultipartAndUnmarshalWithQuery[media.UploadMediaResponse](
		s.client, ctx, "/cgi-bin/media/upload", query, body.Bytes(), contentType,
	)
}

// GetMedia 获取临时素材
// 返回素材的内容。如果素材过大，需使用Range分块下载
// 文档: https://developer.work.weixin.qq.com/document/path/90254
func (s *Service) GetMedia(ctx context.Context, mediaID string) ([]byte, error) {
	query := url.Values{}
	query.Set("media_id", mediaID)

	return s.client.GetMedia(ctx, "/cgi-bin/media/get", query, nil)
}

// GetMediaWithRange 使用Range分块获取临时素材
// rangeHeader 格式如: "bytes=0-1023"
func (s *Service) GetMediaWithRange(ctx context.Context, mediaID string, rangeHeader string) ([]byte, error) {
	query := url.Values{}
	query.Set("media_id", mediaID)

	headers := map[string]string{
		"Range": rangeHeader,
	}

	return s.client.GetMedia(ctx, "/cgi-bin/media/get", query, headers)
}

// GetJSSDKMedia 获取高清语音素材
// 获取从JSSDK的uploadVoice接口上传的临时语音素材，格式为speex，16K采样率
// 文档: https://developer.work.weixin.qq.com/document/path/90255
func (s *Service) GetJSSDKMedia(ctx context.Context, mediaID string) ([]byte, error) {
	query := url.Values{}
	query.Set("media_id", mediaID)

	return s.client.GetMedia(ctx, "/cgi-bin/media/get/jssdk", query, nil)
}

// UploadByURL 异步上传临时素材
// 生成异步上传任务，支持最高200M的大文件
// 文档: https://developer.work.weixin.qq.com/document/path/96219
func (s *Service) UploadByURL(ctx context.Context, req *media.UploadByURLRequest) (*media.UploadByURLResponse, error) {
	return client.PostAndUnmarshal[media.UploadByURLResponse](s.client, ctx, "/cgi-bin/media/upload_by_url", req)
}

// GetUploadByURLResult 查询异步上传任务结果
// 文档: https://developer.work.weixin.qq.com/document/path/96219
func (s *Service) GetUploadByURLResult(ctx context.Context, jobID string) (*media.GetUploadByURLResultResponse, error) {
	req := &media.GetUploadByURLResultRequest{
		JobID: jobID,
	}
	return client.PostAndUnmarshal[media.GetUploadByURLResultResponse](s.client, ctx, "/cgi-bin/media/get_upload_by_url_result", req)
}
