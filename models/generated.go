// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type BookingInput struct {
	Token     int    `json:"token"`
	Volume    int    `json:"volume"`
	Booker    string `json:"booker"`
	Deposit   string `json:"deposit"`
	Delivered bool   `json:"delivered"`
}

type BookingStatusInput struct {
	ID        string `json:"id"`
	Delivered bool   `json:"delivered"`
}

type BookingsQueryInput struct {
	Token int `json:"token"`
}

type CancellationUpdateInput struct {
	BookingID    string `json:"bookingId"`
	SeasonNumber int    `json:"seasonNumber"`
	Token        int    `json:"token"`
	NewSupply    int    `json:"newSupply"`
	NewVolume    int    `json:"newVolume"`
	NewDeposit   string `json:"newDeposit"`
}

type FarmInput struct {
	ID        string `json:"id"`
	Size      string `json:"size"`
	Soil      string `json:"soil"`
	ImageHash string `json:"imageHash"`
	Season    string `json:"season"`
	Owner     string `json:"owner"`
}

type HarvestBookersInput struct {
	Token        int `json:"token"`
	SeasonNumber int `json:"seasonNumber"`
	NoOfBookers  int `json:"noOfBookers"`
}

type HarvestInput struct {
	SeasonNumber int    `json:"seasonNumber"`
	Token        int    `json:"token"`
	TotalSupply  int    `json:"totalSupply"`
	Price        string `json:"price"`
	SupplyUnit   string `json:"supplyUnit"`
}

type HarvestUpdateInput struct {
	Token        int `json:"token"`
	SeasonNumber int `json:"seasonNumber"`
	NewSupply    int `json:"newSupply"`
}

type PlantingInput struct {
	SeasonNumber  int    `json:"seasonNumber"`
	Token         int    `json:"token"`
	SeedUsed      string `json:"seedUsed"`
	ExpectedYield string `json:"expectedYield"`
	SeedSupplier  string `json:"seedSupplier"`
}

type PreparationInput struct {
	SeasonNumber int    `json:"seasonNumber"`
	Token        int    `json:"token"`
	Crop         string `json:"crop"`
	Fertilizer   string `json:"fertilizer"`
}

type ReceivershipUpdateInput struct {
	BookingID        string `json:"bookingId"`
	NewBookerVolume  int    `json:"newBookerVolume"`
	NewBookerDeposit string `json:"newBookerDeposit"`
	Delivered        bool   `json:"delivered"`
}

type SeasonUpdateInput struct {
	Token  int    `json:"token"`
	Season string `json:"season"`
}

type SeasonsQueryInput struct {
	Token int `json:"token"`
}
