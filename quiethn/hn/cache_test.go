package hn

import (
	"testing"
	"time"
)

func TestCache_Add(t *testing.T) {
	id := 123
	c := NewCache(1 * time.Second)
	i := Item{ID: id}

	c.Add(i)
	date1 := c.value[id].CreationDate
	c.Add(i)
	date2 := c.value[id].CreationDate

	if date1 == date2 {
		t.Error("The creation date should have been updated.")
	}
}
