/*
Package enum shows idiomatic way to implement an enumerated type:
 1. Create a new integer type.
 2. List its values using iota.
 3. Give the type a String function.
Based on yourbasic.org/golang/iota.
*/
package enum

type Direction int

// Cardinal directions.
const (
	North Direction = iota
	South
	East
	West
)

func (d Direction) String() string {
	return [...]string{"North", "South", "East", "West"}[d]
}
