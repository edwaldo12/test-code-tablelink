package grpc

import (
	"context"
	"time"

	pb "test_tablelink/internal/delivery/grpc"
	"test_tablelink/internal/domain"
	"test_tablelink/internal/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ItemServer struct {
	pb.UnimplementedItemServiceServer
	uc usecase.ItemUsecase
}

func NewItemServer(uc usecase.ItemUsecase) *ItemServer {
	return &ItemServer{uc: uc}
}

func (s *ItemServer) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	if req.Name == "" || req.Price <= 0 {
		return nil, status.Error(codes.InvalidArgument, "name and price are required")
	}

	item := &domain.Item{
		Name:   req.Name,
		Price:  req.Price,
		Status: int(req.Status),
	}

	if err := s.uc.Create(ctx, item); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateItemResponse{Uuid: item.UUID}, nil
}

func (s *ItemServer) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	item := &domain.Item{
		UUID:   req.Uuid,
		Name:   req.Name,
		Price:  req.Price,
		Status: int(req.Status),
	}

	if err := s.uc.Update(ctx, item); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateItemResponse{}, nil
}

func (s *ItemServer) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	if err := s.uc.HardDelete(ctx, req.Uuid); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.DeleteItemResponse{}, nil
}

func (s *ItemServer) ListItems(ctx context.Context, req *pb.ListItemsRequest) (*pb.ListItemsResponse, error) {
	limit := int(req.Limit)
	offset := int(req.Offset)

	items, err := s.uc.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &pb.ListItemsResponse{
		Items: make([]*pb.Item, len(items)),
	}

	for i, item := range items {
		response.Items[i] = &pb.Item{
			Uuid:      item.UUID,
			Name:      item.Name,
			Price:     item.Price,
			Status:    int32(item.Status),
			CreatedAt: item.CreatedAt.Format(time.RFC3339),
			UpdatedAt: item.UpdatedAt.Format(time.RFC3339),
			DeletedAt: formatTimePtr(item.DeletedAt),
		}
	}

	return response, nil
}
