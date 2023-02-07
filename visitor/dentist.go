package visitor

import (
	"fmt"
	"time"
)

type Sex string

const (
	Female Sex = "female"
	Male   Sex = "male"
)

type Dentist interface {
	GetSex() Sex
	DoScaling(p Patient)
	DoToothFilling(p Patient)
}

type Drg struct {
}

func (d Drg) GetSex() Sex {
	return Female
}
func (d Drg) DoScaling(p Patient) {
	if p.CountKarangGigi() > 10 {
		d.ManualScaling()
	} else {
		d.AutomaticScaling()
	}
}

func (d Drg) ManualScaling() {
	fmt.Println("do manual Scaling")
}

func (d Drg) AutomaticScaling() {
	fmt.Println("do automatic Scaling")
}

func (d Drg) DoToothFilling(p Patient) {

}

type Specialist struct {
}

func (d Specialist) GetSex() Sex {
	return Male
}

func (d Specialist) DoScaling(p Patient) {
	fmt.Println("maaf tidak melayani scaling")
}

func (d Specialist) DoToothFilling(p Patient) {
	record := p.MedicalRecord()
	if time.Now().Sub(record.LastimeScaling) > 30*24*time.Hour {
		if p.Age() > 60 {
			fmt.Println("lakukan gigi palsu")
		} else {
			fmt.Println("Lakukan tambal gigi")
		}
	}
}

type Patient interface {
	Accept(d Dentist)
	CountKarangGigi() int
	Age() int
	Profile() Profile
	MedicalRecord() MedicalRecord
}

type Profile struct {
}

type MedicalRecord struct {
	LastimeScaling time.Time
}

type KarangGigiPatient struct {
	SexPreference Sex
}

func (kgp KarangGigiPatient) Accept(d Dentist) {
	if d.GetSex() == kgp.SexPreference {
		d.DoScaling(kgp)
	}
}

func (kgp KarangGigiPatient) CountKarangGigi() int {
	return 10
}

func (kgp KarangGigiPatient) Age() int {
	return 10
}
func (kgp KarangGigiPatient) Profile() Profile {
	return Profile{}
}
func (kgp KarangGigiPatient) MedicalRecord() MedicalRecord {
	return MedicalRecord{}
}

type GigiBerlubangPatient struct {
	SexPreference Sex
}

func (gbp GigiBerlubangPatient) Accept(d Dentist) {
	d.DoToothFilling(gbp)
}

func (kgp GigiBerlubangPatient) CountKarangGigi() int {
	return 5
}

func (kgp GigiBerlubangPatient) Age() int {
	return 10
}
func (kgp GigiBerlubangPatient) Profile() Profile {
	return Profile{}
}
func (kgp GigiBerlubangPatient) MedicalRecord() MedicalRecord {
	return MedicalRecord{}
}

func init() {
	spc := Specialist{}
	drg := Drg{}

	kgp := KarangGigiPatient{}
	gbp := GigiBerlubangPatient{}

	kgp.Accept(spc)
	gbp.Accept(drg)
}
