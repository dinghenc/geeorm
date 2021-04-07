package session

import (
	"geeorm/log"
	"testing"
)

type Account struct {
	ID       int `geeorm:"PRIMARY KEY"`
	Password string
}

func (account *Account) BeforeInsert(s *Session) error {
	log.Info("before insert", account)
	account.ID += 1000
	return nil
}

func (account *Account) AfterQuery(s *Session) error {
	log.Info("after query", account)
	account.Password = "******"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	s := NewSession().Model(&Account{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, err1 := s.Insert(&Account{1, "123456"}, &Account{2, "qwerty"})

	u := &Account{}

	err2 := s.First(u)
	if err1 != nil || err2 != nil || u.ID != 1001 || u.Password != "******" {
		t.Fatalf("failed to call hooks: %v", u)
	}
}
