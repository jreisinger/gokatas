/*
Package counter implements an integer counter that can only be incremented.

Encapsulation (information hiding) is a key aspect of OOP. It prevents clients
from accessing variables or methods of an object. Benefits:
 1) Clients need to understand fewer statements.
 2) Clients don't depend on things that might change.
 3) Clients can't mess with variables directly.

Adapted from the GoPL ch. 6.6 Encapsulation.
*/
package counter

// Counter encapsulates its field(s).
type Counter struct{ n int }

func (c *Counter) Get() int { return c.n } // gettter
func (c *Counter) Inc()     { c.n++ }      // setter
