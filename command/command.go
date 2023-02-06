package command

// waiter -> command -> receiver

type Interface interface {
	Execute()
}

type FriedRiceCommand struct {
	Chef  Chef
	Param FriedRiceParams
}

func (frc FriedRiceCommand) Execute() {
	frc.Chef.CookFriedRice(frc.Param)
}

type SatePadangCommand struct {
	Chef  Chef
	Param SatePadangParams
}

func (spc SatePadangCommand) Execute() {
	spc.Chef.PrepareSatePadang(spc.Param)
}

func init() {
	lordAdi := &LordAdi{}
	w1 := &Waiter{}
	fc := FriedRiceCommand{
		Chef: lordAdi,
		Param: FriedRiceParams{
			Spiciness: 1,
			Saltiness: 0,
			ExtraEgg:  true,
		},
	}
	w1.AddOrder(fc)

	sp := SatePadangCommand{
		Chef:  lordAdi,
		Param: SatePadangParams{},
	}
	w1.AddOrder(sp)

	w1.Finalize()
}
