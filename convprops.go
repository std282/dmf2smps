package main

import "github.com/std282/dmf2smps/dmfns"

// ConvDetails is JSON-exportable collection of conversion details. It is used
// for tuning conversion algorithm
type ConvDetails struct {
	FileName   string `json:"file"`
	Properties `json:"params"`
	DACMap     []DACMapEntry `json:"dac"`
	PSGMap     []PSGMapEntry `json:"psg"`
}

// Properties desribe global song conversion details
type Properties struct {
	PreferFM6       bool `json:"preferFM6"`
	PreferPSG3      bool `json:"preferPSG3"`
	DecayVibrato    bool `json:"decayVibrato"`
	RestartAfterEnd bool `json:"restartAfterEnd"`
	ExtendedPSG     bool `json:"extendedPSG"`
}

// DACMapEntry describes single DAC sample mapping entry
type DACMapEntry struct {
	Note   string      `json:"note"`
	Bank   int         `json:"bank"`
	Name   string      `json:"name"`
	Sample interface{} `json:"dacSample"`
}

func (cd *ConvDetails) addDACEntry(note int16, bank int, name string) {
	cd.DACMap = append(cd.DACMap, DACMapEntry{
		Note:   noteName[note],
		Bank:   bank,
		Name:   name,
		Sample: nil,
	})
}

// PSGMapEntry decribes single PSG volume envelope mapping entry
type PSGMapEntry struct {
	InstNumber int         `json:"number"`
	Name       string      `json:"name"`
	Envelope   interface{} `json:"psgEnvelope"`
}

func (cd *ConvDetails) addPSGEntry(instNum int, instName string) {
	cd.PSGMap = append(cd.PSGMap, PSGMapEntry{
		Name:       instName,
		InstNumber: instNum,
		Envelope:   nil,
	})
}

var noteName = map[int16]string{
	dmfns.NoteC:  "C",
	dmfns.NoteCs: "C#",
	dmfns.NoteD:  "D",
	dmfns.NoteDs: "D#",
	dmfns.NoteE:  "E",
	dmfns.NoteF:  "F",
	dmfns.NoteFs: "F#",
	dmfns.NoteG:  "G",
	dmfns.NoteGs: "G#",
	dmfns.NoteA:  "A",
	dmfns.NoteAs: "A#",
	dmfns.NoteB:  "B",
}

var noteNameInv = map[string]int16{
	"C":  dmfns.NoteC,
	"C#": dmfns.NoteCs,
	"D":  dmfns.NoteD,
	"D#": dmfns.NoteDs,
	"E":  dmfns.NoteE,
	"F":  dmfns.NoteF,
	"F#": dmfns.NoteFs,
	"G":  dmfns.NoteG,
	"G#": dmfns.NoteGs,
	"A":  dmfns.NoteA,
	"A#": dmfns.NoteAs,
	"B":  dmfns.NoteB,
}