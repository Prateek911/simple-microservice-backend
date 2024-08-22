package entitybuilder

import (
	"simple-microservice-backend/db/model"
	"simple-microservice-backend/pkg/response"
	"time"
)

type ClientResponseBuilder struct {
	response response.ClientResponse
}

func NewClientResponseBuilder() *ClientResponseBuilder {
	return &ClientResponseBuilder{}
}

func (b *ClientResponseBuilder) SetID(id uint) *ClientResponseBuilder {
	b.response.ID = id
	return b
}

func (b *ClientResponseBuilder) SetCreatedAt(createdAt time.Time) *ClientResponseBuilder {
	b.response.CreatedAt = createdAt
	return b
}

func (b *ClientResponseBuilder) SetUpdatedAt(updatedAt time.Time) *ClientResponseBuilder {
	b.response.UpdatedAt = updatedAt
	return b
}

func (b *ClientResponseBuilder) SetDeletedAt(deletedAt *time.Time) *ClientResponseBuilder {
	b.response.DeletedAt = deletedAt
	return b
}

func (b *ClientResponseBuilder) Build() response.ClientResponse {
	return b.response
}

type ContactResponseBuilder struct {
	response    response.ContactResponse
	baseBuilder *ClientResponseBuilder
}

func NewContactResponseBuilder() *ContactResponseBuilder {
	return &ContactResponseBuilder{
		baseBuilder: NewClientResponseBuilder(),
	}
}

func (b *ContactResponseBuilder) SetEmail(email string) *ContactResponseBuilder {
	b.response.Email = email
	return b
}

func (b *ContactResponseBuilder) SetPhoneNo(phoneNo uint) *ContactResponseBuilder {
	b.response.PhoneNo = phoneNo
	return b
}

func (b *ContactResponseBuilder) SetLocation(location string) *ContactResponseBuilder {
	b.response.Location = location
	return b
}

func (b *ContactResponseBuilder) SetAddr1(addr1 string) *ContactResponseBuilder {
	b.response.Addr1 = addr1
	return b
}

func (b *ContactResponseBuilder) SetAddr2(addr2 string) *ContactResponseBuilder {
	b.response.Addr2 = addr2
	return b
}

func (b *ContactResponseBuilder) SetAddr3(addr3 string) *ContactResponseBuilder {
	b.response.Addr3 = addr3
	return b
}

func (b *ContactResponseBuilder) SetBaseResponse(base response.ClientResponse) *ContactResponseBuilder {
	b.response.ClientResponse = base
	return b
}

func (b *ContactResponseBuilder) Build() response.ContactResponse {
	if b.response.ClientResponse.ID == 0 {
		b.response.ClientResponse = b.baseBuilder.Build()
	}
	return b.response
}

type ContactablesResponseBuilder struct {
	response       response.ContactablesResponse
	baseBuilder    *ClientResponseBuilder
	contactBuilder *ContactResponseBuilder
}

func NewContactablesResponseBuilder() *ContactablesResponseBuilder {
	return &ContactablesResponseBuilder{
		baseBuilder:    NewClientResponseBuilder(),
		contactBuilder: NewContactResponseBuilder(),
	}
}

func (b *ContactablesResponseBuilder) SetIsActive(isActive bool) *ContactablesResponseBuilder {
	b.response.IsActive = isActive
	return b
}

func (b *ContactablesResponseBuilder) SetContact(contact response.ContactResponse) *ContactablesResponseBuilder {
	b.response.Contact = contact
	return b
}

func (b *ContactablesResponseBuilder) SetBaseResponse(base response.ClientResponse) *ContactablesResponseBuilder {
	b.response.ClientResponse = base
	return b
}

func (b *ContactablesResponseBuilder) Build() response.ContactablesResponse {
	if b.response.ClientResponse.ID == 0 {
		b.response.ClientResponse = b.baseBuilder.Build()
	}
	return b.response
}

// OwnerResponseBuilder builds the OwnerResponse
type OwnerResponseBuilder struct {
	response            response.OwnerResponse
	baseBuilder         *ClientResponseBuilder
	contactablesBuilder *ContactablesResponseBuilder
}

func NewOwnerResponseBuilder() *OwnerResponseBuilder {
	return &OwnerResponseBuilder{
		baseBuilder:         NewClientResponseBuilder(),
		contactablesBuilder: NewContactablesResponseBuilder(),
	}
}

func (b *OwnerResponseBuilder) SetName(name string) *OwnerResponseBuilder {
	b.response.Name = name
	return b
}

func (b *OwnerResponseBuilder) SetCRNumber(crNumber uint) *OwnerResponseBuilder {
	b.response.CRNumber = crNumber
	return b
}

func (b *OwnerResponseBuilder) SetContactable(contactable response.ContactablesResponse) *OwnerResponseBuilder {
	b.response.Contactable = contactable
	return b
}

func (b *OwnerResponseBuilder) SetBaseResponse(base response.ClientResponse) *OwnerResponseBuilder {
	b.response.ClientResponse = base
	return b
}

func (b *OwnerResponseBuilder) Build() response.OwnerResponse {
	if b.response.ClientResponse.ID == 0 {
		b.response.ClientResponse = b.baseBuilder.Build()
	}
	contactableResponse := b.contactablesBuilder.
		SetBaseResponse(b.response.ClientResponse).
		SetIsActive(b.response.Contactable.IsActive).
		SetContact(b.response.Contactable.Contact).
		Build()

	b.response.Contactable = contactableResponse

	return b.response
}

func BuildResponse(owner *model.Owner) response.OwnerResponse {
	ownerResponse := NewOwnerResponseBuilder().
		SetName(owner.Name).
		SetCRNumber(owner.CRNumber).
		SetBaseResponse(
			NewClientResponseBuilder().
				SetID(owner.ID).
				SetCreatedAt(owner.CreatedAt).
				SetUpdatedAt(owner.UpdatedAt).
				SetDeletedAt(&owner.DeletedAt.Time).
				Build(),
		).
		SetContactable(
			NewContactablesResponseBuilder().
				SetBaseResponse(
					NewClientResponseBuilder().
						SetID(owner.Contactable.ID).
						SetCreatedAt(owner.Contactable.CreatedAt).
						SetUpdatedAt(owner.Contactable.UpdatedAt).
						SetDeletedAt(&owner.Contactable.DeletedAt.Time).
						Build(),
				).
				SetIsActive(owner.Contactable.IsActive).
				SetContact(
					NewContactResponseBuilder().
						SetBaseResponse(
							NewClientResponseBuilder().
								SetID(owner.Contactable.Contact.ID).
								SetCreatedAt(owner.Contactable.Contact.CreatedAt).
								SetUpdatedAt(owner.Contactable.Contact.UpdatedAt).
								SetDeletedAt(&owner.Contactable.Contact.DeletedAt.Time).
								Build(),
						).
						SetEmail(owner.Contactable.Contact.Email).
						SetPhoneNo(owner.Contactable.Contact.PhoneNo).
						SetLocation(owner.Contactable.Contact.Location).
						SetAddr1(owner.Contactable.Contact.Addr1).
						SetAddr2(owner.Contactable.Contact.Addr2).
						SetAddr3(owner.Contactable.Contact.Addr3).
						Build(),
				).
				Build(),
		).
		Build()

	return ownerResponse
}
