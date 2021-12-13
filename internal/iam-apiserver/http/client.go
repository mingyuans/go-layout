package http

var factoryClient Factory

type Factory interface {
	WXClient() WXClient
}

func Client() Factory {
	return factoryClient
}

func SetFactory(client Factory) {
	factoryClient = client
}
