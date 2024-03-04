package models

type Customer struct {
	ID uint64 `pg:",pk"`
	BranchID   uint `pg:"on_delete:CASCADE"`
	Name     string `pg:",notnull"`
	PAN      string `pg:",notnull, unique"`
	DOB      string
	Age      int
	PhoneNo  int
	Address  string
	Accounts []*Account `pg:"many2many:account_to_customers"`
}
