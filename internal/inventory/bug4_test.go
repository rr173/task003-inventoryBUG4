package inventory

import (
	"testing"
	"time"
)

func TestBug4_ViewPointerIsolation(t *testing.T) {
	s := New()
	now := time.Date(2026, 8, 14, 12, 0, 0, 0, time.UTC)
	inTime := now.Add(time.Hour)
	if _, err := s.Create(CreateInput{SKU: "PTR", Name: "ptr", Stock: 5}, now); err != nil {
		t.Fatal(err)
	}
	p1, err := s.StockIn("PTR", AmountInput{Amount: 1}, inTime)
	if err != nil {
		t.Fatal(err)
	}
	if p1.LastInAt == nil {
		t.Fatal("LastInAt should not be nil")
	}
	*p1.LastInAt = time.Time{}
	p2, _ := s.Get("PTR")
	if p2.LastInAt == nil || p2.LastInAt.IsZero() {
		t.Fatal("BUG: external mutation of view leaked to internal state")
	}
}
