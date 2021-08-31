package fitbuddy

func InitCommands() *CmdStorage {
	storage := InitCmdStorage()

	storage.addCallback("/start", StartCallback)
	storage.addCallback("/menu", PrintMenuCallback)

	return storage
}
