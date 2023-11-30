package chain_of_resp

type Department interface {
	execute(*Patient)
	setNext(Department)
}
