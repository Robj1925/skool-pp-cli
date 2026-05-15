package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

func newInstallCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "self-install",
		Short: "Install the CLI to your local path",
		Long: `Installs the skool-pp-cli binary to ~/.local/bin and provides 
instructions for adding it to your PATH if it isn't already. 
This avoids the need for sudo/root permissions.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			home, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("getting home directory: %w", err)
			}

			binDir := filepath.Join(home, ".local", "bin")
			if err := os.MkdirAll(binDir, 0755); err != nil {
				return fmt.Errorf("creating bin directory: %w", err)
			}

			exe, err := os.Executable()
			if err != nil {
				return fmt.Errorf("getting executable path: %w", err)
			}

			target := filepath.Join(binDir, "skool-pp-cli")
			
			// Copy binary
			input, err := os.ReadFile(exe)
			if err != nil {
				return fmt.Errorf("reading current binary: %w", err)
			}
			if err := os.WriteFile(target, input, 0755); err != nil {
				return fmt.Errorf("writing binary to %s: %w", target, binDir)
			}

			fmt.Printf("✓ Binary installed to %s\n", target)

			// Check if in PATH
			pathEnv := os.Getenv("PATH")
			if !strings.Contains(pathEnv, binDir) {
				fmt.Println("\n⚠ This directory is not in your PATH.")
				fmt.Println("To add it, run the following command or add it to your shell profile (.zshrc or .bash_profile):")
				
				shell := filepath.Base(os.Getenv("SHELL"))
				if shell == "" {
					shell = "your shell"
				}

				exportCmd := fmt.Sprintf("export PATH=\"%s:$PATH\"", binDir)
				fmt.Printf("\n  %s\n\n", exportCmd)
				
				if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
					profile := ".zshrc"
					if strings.Contains(shell, "bash") {
						profile = ".bash_profile"
					}
					fmt.Printf("To make this permanent, run:\n  echo '%s' >> ~/%s\n", exportCmd, profile)
				}
			} else {
				fmt.Println("✓ Directory is already in your PATH.")
			}

			return nil
		},
	}
}
