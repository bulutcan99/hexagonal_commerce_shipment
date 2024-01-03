package port

import (
	"context"
	"github.com/bulutcan99/shipment/internal/core/domain"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
	UpdatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
	GetPaymentByID(ctx context.Context, id uint64) (*domain.Payment, error)
	ListPayments(ctx context.Context, skip, limit uint64) ([]domain.Payment, error)
	DeletePayment(ctx context.Context, id uint64) error
}

type PaymentService interface {
	CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
	UpdatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
	GetPayment(ctx context.Context, id uint64) (*domain.Payment, error)
	ListPayments(ctx context.Context, skip, limit uint64) ([]domain.Payment, error)
	DeletePayment(ctx context.Context, id uint64) error
}
