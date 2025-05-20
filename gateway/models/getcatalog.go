package models

import "google.golang.org/protobuf/types/known/wrapperspb"

type CatalogQueryParams struct {
	AgeFrom    *int32  `schema:"ageFrom"`
	AgeTo      *int32  `schema:"ageTo"`
	Sex        *string `schema:"sex"`
	HeightFrom *int32  `schema:"heightFrom"`
	HeightTo   *int32  `schema:"heightTo"`
	WeightFrom *int32  `schema:"weightFrom"`
	WeightTo   *int32  `schema:"weightTo"`
	Page       int32   `schema:"page" validate:"required,min=1"`
	Limit      int32   `schema:"limit" validate:"required,min=1,max=100"`
}

type GetCatalogRequest struct {
	AgeFrom    *wrapperspb.Int32Value  `protobuf:"bytes,1,opt,name=age_from,json=ageFrom,proto3" json:"age_from,omitempty"`
	AgeTo      *wrapperspb.Int32Value  `protobuf:"bytes,2,opt,name=age_to,json=ageTo,proto3" json:"age_to,omitempty"`
	Sex        *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=sex,proto3" json:"sex,omitempty"`
	HeightFrom *wrapperspb.Int32Value  `protobuf:"bytes,4,opt,name=height_from,json=heightFrom,proto3" json:"height_from,omitempty"`
	HeightTo   *wrapperspb.Int32Value  `protobuf:"bytes,5,opt,name=height_to,json=heightTo,proto3" json:"height_to,omitempty"`
	WeightFrom *wrapperspb.Int32Value  `protobuf:"bytes,6,opt,name=weight_from,json=weightFrom,proto3" json:"weight_from,omitempty"`
	WeightTo   *wrapperspb.Int32Value  `protobuf:"bytes,7,opt,name=weight_to,json=weightTo,proto3" json:"weight_to,omitempty"`
	Page       int32                   `protobuf:"varint,8,opt,name=page,proto3" json:"page,omitempty"`
	Limit      int32                   `protobuf:"varint,9,opt,name=limit,proto3" json:"limit,omitempty"`
}
