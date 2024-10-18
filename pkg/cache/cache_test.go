package cache_test

import (
	"testing"

	"lrucache/pkg/cache"
)

func TestInitialization(t *testing.T) {
	c := &cache.Cache{
		MaxSize:    3,
		RecentData: make([]map[string]string, 0),
	}

	got := c.ToH()
	want := "{}"

	if got != want {
		t.Errorf("want %s; got %s", want, got)
	}
}

func TestWrite(t *testing.T) {
	c := &cache.Cache{
		MaxSize:    3,
		RecentData: make([]map[string]string, 0),
	}

	got := c.Write("key1", "val1")
	want := "val1"

	if got != want {
		t.Errorf("want %s; got %s", want, got)
	}
}

func TestToH(t *testing.T) {
	c := &cache.Cache{
		MaxSize:    3,
		RecentData: make([]map[string]string, 0),
	}
	c.Write("key1", "val1")

	got := c.ToH()
	want := "{\"key1\":\"val1\"}"

	if got != want {
		t.Errorf("want %s; got %s", want, got)
	}
}

func TestToHMultiple(t *testing.T) {
	c := &cache.Cache{
		MaxSize:    3,
		RecentData: make([]map[string]string, 0),
	}
	c.Write("key1", "val1")
	c.Write("key2", "val2")

	got := c.ToH()
	want := "{\"key1\":\"val1\",\"key2\":\"val2\"}"

	if got != want {
		t.Errorf("want %s; got %s", want, got)
	}
}

func TestWriteOverMaxSizeExpiresOldestKVs(t *testing.T) {
	c := &cache.Cache{
		MaxSize:    3,
		RecentData: make([]map[string]string, 0),
	}
	c.Write("key1", "val1")
	c.Write("key2", "val2")
	c.Write("key3", "val3")
	c.Write("key4", "val4")

	got := c.ToH()
	want := "{\"key2\":\"val2\",\"key3\":\"val3\",\"key4\":\"val4\"}"

	if got != want {
		t.Errorf("want %s; got %s", want, got)
	}
}

func TestReadNotFound(t *testing.T) {
	c := &cache.Cache{
		MaxSize:    3,
		RecentData: make([]map[string]string, 0),
	}
	c.Write("key1", "val1")
	c.Write("key2", "val2")
	c.Write("key3", "val3")
	c.Write("key4", "val4")

	got := c.Read("key1")
	// TODO: Change to nil?
	want := ""

	if got != want {
		t.Errorf("want %s; got %s", want, got)
	}
}

func TestRead(t *testing.T) {
	c := &cache.Cache{
		MaxSize:    3,
		RecentData: make([]map[string]string, 0),
	}
	c.Write("key1", "val1")
	c.Write("key2", "val2")
	c.Write("key3", "val3")
	c.Write("key4", "val4")

	got := c.Read("key2")
	want := "val2"

	if got != want {
		t.Errorf("want %s; got %s", want, got)
	}
}

func TestReadMarksKVAsRecent(t *testing.T) {
	c := &cache.Cache{
		MaxSize:    3,
		RecentData: make([]map[string]string, 0),
	}
	c.Write("key1", "val1")
	c.Write("key2", "val2")
	c.Write("key3", "val3")
	c.Write("key4", "val4")
	c.Read("key2")
	c.Write("key5", "val5")
	c.Write("key6", "val6")

	got := c.ToH()
	want := "{\"key2\":\"val2\",\"key5\":\"val5\",\"key6\":\"val6\"}"

	if got != want {
		t.Errorf("want %s; got %s", want, got)
	}
}

func TestCount(t *testing.T) {
	c := &cache.Cache{
		MaxSize:    3,
		RecentData: make([]map[string]string, 0),
	}
	c.Write("key1", "val1")
	c.Write("key2", "val2")
	c.Write("key3", "val3")

	got := c.Count()
	want := 3

	if got != want {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestOverwriteReturn(t *testing.T) {
	c := &cache.Cache{
		MaxSize:    3,
		RecentData: make([]map[string]string, 0),
	}
	c.Write("key5", "val5")

	got := c.Write("key5", "value5-overwrite")
	want := "value5-overwrite"

	if got != want {
		t.Errorf("want %s; got %s", want, got)
	}
}

func TestOverwriteRead(t *testing.T) {
	c := &cache.Cache{
		MaxSize:    3,
		RecentData: make([]map[string]string, 0),
	}
	c.Write("key5", "val5")
	c.Write("key5", "value5-overwrite")

	got := c.Read("key5")
	want := "value5-overwrite"

	if got != want {
		t.Errorf("want %s; got %s", want, got)
	}
}

func TestDeleteReturn(t *testing.T) {
	c := &cache.Cache{
		MaxSize:    3,
		RecentData: make([]map[string]string, 0),
	}
	c.Write("key5", "value5-overwrite")

	got := c.Delete("key5")
	want := "value5-overwrite"

	if got != want {
		t.Errorf("want %s; got %s", want, got)
	}
}

func TestDeleteCount(t *testing.T) {
	c := &cache.Cache{
		MaxSize:    3,
		RecentData: make([]map[string]string, 0),
	}
	c.Write("key1", "val1")
	c.Delete("key1")

	got := c.Count()
	want := 0

	if got != want {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestClear(t *testing.T) {
	c := &cache.Cache{
		MaxSize:    3,
		RecentData: make([]map[string]string, 0),
	}
	c.Write("key1", "val1")

	got := c.Clear()
	want := 1

	if got != want {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestClearReturnsEmptyToH(t *testing.T) {
	c := &cache.Cache{
		MaxSize:    3,
		RecentData: make([]map[string]string, 0),
	}
	c.Write("key1", "val1")
	c.Clear()

	got := c.ToH()
	want := "{}"

	if got != want {
		t.Errorf("want %s; got %s", want, got)
	}
}
