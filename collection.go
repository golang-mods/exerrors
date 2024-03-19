package exerrors

import (
	"errors"
	"sync"
)

func Map0[T any](collection []T, iteratee func(item T, index int) error) error {
	errs := make([]error, len(collection))

	for index, item := range collection {
		errs[index] = iteratee(item, index)
	}

	return errors.Join(errs...)
}

func Map[T, R any](collection []T, iteratee func(item T, index int) (R, error)) ([]R, error) {
	result := make([]R, len(collection))

	return result, Map0(collection, func(item T, index int) error {
		var err error
		result[index], err = iteratee(item, index)
		return err
	})
}

func ParallelMap0[T any](collection []T, iteratee func(item T, index int) error) error {
	errs := make([]error, len(collection))

	var wg sync.WaitGroup
	wg.Add(len(collection))
	for index, item := range collection {
		go func() {
			errs[index] = iteratee(item, index)
			wg.Done()
		}()
	}
	wg.Wait()

	return errors.Join(errs...)
}

func ParallelMap[T, R any](collection []T, iteratee func(item T, index int) (R, error)) ([]R, error) {
	result := make([]R, len(collection))

	return result, ParallelMap0(collection, func(item T, index int) error {
		var err error
		result[index], err = iteratee(item, index)
		return err
	})
}
