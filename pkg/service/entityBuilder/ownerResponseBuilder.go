package entitybuilder

import (
	"simple-microservice-backend/pkg/response"
	"time"
)

type BaseResponseBuilder struct {
	response response.ClientResponse
}

func (b *BaseResponseBuilder) SetID(id uint) *BaseResponseBuilder {
	b.response.ID = id
	return b
}

func (b *BaseResponseBuilder) SetCreatedAt(createdAt time.Time) *BaseResponseBuilder {
	b.response.CreatedAt = createdAt
	return b
}

func (b *BaseResponseBuilder) SetUpdatedAt(updatedAt time.Time) *BaseResponseBuilder {
	b.response.UpdatedAt = updatedAt
	return b
}

func (b *BaseResponseBuilder) SetDeletedAt(deletedAt *time.Time) *BaseResponseBuilder {
	b.response.DeletedAt = deletedAt
	return b
}

func (b *BaseResponseBuilder) Build() response.ClientResponse {
	return b.response
}

type ContactablesResponseBuilder struct {
	response    response.ContactablesResponse
	baseBuilder BaseResponseBuilder
}

func NewContactablesResponseBuilder() *ContactablesResponseBuilder {
	return &ContactablesResponseBuilder{
		baseBuilder: BaseResponseBuilder{},
	}
}

func (b *ContactablesResponseBuilder) SetContact(contact response.ContactResponse) *ContactablesResponseBuilder {
	b.response.Contact = contact
	return b
}

func (b *ContactablesResponseBuilder) SetIsActive(isActive bool) *ContactablesResponseBuilder {
	b.response.IsActive = isActive
	return b
}

func (b *ContactablesResponseBuilder) Build() response.ContactablesResponse {
	b.response.ClientResponse = b.baseBuilder.Build()
	return b.response
}

type ContactResponseBuilder struct {
	response    response.ContactResponse
	baseBuilder BaseResponseBuilder
}

func NewContactResponseBuilder() *ContactResponseBuilder {
	return &ContactResponseBuilder{
		baseBuilder: BaseResponseBuilder{},
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

func (b *ContactResponseBuilder) Build() response.ContactResponse {
	b.response.ClientResponse = b.baseBuilder.Build()
	return b.response
}

type OwnerResponseBuilder struct {
	response     response.OwnerResponse
	baseResponse BaseResponseBuilder
}

func NewOwnerResponseBuilder() *OwnerResponseBuilder {
	return &OwnerResponseBuilder{
		baseResponse: BaseResponseBuilder{},
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

func (b *OwnerResponseBuilder) Build() response.OwnerResponse {
	b.response.ClientResponse = b.baseResponse.Build()
	return b.response
}
