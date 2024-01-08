package maskerspam

import (
	"fmt"
)

type producer interface {
	produce() ([]string, error)
}

type presenter interface {
	present(data []string) error
}

// Service is structure for masking url service
// Including inside 2 fields:
// producer - for data provider unit
// presenter - for data presenter unit.
type Service struct {
	prod producer
	pres presenter
}

// NewService is constructor of Service
func NewService(prod producer, pres presenter) *Service {
	return &Service{
		prod: prod,
		pres: pres,
	}
}

// Run is method for start Service working
func (s *Service) Run() error {
	data, err := s.prod.produce()
	if err != nil {
		return fmt.Errorf("service.producer.produce: %w", err)
	}

	data = s.process(data)

	if err = s.pres.present(data); err != nil {
		return fmt.Errorf("service.presentor.present: %w", err)
	}

	return nil
}

func (s *Service) process(data []string) []string {
	const maxRoutineCount = 10
	resultData := make([]string, 0, cap(data))
	tasks := make(chan string)
	results := make(chan string)

	for i := 0; i < maxRoutineCount; i++ {
		go s.worker(tasks, results)
	}

	go func() {
		for _, task := range data {
			tasks <- task
		}

		close(tasks)
	}()

	for rowIndex := 0; rowIndex < len(data); rowIndex++ {
		resultData = append(resultData, <-results)
	}
	close(results)

	return resultData
}

func (s *Service) worker(tasks <-chan string, result chan<- string) {
	for task := range tasks {
		result <- s.maskingURL(task)
	}
}

func (s *Service) maskingURL(str string) string {
	const symbolsDetectedCount int = 7

	startURLIndex := 0
	isMasking := false
	buffer := []byte(str)

	for index := range buffer {
		if buffer[index] == 'h' && string(buffer[index:index+7]) == "http://" {
			startURLIndex = index + symbolsDetectedCount
			isMasking = true
		}

		if startURLIndex != 0 && index >= startURLIndex && isMasking {
			if buffer[index] == ' ' {
				isMasking = false

				continue
			}

			buffer[index] = '*'
		}
	}

	return string(buffer)
}
