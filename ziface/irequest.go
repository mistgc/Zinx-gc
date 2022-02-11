package ziface

/*
	Interface 'IRequset':
		Actually, this interface combine client connection information 
		and request data into a single 'Request'.
*/

type IRequset interface {
	// Get current connection
	GetConnection() IConnection

	// Get request data
	GetData() []byte
}
