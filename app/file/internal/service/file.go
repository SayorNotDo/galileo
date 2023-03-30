package service

import (
	"context"
	"galileo/app/file/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	v1 "galileo/api/file/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FileService struct {
	v1.UnimplementedFileServer
	uc     *biz.FileUseCase
	logger *log.Helper
}

func NewFileService(uc *biz.FileUseCase, logger log.Logger) *FileService {
	return &FileService{
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "module", "file")),
	}
}

func (s *FileService) GetOssStsToken(ctx context.Context, req *emptypb.Empty) (*v1.OssStsTokenReply, error) {
	return &v1.OssStsTokenReply{}, nil
}

func (s *FileService) CreateFile(ctx context.Context, req *v1.CreateFileRequest) (*v1.CreateFileReply, error) {
	return &v1.CreateFileReply{}, nil
}

func (s *FileService) UploadFile(ctx context.Context, req *v1.UploadFileRequest) (*v1.UploadFileReply, error) {
	url, err := s.uc.UploadFile(ctx, req.Name, req.Type, req.Content)
	if err != nil {
		return nil, err
	}
	response := &v1.UploadFileReply{}
	response.Url = url
	return response, nil
}

func (s *FileService) UpdateFile(ctx context.Context, req *v1.UpdateFileRequest) (*v1.UpdateFileReply, error) {
	return &v1.UpdateFileReply{}, nil
}

func (s *FileService) DeleteFile(ctx context.Context, req *v1.DeleteFileRequest) (*v1.DeleteFileReply, error) {
	return &v1.DeleteFileReply{}, nil
}

func (s *FileService) GetFile(ctx context.Context, req *v1.GetFileRequest) (*v1.GetFileReply, error) {
	return &v1.GetFileReply{}, nil
}

func (s *FileService) ListFile(ctx context.Context, req *v1.ListFileRequest) (*v1.ListFileReply, error) {
	return &v1.ListFileReply{}, nil
}
