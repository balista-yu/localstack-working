package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	err := godotenv.Load("../../.env.test")

	if err != nil {
		os.Exit(1)
	}

	code := m.Run()

	os.Exit(code)
}

func TestGetObjectKeys(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			objectKeys, err := getObjectKeys()
			if (err != nil) != tt.wantErr {
				t.Errorf("getObjectKeys() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NotEmpty(t, objectKeys, "The object keys should be anything.")
		})
	}
}
