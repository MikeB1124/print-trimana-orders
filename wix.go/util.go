package wix

import (
	"time"
)

type Order struct {
	ID                   string        `json:"id"`
	RestaurantID         string        `json:"restaurantId"`
	LocationID           string        `json:"locationId"`
	RestaurantLocationID string        `json:"restaurantLocationId"`
	CreatedDate          time.Time     `json:"createdDate"`
	UpdatedDate          time.Time     `json:"updatedDate"`
	Comment              string        `json:"comment"`
	Currency             string        `json:"currency"`
	Status               string        `json:"status"`
	LineItems            []LineItem    `json:"lineItems"`
	Discounts            []interface{} `json:"discounts"`
	Payments             []Payment     `json:"payments"`
	Fulfillment          Fulfillment   `json:"fulfillment"`
	Customer             Customer      `json:"customer"`
	Totals               Totals        `json:"totals"`
	Activities           []Activity    `json:"activities"`
	ChannelInfo          ChannelInfo   `json:"channelInfo"`
}

type LineItem struct {
	Quantity         int              `json:"quantity"`
	Price            string           `json:"price"`
	Comment          string           `json:"comment"`
	DishOptions      []DishOption     `json:"dishOptions"`
	CatalogReference CatalogReference `json:"catalogReference"`
}

type DishOption struct {
	Name             string            `json:"name"`
	SelectedChoices  []SelectedChoice  `json:"selectedChoices"`
	MinChoices       int               `json:"minChoices"`
	MaxChoices       int               `json:"maxChoices"`
	Type             string            `json:"type"`
	AvailableChoices []AvailableChoice `json:"availableChoices"`
	DefaultChoices   []interface{}     `json:"defaultChoices"`
}

type SelectedChoice struct {
	Quantity         int              `json:"quantity"`
	Price            string           `json:"price"`
	DishOptions      []interface{}    `json:"dishOptions"`
	CatalogReference CatalogReference `json:"catalogReference"`
}

type AvailableChoice struct {
	ItemID string `json:"itemId"`
	Price  string `json:"price"`
	Name   string `json:"name"`
}

type CatalogReference struct {
	CatalogItemID          string `json:"catalogItemId"`
	CatalogItemName        string `json:"catalogItemName"`
	CatalogItemDescription string `json:"catalogItemDescription"`
	CatalogItemMedia       string `json:"catalogItemMedia"`
}

type Payment struct {
	Type                  string `json:"type"`
	Amount                string `json:"amount"`
	Method                string `json:"method"`
	ProviderTransactionID string `json:"providerTransactionId"`
}

type Fulfillment struct {
	Type            string          `json:"type"`
	PromisedTime    string          `json:"promisedTime"`
	ASAP            bool            `json:"asap"`
	PreparationTime int             `json:"preparationTime"`
	PickupDetails   PickupDetails   `json:"pickupDetails"`
	DeliveryDetails DeliveryDetails `json:"deliveryDetails"`
}

type DeliveryDetails struct {
	Address Address `json:"address"`
}

type Address struct {
	Formatted    string      `json:"formatted"`
	Country      interface{} `json:"country"`
	City         string      `json:"city"`
	Street       string      `json:"street"`
	StreetNumber string      `json:"streetNumber"`
	Apt          string      `json:"apt"`
	Floor        string      `json:"floor"`
	Entrance     interface{} `json:"entrance"`
	ZipCode      string      `json:"zipCode"`
	SubDivision  string      `json:"subdivision"`
}

type PickupDetails struct {
	Fee string `json:"fee"`
}

type Customer struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	ContactID string `json:"contactId"`
}

type Totals struct {
	Subtotal    string        `json:"subtotal"`
	Total       string        `json:"total"`
	Tax         string        `json:"tax"`
	Quantity    int           `json:"quantity"`
	Tip         string        `json:"tip"`
	ServiceFees []interface{} `json:"serviceFees"`
}

type Activity struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

type ChannelInfo struct {
	Type string `json:"type"`
}

type WixOrdersResponse struct {
	Orders []Order `json:"orders"`
}

type WixOrderAcceptedResponse struct {
	Order Order `json:"order"`
}
