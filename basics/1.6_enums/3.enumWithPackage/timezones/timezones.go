package timezones

//Timezone represents all timezones in the USA
type Timezone int

const (

	//EST Eastern Standard time (-5)
	EST Timezone = -(5 + iota)

	//CST Central Standard time (-6)
	CST

	//MST Mountain Standard time (-7)
	MST

	//PST Pacific Standard time (-8)
	PST
)
