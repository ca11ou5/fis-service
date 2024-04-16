package handler

import (
	"context"
	"fis/internal/pb"
	"fis/internal/service"
	"github.com/skbt-ecom/logging"
)

type Handler struct {
	service *service.Service
	log     *logging.Logger
}

func NewHandler(service *service.Service, log *logging.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) FullMappingAndSave(ctx context.Context, req *pb.FullMappingAndSaveRequest) (*pb.FullMappingAndSaveResponse, error) {
	err := h.service.FullMappingAndSave(req.Application, req.RequestId)
}

func (h *Handler) Score(ctx context.Context, req *pb.ScoreRequest) (*pb.ScoreResponse, error) {

}

func (h *Handler) UpdateStatus(ctx context.Context, req *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error) {
	err := h.service.UpdateStatus()
}

func (h *Handler) SendToFis(ctx context.Context, req *pb.SendToFisRequest) (*pb.SendToFisResponse, error) {
	err := h.service.SendToFis()
}
