package tdx

import (
	"context"
	"fmt"

	rawtdx "github.com/Hyodar/tdxs/internal/constellation/attestation/tdx"
	"github.com/Hyodar/tdxs/internal/constellation/config"

	"github.com/Hyodar/tdxs/pkg/api"
	"github.com/Hyodar/tdxs/pkg/logger"
	"github.com/Hyodar/tdxs/pkg/validator"
)

type TDXValidator struct {
	validator.Validator

	logger  logger.Logger
	backend *rawtdx.Validator
}

type TDXValidatorConfig struct {
	*config.QEMUTDX `yaml:",inline"`
}

func NewTDXValidator(cfg *TDXValidatorConfig, logger logger.Logger) *TDXValidator {
	return &TDXValidator{
		backend: rawtdx.NewValidator(cfg.QEMUTDX, logger),
		logger:  logger,
	}
}

func (i *TDXValidator) Start(_ context.Context) error {
	return nil
}

func (i *TDXValidator) Validate(ctx context.Context, req *api.ValidateRequest) *api.ValidateResponse {
	if i.backend == nil {
		return &api.ValidateResponse{Error: fmt.Errorf("backend not initialized")}
	}

	userData, err := i.backend.Validate(ctx, req.Document, req.Nonce)
	if err != nil {
		return &api.ValidateResponse{Error: err}
	}
	return &api.ValidateResponse{UserData: userData, Valid: true}
}
