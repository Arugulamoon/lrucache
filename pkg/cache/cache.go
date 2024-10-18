package cache

import "encoding/json"

type Cache struct {
	MaxSize    int
	RecentData []map[string]string
}

func (c *Cache) ToH() string {
	jsonStr, err := json.Marshal(c.toH())
	if err != nil {
		panic(err)
	}

	return string(jsonStr)
}

func (c *Cache) toH() map[string]string {
	data := make(map[string]string, 0) // TODO: MaxSize?
	for _, e := range c.RecentData {
		for k, v := range e {
			data[k] = v
		}
	}
	return data
}

func (c *Cache) Read(k string) string {
	found := c.delete(k)
	c.RecentData = append(c.RecentData, found)
	return found[k]
}

func (c *Cache) Write(k, v string) string {
	if len(c.RecentData) >= c.MaxSize {
		c.RecentData = c.RecentData[1:]
	}
	c.RecentData = append(c.RecentData,
		map[string]string{k: v})
	return v
}

func (c *Cache) Delete(k string) string {
	found := c.delete(k)
	return found[k]
}

func (c *Cache) delete(k string) map[string]string {
	var found map[string]string
	recentData := make([]map[string]string, 0)
	for _, e := range c.RecentData {
		for j := range e {
			if j != k {
				recentData = append(recentData, e)
			} else {
				found = e
			}
		}
	}
	c.RecentData = recentData
	return found
}

func (c *Cache) Count() int {
	return len(c.RecentData)
}

func (c *Cache) Clear() int {
	count := c.Count()
	c.RecentData = make([]map[string]string, 0)
	return count
}
