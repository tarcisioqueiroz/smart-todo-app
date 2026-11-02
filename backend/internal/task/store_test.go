package task

import "testing"

func TestCreateAssignsSequentialIDs(t *testing.T) {
	s := NewStore()
	t1 := s.Create("tarefa 1")
	t2 := s.Create("tarefa 2")
	t3 := s.Create("tarefa 3")

	if t1.ID != 1 {
		t.Errorf("Primeira tarefa: ID = %d, esperado %d", t1.ID, 1)
	}
	if t2.ID != 2 {
		t.Errorf("Segunda tarefa: ID = %d, esperado %d", t2.ID, 2)
	}
	if t3.ID != 3 {
		t.Errorf("Terceira tarefa: ID = %d, esperado %d", t3.ID, 3)
	}
}

func TestGetUnknownIDReturnsFalse(t *testing.T) {
	s := NewStore()
	_, ok := s.Get(99)
	if ok {
		t.Errorf("Get(99): ok = %v, esperado false", ok)
	}
}

func TestUpdateKeepsFieldsNotSent(t *testing.T) {
	s := NewStore()
	t1 := s.Create("tarefa 1")
	done := true
	_, ok := s.Update(t1.ID, nil, &done)
	if !ok {
		t.Fatalf("Update (done): tarefa não encontrada")
	}

	title := "novo title"
	updated, ok := s.Update(t1.ID, &title, nil)
	if !ok {
		t.Fatalf("Update (title): tarefa não encontrada")
	}
	if updated.Title != title {
		t.Errorf("Update devolveu Title = %q, esperado %q", updated.Title, title)
	}

	got, ok := s.Get(t1.ID)

	if !ok {
		t.Fatalf("Get(%d): tarefa não encontrada", t1.ID)
	}
	if got.Title != title {
		t.Errorf("Title = %q, esperado %q", got.Title, title)
	}

	if !got.Done {
		t.Errorf("Done = %v, esperado true", got.Done)
	}
}

func TestUpdateUnknownIDReturnsFalse(t *testing.T) {
	s := NewStore()
	done := true
	_, ok := s.Update(99, nil, &done)
	if ok {
		t.Errorf("Update(99): ok = true, esperado false")
	}
}
