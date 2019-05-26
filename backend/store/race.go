package store

// AddRace adds a race to a character
func (c *Character) AddRace(r Race) (err error) {
	if err = c.update(r.Update); err != nil {
		return
	}
	if err = c.aggregateUpdate(r.Aggregate); err != nil {
		return
	}
	return
}
