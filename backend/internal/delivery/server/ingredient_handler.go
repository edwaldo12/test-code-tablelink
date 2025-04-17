package grpc

import (
	"context"
	"database/sql"
	"time"

	pb "test_tablelink/internal/delivery/grpc"
	"test_tablelink/internal/domain"
	"test_tablelink/internal/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IngredientServer struct {
	pb.UnimplementedIngredientServiceServer
	uc usecase.IngredientUsecase
}

func NewIngredientServer(uc usecase.IngredientUsecase) *IngredientServer {
	return &IngredientServer{uc: uc}
}

func (s *IngredientServer) CreateIngredient(ctx context.Context, req *pb.CreateIngredientRequest) (*pb.CreateIngredientResponse, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}

	ing := &domain.Ingredient{
		Name:         req.Name,
		CauseAllergy: req.CauseAlergy,
		Type:         int(req.Type),
		Status:       int(req.Status),
	}

	if err := s.uc.Create(ctx, ing); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateIngredientResponse{Uuid: ing.UUID}, nil
}

func (s *IngredientServer) UpdateIngredient(ctx context.Context, req *pb.UpdateIngredientRequest) (*pb.UpdateIngredientResponse, error) {
	ing := &domain.Ingredient{
		UUID:         req.Uuid,
		Name:         req.Name,
		CauseAllergy: req.CauseAlergy,
		Type:         int(req.Type),
		Status:       int(req.Status),
	}

	if err := s.uc.Update(ctx, ing); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateIngredientResponse{}, nil
}

func (s *IngredientServer) DeleteIngredient(ctx context.Context, req *pb.DeleteIngredientRequest) (*pb.DeleteIngredientResponse, error) {
	if err := s.uc.HardDelete(ctx, req.Uuid); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.DeleteIngredientResponse{}, nil
}

func (s *IngredientServer) ListIngredients(ctx context.Context, req *pb.ListIngredientsRequest) (*pb.ListIngredientsResponse, error) {
	limit := int(req.Limit)
	offset := int(req.Offset)

	ingredients, err := s.uc.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &pb.ListIngredientsResponse{
		Ingredients: make([]*pb.Ingredient, len(ingredients)),
	}

	for i, ing := range ingredients {
		response.Ingredients[i] = &pb.Ingredient{
			Uuid:        ing.UUID,
			Name:        ing.Name,
			CauseAlergy: ing.CauseAllergy,
			Type:        int32(ing.Type),
			Status:      int32(ing.Status),
			CreatedAt:   formatNullTime(ing.CreatedAt),
			UpdatedAt:   formatNullTime(ing.UpdatedAt),
			DeletedAt:   formatNullTime(ing.DeletedAt),
		}
	}

	return response, nil
}

func formatTimePtr(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(time.RFC3339)
}

func formatNullTime(nt sql.NullTime) string {
	if !nt.Valid {
		return ""
	}
	return nt.Time.Format(time.RFC3339)
}
