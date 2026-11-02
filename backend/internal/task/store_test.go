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
		t.Errorf("Bsuca tarefa com ID desconhecido: \nesperado: %v\nretornou: %v", false, ok)
	}
}

func TestUpdateKeepsFieldsNotSent(t *testing.T) {
	s := NewStore()
	t1 := s.Create("tarefa 1")
	done := true
	s.Update(t1.ID, nil, &done)

	title := "novo title"
	s.Update(t1.ID, &title, nil)

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
