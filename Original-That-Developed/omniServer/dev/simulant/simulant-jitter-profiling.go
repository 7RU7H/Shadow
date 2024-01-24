package main

import (
	"math/rand"
)

// clicksPerAge

// age
// 18 - 65 ( 1 -  )
// bigger the value the slower
// Beats per minute idea for click 60 is still quiet fast, but that also needs to be a seeded value between
// 18 -> (40 .. 80)
// Lower bound needs to step up slower than the upper bound

// techlevel
// 1 - 3 (normal, literate, devOrIt)
// - hotkey, click and typing speed

// tool
// switch statement on tool
// is it regular brower usage or automated cli

// filetype & web service
// average click depth per service is probably designed to be around 3 - 5 (as adverts, UI, service choices)
// average text on a page prior to download

// picture - slow  visuals for the brains
// video - slow visuals for the brains
// .doc, pdf, etc
// backups - slow if browser as automated would be queued

type SimulantJitterProfile struct {
	age      int
	ageValue // use mathematical name
}

func JitterCyclicOrdering() error {

}

func simulantJitter() error {
	err := rand.Shuffle(len(VALUE), JitterCyclicOrdering())
}
