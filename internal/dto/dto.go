package dto

type DomainTransferObject interface {
	Validate() error
}
