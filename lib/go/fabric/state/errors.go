package state

import "fmt"

//type InvalidStateObjectError struct {
//	MSG string
//}

type KeyNotFoundError struct {
	Key       string
	Namespace string
	MSG       string
}

func (e *KeyNotFoundError) Error() string {
	return fmt.Sprintf("Key %s not found in namespace %s => msg: %s", e.Key, e.Namespace, e.MSG)
}

type KeyAlreadyExistsError struct {
	Key       string
	Namespace string
	MSG       string
}

func (e *KeyAlreadyExistsError) Error() string {
	return fmt.Sprintf("Key %s already exists in namespace %s => msg: %s", e.Key, e.Namespace, e.MSG)
}

//var KeyNotFoundError = errors.New("Key not found")
