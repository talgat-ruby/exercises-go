package service

import "errors"

func (s *serviceExpence) ServiceBalance(id int) (map[string]float64, error) {
	if id <= 0 {
		return nil, errors.New("incorrect id")
	}
	return s.dbExpence.DBBalance(id)
}
