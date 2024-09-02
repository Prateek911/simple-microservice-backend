package responseBuilder

import (
	"simple-microservice-backend/db/model"
	"simple-microservice-backend/pkg/response"
	"time"
)

type BaseResponseBuilder struct {
	response response.BaseResponse
}

func NewBaseResponseBuilder() *BaseResponseBuilder {
	return &BaseResponseBuilder{}
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
	if deletedAt != nil && !deletedAt.IsZero() {
		b.response.DeletedAt = deletedAt
	}
	return b
}

func (b *BaseResponseBuilder) Build() response.BaseResponse {
	return b.response
}

type ContactResponseBuilder struct {
	response    response.ContactResponse
	baseBuilder *BaseResponseBuilder
}

func NewContactResponseBuilder() *ContactResponseBuilder {
	return &ContactResponseBuilder{
		baseBuilder: NewBaseResponseBuilder(),
	}
}

func (b *ContactResponseBuilder) SetID(id uint) *ContactResponseBuilder {
	b.response.ID = id
	return b
}

func (b *ContactResponseBuilder) SetCreatedAt(createdAt time.Time) *ContactResponseBuilder {
	b.response.CreatedAt = createdAt
	return b
}

func (b *ContactResponseBuilder) SetUpdatedAt(updatedAt time.Time) *ContactResponseBuilder {
	b.response.UpdatedAt = updatedAt
	return b
}

func (b *ContactResponseBuilder) SetDeletedAt(deletedAt *time.Time) *ContactResponseBuilder {
	if deletedAt != nil && !deletedAt.IsZero() {
		b.response.DeletedAt = deletedAt
	}
	return b
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

func (b *ContactResponseBuilder) SetBaseResponse(base response.BaseResponse) *ContactResponseBuilder {
	b.response.BaseResponse = base
	return b
}

func (b *ContactResponseBuilder) SetIsActive(isActive bool) *ContactResponseBuilder {
	b.response.IsActive = isActive
	return b
}

func (b *ContactResponseBuilder) Build() response.ContactResponse {
	if b.response.BaseResponse.ID == 0 {
		b.response.BaseResponse = b.baseBuilder.Build()
	}
	return b.response
}

type OwnerResponseBuilder struct {
	response       response.OwnerResponse
	contactBuilder *ContactResponseBuilder
}

func NewOwnerResponseBuilder() *OwnerResponseBuilder {
	return &OwnerResponseBuilder{
		contactBuilder: NewContactResponseBuilder(),
	}
}

func (b *OwnerResponseBuilder) SetID(id uint) *OwnerResponseBuilder {
	b.response.ID = id
	return b
}

func (b *OwnerResponseBuilder) SetName(name string) *OwnerResponseBuilder {
	b.response.Name = name
	return b
}

func (b *OwnerResponseBuilder) SetCRNumber(crNumber uint) *OwnerResponseBuilder {
	b.response.CRNumber = crNumber
	return b
}

func (b *OwnerResponseBuilder) SetContact(contact response.ContactResponse) *OwnerResponseBuilder {
	b.response.Contact = contact
	return b
}

func (b *OwnerResponseBuilder) SetCreatedAt(createdAt time.Time) *OwnerResponseBuilder {
	b.response.CreatedAt = createdAt
	return b
}

func (b *OwnerResponseBuilder) SetUpdatedAt(updatedAt time.Time) *OwnerResponseBuilder {
	b.response.UpdatedAt = updatedAt
	return b
}

func (b *OwnerResponseBuilder) SetDeletedAt(deletedAt *time.Time) *OwnerResponseBuilder {
	if deletedAt != nil && !deletedAt.IsZero() {
		b.response.DeletedAt = deletedAt
	}
	return b
}

func (b *OwnerResponseBuilder) Build() response.OwnerResponse {
	return b.response
}

func BuildResponse(owner *model.Owner) response.OwnerResponse {
	contactResponse := NewContactResponseBuilder().
		SetID(owner.Contact.ID).
		SetCreatedAt(owner.Contact.CreatedAt).
		SetUpdatedAt(owner.Contact.UpdatedAt).
		SetDeletedAt(&owner.Contact.DeletedAt.Time).
		SetEmail(owner.Contact.Email).
		SetPhoneNo(owner.Contact.PhoneNo).
		SetLocation(owner.Contact.Location).
		SetAddr1(owner.Contact.Addr1).
		SetAddr2(owner.Contact.Addr2).
		SetAddr3(owner.Contact.Addr3).
		SetIsActive(owner.Contact.IsActive).
		Build()

	ownerResponse := NewOwnerResponseBuilder().
		SetID(owner.ID).
		SetCreatedAt(owner.CreatedAt).
		SetUpdatedAt(owner.UpdatedAt).
		SetDeletedAt(&owner.DeletedAt.Time).
		SetName(owner.Name).
		SetCRNumber(owner.CRNumber).
		SetContact(contactResponse).
		Build()

	return ownerResponse
}
