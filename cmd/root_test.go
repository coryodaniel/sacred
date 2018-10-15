package cmd

import (
	"testing"
  "os"
  "github.com/spf13/viper"
)

func TestExecute(t *testing.T) {
  t.Run("Flag defaults", func(t *testing.T) {
    RootCmd.Execute()

    if cfgFile != "" {
      t.Errorf("expected config file to be blank by default; got %s", cfgFile)
    }

    if outputDir != "" {
      t.Errorf("expected output directory to be blank by default; got %s", outputDir)
    }

    if verbose {
      t.Errorf("expeced verbose to be false by default")
    }

    if dryRun {
      t.Errorf("expeced dry-run to be false by default")
    }
  })

  t.Run("Reads SACRED_TOKEN from ENV", func(t *testing.T) {
    want := "hello"
    os.Setenv("SACRED_TOKEN", want)

    RootCmd.Execute()
    got := viper.Get("token")

    if want != got {
      t.Errorf("expected to read SACRED_TOKEN from ENV vars")
    }
  })

  t.Run("Reads SACRED_DOMAIN from ENV", func(t *testing.T) {
    want := "example.com"
    os.Setenv("SACRED_DOMAIN", want)

    RootCmd.Execute()
    got := viper.Get("domain")

    if want != got {
      t.Errorf("expected to read SACRED_TOKEN from ENV vars")
    }
  })
}

func TestMergeCredentials(t *testing.T) {
	t.Run("Overrides auth token when the SACRED_TOKEN is set", func(t *testing.T) {
		wantToken := "env-token"
	  os.Setenv("SACRED_TOKEN", wantToken)

		MergeCredentials(&cfg)
		gotToken := cfg.Auth.Token

		if wantToken != gotToken {
			t.Errorf("expected %s got %s", wantToken, gotToken)
		}
	})

	t.Run("Overrides auth domain when the SACRED_DOMAIN is set", func(t *testing.T) {
		wantDomain := "env-domain"
	  os.Setenv("SACRED_DOMAIN", wantDomain)
		RootCmd.Execute()

		MergeCredentials(&cfg)
		gotDomain := cfg.Auth.Domain

		if wantDomain != gotDomain {
			t.Errorf("expected %s got %s", wantDomain, gotDomain)
		}
	})
}
