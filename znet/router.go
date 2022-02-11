package znet

import (
	"Zinx/ziface"
)

// When implementing router, first embed 'BaseRouter', 
// and then override the methods of 'BaseRouter' according to the requirement.

type BaseRouter struct {}

/*
	Because some 'Router' don't want implement business 'PreHandle' and business 'PostHandle',
	if 'Router' inherit 'BaseRouter' completely,
	it doesn't need to implement 'PreHandle' and 'PostHandle'.
*/

// The hook method before handle connection business
func (br *BaseRouter) PreHandle(request ziface.IRequset) {}

// The main method to handle connection business
func (br *BaseRouter) Handle(request ziface.IRequset) {}

// The hook method after handle connection business
func (br *BaseRouter) PostHandle(request ziface.IRequset) {}
