/*
Package warehouse lets a Blacksmith source or destination be considered as a data
warehouse. This allows TLT, operations, and queries with Pythonistic SQL direcly
on top of databases with no other layer of abstraction.

Note: It is implemented by the third-party package sqlike to easily leverage the
standard database/sql and run operations / queries on top of the SQL database. See
Go module at https://pkg.go.dev/github.com/nunchistudio/blacksmith-modules/sqlike
for more details.
*/
package warehouse
