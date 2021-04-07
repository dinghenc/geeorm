package session

import "testing"

var (
	user1 = &User{"Tom", 18}
	user2 = &User{"Sam", 22}
	user3 = &User{"Iack", 23}
)

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	s := NewSession().Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed init test records")
	}
	return s
}

func TestSession_Insert(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Insert(user3)
	if err != nil || affected != 1 {
		t.Fatal("failed to create record")
	}
}

func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Fatal("failed to query records")
	}
}

func TestSession_Limit(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	err := s.Limit(1).Find(&users)
	if err != nil || len(users) != 1 {
		t.Fatal("failed to query with limit condition")
	}
}

func TestSession_Update(t *testing.T) {
	s := testRecordInit(t)
	affected, err1 := s.Where("Name = ?", "Tom").Update("Age", 30)
	u := &User{}
	err2 := s.OrderBy("Age DESC").First(u)
	if err1 != nil || err2 != nil || affected != 1 || u.Age != 30 {
		t.Fatal("failed to update")
	}
}

func TestSession_DeleteAndCount(t *testing.T) {
	s := testRecordInit(t)
	affected, err1 := s.Where("Name = ?", "Tom").Delete()
	count, err2 := s.Count()
	if err1 != nil || err2 != nil || affected != 1 || count != 1 {
		t.Fatal("failed to delete or count")
	}
}
