package database

//***these code implement into package-initialization

var connection string

// the init function is a special function wich will be the first to implement by the system
// when the package was called
// it's like a magic method in php __construct
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
