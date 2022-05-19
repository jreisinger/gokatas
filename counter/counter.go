/*
Package counter implements an integer counter that can only be incremented.

Encapsulation (information hiding) is a key aspect of OOP. It prevents clients
from accessing variables or methods of an object. Benefits:

 1) Clients need to understand fewer statements.
 2) Clients don't depend on things that might change.
 3) Clients can't mess with variables directly.

Adapted from the GoPL (gopl.io) ch. 6.6 Encapsulation.
*/
package counter

type Counter struct{ n int } // counter encapsulates its field(s)
func (c *Counter) Get() int  { return c.n } // getter
func (c *Counter) Inc()      { c.n++ }      // setter
