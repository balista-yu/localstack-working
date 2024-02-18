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

func TestGetTableNames(t *testing.T) {
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
			tableNames, err := getTableNames()
			if (err != nil) != tt.wantErr {
				t.Errorf("getTableNames() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, "working-demo-table", tableNames, "The two words should be the same.")
		})
	}
}
