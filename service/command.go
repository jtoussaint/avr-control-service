package service

//
// A CommandHandler processing incoming commands to mutate data
//
type CommandHandler interface {
	Handle()
}
