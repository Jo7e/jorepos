# OnePass Package

A simplified wrapper around the [1Password SDK for Go](https://github.com/1password/onepassword-sdk-go), making it easier to integrate 1Password credentials in your applications.

## Features

- Simple interface for accessing 1Password vaults
- Integration with 1Password service accounts
- Error handling with descriptive error messages

## Usage

### Basic Usage

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jo7e/jorepos/onepass"
)

func main() {
	// Create a client with service account token from environment variable
	ctx := context.Background()
	client, err := onepass.NewClient(ctx)
	if err != nil {
		fmt.Printf("Error initializing 1Password client: %v\n", err)
		os.Exit(1)
	}

	// List available vaults
	vaults, err := client.GetVaults(ctx)
	if err != nil {
		fmt.Printf("Error retrieving vaults: %v\n", err)
		os.Exit(1)
	}

	// Print vault information
	fmt.Println("Available vaults:")
	for _, vault := range vaults {
		fmt.Printf("- %s (ID: %s)\n", vault.Title, vault.ID)
	}

	// Get a specific vault by title
	vault, err := client.GetVaultByTitle(ctx, "Private")
	if err != nil {
		fmt.Printf("Error retrieving vault: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found vault: %s (ID: %s)\n", vault.Title, vault.ID)
}
```

## Environment Variables

- `OP_SERVICE_ACCOUNT_TOKEN`: The 1Password service account token (required)
