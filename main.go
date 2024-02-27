package main

import (
	"fmt"
)

func main() {
	fmt.Println("--------- Testing Mobil ------------")

	//Kanan
	// diketuk --> "knock knock"
	// dibuka  --> "beep beep"

	//Kiri
	// diketuk --> "beep beep"
	// dibuka  --> "knock knock"

	rodaDepan := Karet{}
	rodaBelakang := Besi{}

	// Membuat objek pintu dengan bunyi yang berbeda
	pintuKanan := Pintu{Diketuk: "Knock", Dibuka: "beep"}
	pintuKiri := Pintu{Diketuk: "beep", Dibuka: "Knock"}

	// Membuat objek mobil
	mobil := Mobil{
		RodaDepan:    rodaDepan,
		RodaBelakang: rodaBelakang,
		PintuKanan:   pintuKanan,
		PintuKiri:    pintuKiri,
	}

	mobil.RodaDepan = Kayu{}
	mobil.RodaBelakang = Karet{}

	fmt.Println("Roda Depan:", mobil.RodaDepan.GetType())
	fmt.Println("Roda Belakang:", mobil.RodaBelakang.GetType())

	fmt.Println("Ketuk Pintu Kanan:", mobil.PintuKanan.Diketuk)
	fmt.Println("Buka Pintu Kanan:", mobil.PintuKanan.Dibuka)

	fmt.Println("Ketuk Pintu Kiri:", mobil.PintuKiri.Diketuk)
	fmt.Println("Buka Pintu Kiri:", mobil.PintuKiri.Dibuka)
}

type Roda interface {
	GetType() string
}

type Karet struct{}

func (bk Karet) GetType() string {
	return "Ban Karet"
}

type Kayu struct{}

func (bk Kayu) GetType() string {
	return "Ban Kayu"
}

type Besi struct{}

func (bb Besi) GetType() string {
	return "Ban Besi"
}

type Pintu struct {
	Diketuk string
	Dibuka  string
}

type Mobil struct {
	RodaDepan    Roda
	RodaBelakang Roda
	PintuKanan   Pintu
	PintuKiri    Pintu
}
