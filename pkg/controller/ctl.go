package controller

var _ Interface = &CIController{}
var _ Interface = &CDController{}
// Interface ....
type Interface interface {
	Run(addr string, stop <-chan struct{}) error
	Handle(addr string)
}
