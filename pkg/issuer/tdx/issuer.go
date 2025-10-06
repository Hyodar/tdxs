package tdx

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"

	rawtdx "github.com/Hyodar/tdxs/internal/constellation/attestation/tdx"
	"github.com/google/go-tdx-guest/abi"
	"github.com/google/go-tdx-guest/proto/tdx"

	"github.com/Hyodar/tdxs/pkg/api"
	"github.com/Hyodar/tdxs/pkg/issuer"
	"github.com/Hyodar/tdxs/pkg/logger"
)

type TDXIssuer struct {
	issuer.Issuer

	logger  logger.Logger
	backend *rawtdx.Issuer
}

type TDXMetadata struct {
	MrTd    string `json:"mrtd"`    // Measurement of initial TD contents (hex)
	MrOwner string `json:"mrowner"` // Software-defined ID for TD owner (hex)
	MrSeam  string `json:"mrseam"`  // Measurement of TDX Module (hex)
	Rtmr0   string `json:"rtmr0"`   // Runtime measurement register 0 (hex)
	Rtmr1   string `json:"rtmr1"`   // Runtime measurement register 1 (hex)
	Rtmr2   string `json:"rtmr2"`   // Runtime measurement register 2 (hex)
	Rtmr3   string `json:"rtmr3"`   // Runtime measurement register 3 (hex)
	XFAM    string `json:"xfam"`    // Extended features available mask (hex)
}

func NewTDXIssuer(logger logger.Logger) *TDXIssuer {
	return &TDXIssuer{
		backend: rawtdx.NewIssuer(logger),
		logger:  logger,
	}
}

func (i *TDXIssuer) Start(_ context.Context) error {
	return nil
}

func (i *TDXIssuer) Issue(ctx context.Context, req *api.IssueRequest) *api.IssueResponse {
	doc, err := i.backend.Issue(ctx, req.UserData, req.Nonce)
	if err != nil {
		return &api.IssueResponse{Error: err}
	}
	return &api.IssueResponse{Document: doc}
}

func (i *TDXIssuer) Metadata(ctx context.Context, req *api.MetadataRequest) *api.MetadataResponse {
	userData := []byte(issuer.MetadataUserData)
	nonce := []byte(issuer.MetadataNonce)

	doc, err := i.backend.Issue(ctx, userData, nonce)
	if err != nil {
		return &api.MetadataResponse{Error: err}
	}

	metadata, err := i.extractMetadata(doc)
	if err != nil {
		return &api.MetadataResponse{Error: fmt.Errorf("extract metadata: %w", err)}
	}

	return &api.MetadataResponse{
		IssuerType: string(issuer.IssuerTypeTDX),
		UserData:   userData,
		Nonce:      nonce,
		Metadata:   metadata,
	}
}

func (i *TDXIssuer) extractMetadata(doc []byte) (*TDXMetadata, error) {
	var attDoc struct {
		RawQuote []byte `json:"RawQuote"`
		UserData []byte `json:"UserData"`
	}

	if err := json.Unmarshal(doc, &attDoc); err != nil {
		return nil, fmt.Errorf("unmarshal attestation document: %w", err)
	}

	quotePb, err := abi.QuoteToProto(attDoc.RawQuote)
	if err != nil {
		return nil, fmt.Errorf("parse TDX quote: %w", err)
	}

	quote, ok := quotePb.(*tdx.QuoteV4)
	if !ok {
		return nil, fmt.Errorf("unexpected quote type: %T", quotePb)
	}

	metadata := &TDXMetadata{}

	if quote.TdQuoteBody != nil {
		metadata.XFAM = prefixedHexEncode(quote.TdQuoteBody.Xfam)
		metadata.MrTd = prefixedHexEncode(quote.TdQuoteBody.MrTd)
		metadata.MrOwner = prefixedHexEncode(quote.TdQuoteBody.MrOwner)
		metadata.MrSeam = prefixedHexEncode(quote.TdQuoteBody.MrSeam)

		if len(quote.TdQuoteBody.Rtmrs) > 0 {
			metadata.Rtmr0 = prefixedHexEncode(quote.TdQuoteBody.Rtmrs[0])
		}
		if len(quote.TdQuoteBody.Rtmrs) > 1 {
			metadata.Rtmr1 = prefixedHexEncode(quote.TdQuoteBody.Rtmrs[1])
		}
		if len(quote.TdQuoteBody.Rtmrs) > 2 {
			metadata.Rtmr2 = prefixedHexEncode(quote.TdQuoteBody.Rtmrs[2])
		}
		if len(quote.TdQuoteBody.Rtmrs) > 3 {
			metadata.Rtmr3 = prefixedHexEncode(quote.TdQuoteBody.Rtmrs[3])
		}
	}

	return metadata, nil
}

func prefixedHexEncode(data []byte) string {
	return "0x" + hex.EncodeToString(data)
}

