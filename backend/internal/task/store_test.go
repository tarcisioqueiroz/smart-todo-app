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
