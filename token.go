package main

import (
	"github.com/thanhpk/randstr"
)

func (s *Service) NewToken(uid uint) string{
	str := randstr.String(32)
	token := &Token{
		Uid:   uid,
		Str:   str,
	}
	tx := s.DB.Begin()
	if tx.Create(&token).RowsAffected != 1 {
		tx.Rollback()
		return ""
	}
	tx.Commit()
	return str
}
