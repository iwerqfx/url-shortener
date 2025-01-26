package logger

import (
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{
			name: "valid debug level with text format",
			cfg: Config{
				Level:  "debug",
				Format: "text",
			},
			wantErr: false,
		},
		{
			name: "valid info level with JSON format",
			cfg: Config{
				Level:  "info",
				Format: "json",
			},
			wantErr: false,
		},
		{
			name: "invalid log level",
			cfg: Config{
				Level:  "invalid",
				Format: "text",
			},
			wantErr: true,
		},
		{
			name: "valid warn level with text format",
			cfg: Config{
				Level:  "warn",
				Format: "text",
			},
			wantErr: false,
		},
		{
			name: "valid error level with text format",
			cfg: Config{
				Level:  "error",
				Format: "text",
			},
			wantErr: false,
		},
		{
			name: "unsupported format",
			cfg: Config{
				Level:  "info",
				Format: "unsupported",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := New(tt.cfg)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, logger)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, logger)
				_, ok := logger.Handler().(slog.Handler)
				assert.True(t, ok, "Logger handler should implement slog.Handler")
			}
		})
	}
}
