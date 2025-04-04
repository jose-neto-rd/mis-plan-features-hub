package datadog_test

import (
	"mis-plan-features-hub/internal/infra/datadog"
	logger_mock "mis-plan-features-hub/mocks/infra/logger"
	"testing"
)

func MakeDataDogSut() (*datadog.DatadogService, *logger_mock.MockLogger) {
	logger := &logger_mock.MockLogger{}

	datadog := &datadog.DatadogService{DD_ENV: "test", IsDDEnabled: "test", Logger: logger}

	return datadog, logger
}

func TestDatadog(t *testing.T) {
	t.Run("it should execute contains correctly", func(t *testing.T) {
		datadog, _ := MakeDataDogSut()

		tests := []struct {
			slice  []string
			str    string
			expect bool
		}{
			{[]string{"homologacao", "projeto", "producao"}, "homologacao", true},
			{[]string{"homologacao", "projeto", "producao"}, "teste", false},
			{[]string{}, "homologacao", false},
		}

		for _, tt := range tests {
			result := datadog.Contains(tt.slice, tt.str)
			if result != tt.expect {
				t.Errorf("contains(%v, %s) = %v; want %v", tt.slice, tt.str, result, tt.expect)
			}
		}
	})

	t.Run("it should execute InitDatadog correctly", func(t *testing.T) {
		tests := []struct {
			DD_ENV      string
			IsDDEnabled string
			expectMsg   string
		}{
			{"homologacao", "true", "Monitoring is running"},
			{"projeto", "true", "Monitoring is running"},
			{"producao", "true", "Monitoring is running"},
			{"teste", "true", "Monitoring is not running"},
			{"homologacao", "false", "Monitoring is not running"},
		}

		for _, tt := range tests {
			logger := &logger_mock.MockLogger{}
			datadog := &datadog.DatadogService{
				DD_ENV:      tt.DD_ENV,
				IsDDEnabled: tt.IsDDEnabled,
				Logger:      logger,
			}

			logger.On("Info", tt.expectMsg).Return()

			datadog.InitDatadog()

			logger.AssertExpectations(t)
		}
	})
}
