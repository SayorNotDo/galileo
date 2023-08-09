package service

import (
	"context"
	v1 "galileo/api/file/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *FileService) GetOssStsToken(ctx context.Context, req *emptypb.Empty) (*v1.OssStsTokenReply, error) {
	stsRes, err := s.uc.GetOssStsToken(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.OssStsTokenReply{
		AccessKey:     stsRes.AccessKey,
		AccessSecret:  stsRes.AccessSecret,
		Expiration:    stsRes.Expiration,
		SecurityToken: stsRes.SecurityToken,
		Endpoint:      stsRes.EndPoint,
		BucketName:    stsRes.BucketName,
		Region:        stsRes.Region,
		Url:           stsRes.Url,
	}, nil
}

func (s *FileService) UploadFile(ctx context.Context, req *v1.UploadFileRequest) (*v1.UploadFileReply, error) {
	url, err := s.uc.UploadFile(ctx, req.Name, req.Type, req.Content)
	if err != nil {
		return nil, err
	}
	return &v1.UploadFileReply{Url: url}, nil
}

func (s *FileService) DeleteFile(ctx context.Context, req *v1.DeleteFileRequest) (*v1.DeleteFileReply, error) {
	return &v1.DeleteFileReply{}, nil
}

func (s *FileService) DownloadFile(ctx context.Context, req *v1.DownloadFileRequest) (*v1.DownloadFileReply, error) {
	return &v1.DownloadFileReply{}, nil
}
