package db

import (
	"errors"
)

type RecordDBAPI[S any, T any] struct {
	IdGetter  func(T) S
	isDefault func(T) bool
}

func NewAPI[S any, T any](idGetter func(T) S, defaultChecker func(T) bool) *RecordDBAPI[S, T] {
	return &RecordDBAPI[S, T]{
		IdGetter:  idGetter,
		isDefault: defaultChecker,
	}
}

func (s *RecordDBAPI[S, T]) GetAll() ([]T, error) {
	var value []T
	r := DB.Find(&value)
	if err := HandleDBError(r); err != nil {
		return value, err
	}

	return value, nil
}

func (s *RecordDBAPI[S, T]) Create(v T) (T, error) {
	value := v
	r := DB.Create(&value)
	if err := HandleDBError(r); err != nil {
		return value, err
	}

	return s.GetByID(s.IdGetter(value))
}

func (s *RecordDBAPI[S, T]) GetByID(id S) (T, error) {
	var value T

	r := DB.Find(&value, id)
	if err := HandleDBError(r); err != nil {
		return value, err
	}

	if s.isDefault(value) {
		return value, errors.New(ERR_RECORD_NOT_FOUND)
	}

	return value, nil
}

func (s *RecordDBAPI[S, T]) Delete(id S) error {
	var value T
	r := DB.Delete(&value, id)
	if err := HandleDBError(r); err != nil {
		return err
	}

	return nil
}
