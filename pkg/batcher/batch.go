package batcher

import (
	"fmt"
	"time"
)

// BatchProcessor определяет структуру для универсальной батчевой обработки.
type BatchProcessor[T any] struct {
	BatchSize int
	Delay     time.Duration
	Process   func(batch []T) error
}

// NewBatchProcessor создает новый экземпляр BatchProcessor с заданными параметрами.
func NewBatchProcessor[T any](batchSize int, delay time.Duration, processFunc func(batch []T) error) *BatchProcessor[T] {
	return &BatchProcessor[T]{
		BatchSize: batchSize,
		Delay:     delay,
		Process:   processFunc,
	}
}

// Run выполняет батчевую обработку данных.
func (bp *BatchProcessor[T]) Run(data []T) error {
	for i := 0; i < len(data); i += bp.BatchSize {
		end := i + bp.BatchSize

		if end > len(data) {
			end = len(data)
		}

		batch := data[i:end]
		err := bp.Process(batch)
		if err != nil {
			return fmt.Errorf("ошибка при обработке батча: %w", err)
		}

		if end < len(data) {
			time.Sleep(bp.Delay)
		}
	}
	return nil
}
