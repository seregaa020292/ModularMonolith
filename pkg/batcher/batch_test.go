package batcher

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

// testProcessBatch создает замыкание, которое записывает каждый обработанный батч в предоставленный срез processedBatches.
// Возвращает ошибку для определенного батча, если это необходимо.
func testProcessBatch[T any](processedBatches *[][]T, failOnBatch int) func(batch []T) error {
	return func(batch []T) error {
		if len(*processedBatches)+1 == failOnBatch {
			return errors.New("test error")
		}
		*processedBatches = append(*processedBatches, batch)
		return nil
	}
}

func TestBatchProcessor_Run(t *testing.T) {
	tests := []struct {
		name            string
		data            []int
		batchSize       int
		failOnBatch     int
		expectedBatches [][]int
		expectError     bool
	}{
		{
			name:            "Successful processing",
			data:            []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			batchSize:       3,
			failOnBatch:     0, // Не вызывать ошибку
			expectedBatches: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}},
			expectError:     false,
		},
		{
			name:            "Error on third batch",
			data:            []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			batchSize:       3,
			failOnBatch:     3, // Ошибка на третьем батче
			expectedBatches: [][]int{{1, 2, 3}, {4, 5, 6}},
			expectError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var processedBatches [][]int
			processFunc := testProcessBatch(&processedBatches, tt.failOnBatch)

			batchProcessor := NewBatchProcessor(tt.batchSize, 0*time.Second, processFunc)

			err := batchProcessor.Run(tt.data)
			if tt.expectError && err == nil {
				t.Errorf("%s: expected an error, but got nil", tt.name)
			} else if !tt.expectError && err != nil {
				t.Errorf("%s: unexpected error: %v", tt.name, err)
			}

			if !reflect.DeepEqual(processedBatches, tt.expectedBatches) {
				t.Errorf("%s: expected batches %v, but got %v", tt.name, tt.expectedBatches, processedBatches)
			}
		})
	}
}
