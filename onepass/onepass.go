package onepass

import (
	"context"
	"fmt"
	"os"

	"github.com/1password/onepassword-sdk-go"
)

// Client represents a 1Password client
type Client struct {
	opc *onepassword.Client
}

// NewClient creates a new 1Password client using the service account token
func NewClient(ctx context.Context) (*Client, error) {
	token := os.Getenv("OP_SERVICE_ACCOUNT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("%w: OP_SERVICE_ACCOUNT_TOKEN", ErrEnvVarNotSet)
	}

	client, err := onepassword.NewClient(
		ctx,
		onepassword.WithServiceAccountToken(token),
		onepassword.WithIntegrationInfo("jorepos/onepass", "v0.0.1"),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create 1Password client: %w", err)
	}

	return &Client{
		opc: client,
	}, nil
}

// GetVaults returns all accessible vaults
func (c *Client) GetVaults(ctx context.Context) ([]onepassword.VaultOverview, error) {
	return c.opc.Vaults().List(ctx)
}

// GetVaultByID returns a vault by its ID
func (c *Client) GetVaultByTitle(ctx context.Context, title string) (*onepassword.VaultOverview, error) {
	if title == "" {
		return nil, fmt.Errorf("vault title cannot be empty")
	}

	vaults, err := c.GetVaults(ctx)
	if err != nil {
		return nil, err
	}

	for i, vault := range vaults {
		if vault.Title == title {
			return &vaults[i], nil
		}
	}

	return nil, fmt.Errorf("vault with title '%s' not found", title)
}
