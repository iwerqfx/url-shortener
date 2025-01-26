package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"sync"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name       string
		envVars    map[string]string
		wantConfig Config
		wantPanic  bool
	}{
		{
			name: "valid environment variables",
			envVars: map[string]string{
				"APP_NAME":   "MyApp",
				"LOG_LEVEL":  "debug",
				"LOG_FORMAT": "json",
			},
			wantConfig: Config{
				App: AppConfig{
					Name: "MyApp",
				},
				Log: LogConfig{
					Level:  "debug",
					Format: "json",
				},
			},
			wantPanic: false,
		},
		{
			name: "missing required environment variable",
			envVars: map[string]string{
				"APP_NAME":  "MyApp",
				"LOG_LEVEL": "debug",
			},
			wantConfig: Config{},
			wantPanic:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			instance = nil
			once = sync.Once{}

			if tt.wantPanic {
				assert.Panics(t, func() {
					_ = Get()
				})
			} else {
				assert.NotPanics(t, func() {
					cfg := Get()
					assert.Equal(t, tt.wantConfig, *cfg)
				})
			}

			for key := range tt.envVars {
				os.Unsetenv(key)
			}
		})
	}
}
