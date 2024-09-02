package entitybuilder

import (
	"simple-microservice-backend/db/model"
	"simple-microservice-backend/pkg/request"
)

type ContactBuilder struct {
	contact model.Contact
}

func NewContactBuilder() *ContactBuilder {
	return &ContactBuilder{}
}

func (b *ContactBuilder) SetEmail(email string) *ContactBuilder {
	b.contact.Email = email
	return b
}

func (b *ContactBuilder) SetPhone(phoneNo uint) *ContactBuilder {
	b.contact.PhoneNo = phoneNo
	return b
}

func (b *ContactBuilder) SetLocation(location string) *ContactBuilder {
	b.contact.Location = location
	return b
}

func (b *ContactBuilder) SetAddr1(addr1 string) *ContactBuilder {
	b.contact.Addr1 = addr1
	return b
}

func (b *ContactBuilder) SetAddr2(addr2 string) *ContactBuilder {
	b.contact.Addr2 = addr2
	return b
}

func (b *ContactBuilder) SetAddr3(addr3 string) *ContactBuilder {
	b.contact.Addr3 = addr3
	return b
}

func (b *ContactBuilder) SetIsActive(isActive bool) *ContactBuilder {
	b.contact.IsActive = isActive
	return b
}

func (b *ContactBuilder) Build() model.Contact {
	return b.contact
}

type OwnerBuilder struct {
	owner model.Owner
}

func NewOwnerBuilder() *OwnerBuilder {
	return &OwnerBuilder{}
}

func (b *OwnerBuilder) SetName(name string) *OwnerBuilder {
	b.owner.Name = name
	return b
}

func (b *OwnerBuilder) SetCRNumber(crNumber uint) *OwnerBuilder {
	b.owner.CRNumber = crNumber
	return b
}

func (b *OwnerBuilder) SetContact(contact model.Contact) *OwnerBuilder {
	b.owner.Contact = contact
	return b
}

func (b *OwnerBuilder) Build() model.Owner {
	return b.owner
}

func CreateOwner(request request.OwnerCreate) *model.Owner {
	contact := NewContactBuilder().
		SetPhone(request.Contact.PhoneNo).
		SetLocation(request.Contact.Location).
		SetEmail(request.Contact.Email).
		SetAddr1(request.Contact.Addr1).
		SetAddr2(request.Contact.Addr2).
		SetAddr3(request.Contact.Addr3).
		SetIsActive(true).
		Build()

	owner := NewOwnerBuilder().
		SetCRNumber(request.CRNumber).
		SetContact(contact).
		SetName(request.Name).
		Build()

	return &owner
}
