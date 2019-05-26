package store

// AddStatus adds a status to a character
func (c *Character) AddStatus(s Status) (err error) {
	if err = c.update(s.Effect); err != nil {
		return
	}
	if err = c.aggregateUpdate(s.EffectAggregate); err != nil {
		return
	}
	return
}
