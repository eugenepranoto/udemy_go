package command

import "fmt"

type FriedRiceParams struct {
	Spiciness int
	Saltiness int
	ExtraEgg  bool
}

type SatePadangParams struct {
}

type Chef interface {
	CookFriedRice(param FriedRiceParams)
	PrepareSatePadang(param SatePadangParams)
}

type LordAdi struct {
}

func (la LordAdi) CookFriedRice(params FriedRiceParams) {
	fmt.Println("Lord adi cook rice %w", params)
}

func (la LordAdi) PrepareSatePadang(params SatePadangParams) {
	fmt.Println("Lord adi cook sate padang %w", params)
}
