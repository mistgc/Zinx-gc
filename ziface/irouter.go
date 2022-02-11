package ziface

type IRouter interface {
	// The hook method before handle connection business
	PreHandle(request IRequset)

	// The main method to handle connection business
	Handle(request IRequset)

	// The hook method after handle connection business
	PostHandle(request IRequset)
}
