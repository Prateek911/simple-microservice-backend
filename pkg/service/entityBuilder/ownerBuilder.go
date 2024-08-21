package entitybuilder

import "simple-microservice-backend/db/model"

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

func (b *ContactBuilder) Build() model.Contact {
	return b.contact
}

type ContactablesBuilder struct {
	contactables model.Contactables
}

func NewContactablesBuilder() *ContactablesBuilder {
	return &ContactablesBuilder{}
}

func (b *ContactablesBuilder) SetContact(contact model.Contact) *ContactablesBuilder {
	b.contactables.Contact = contact
	return b
}

func (b *ContactablesBuilder) SetIsActive(isActive bool) *ContactablesBuilder {
	b.contactables.IsActive = isActive
	return b
}

func (b *ContactablesBuilder) Build() model.Contactables {
	return b.contactables
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

func (b *OwnerBuilder) SetContactable(contact model.Contactables) *OwnerBuilder {
	b.owner.Contactable = contact
	return b
}

func (b *OwnerBuilder) Build() model.Owner {
	return b.owner
}
