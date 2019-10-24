/// Pkg created for connecting with the database, making it transparent (both MariaDB and Cassandra).
package connector

/// Connector element
type Connector struct {
	// Info of the database
}

func (Connector) Connect() error {
	// Connects to the DB
	return nil
}

func (Connector) Close() error {
	// Connects to the DB
	return nil
}

func (Connector) Execute(query string) (int, error) {
	return 0, nil
}
