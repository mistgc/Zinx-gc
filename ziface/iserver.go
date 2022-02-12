package ziface

type IServer interface {
	// Start server
	Start()

	// Stop server
	Stop()

	// Run server
	Serve()

	// Register a router for client connection to handle business
	AddRouter(router IRouter)
}
