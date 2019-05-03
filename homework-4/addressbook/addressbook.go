// implements addressbook & sorting
package addressbook

type Person struct {
	Name    string
	Age     int
	Phones  []string
	Address string
}

//ByAge implements sort interface for Person's ages
type ByAge []Person

//ByAddr implements sort interface for Person's names
type ByAddr []Person

// Len - num records of the address book
func (ba ByAge) Len() int {
	return len(ba)
}

// Less - is need to swap records?
func (ba ByAge) Less(i, j int) bool {
	return ba[i].Age < ba[j].Age
}

// Swap swaps records
func (ba ByAge) Swap(i, j int) {
	ba[i], ba[j] = ba[j], ba[i]
}

// Len - num records of the address book
func (ba ByAddr) Len() int {
	return len(ba)
}

// Less - is need to swap records?
func (ba ByAddr) Less(i, j int) bool {
	return ba[i].Address < ba[j].Address
}

// Swap swaps records
func (ba ByAddr) Swap(i, j int) {
	ba[i], ba[j] = ba[j], ba[i]
}
