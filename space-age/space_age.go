package space

// Age converts seconds and a planet choice to the
// age of a person in years of the chosen planet
func Age(seconds float64, planet string) float64 {
	// Earth: orbital period 365.25 Earth days, or 31557600 seconds
	ageOnEarth := seconds / 31557600
	switch planet {
	case "Mercury":
		// Mercury: orbital period 0.2408467 Earth years
		age := ageOnEarth / 0.2408467
		return age
	case "Venus":
		// Venus: orbital period 0.61519726 Earth years
		age := ageOnEarth / 0.61519726
		return age
	case "Mars":
		// Mars: orbital period 1.8808158 Earth years
		age := ageOnEarth / 1.8808158
		return age
	case "Jupiter":
		// Jupiter: orbital period 11.862615 Earth years
		age := ageOnEarth / 11.862615
		return age
	case "Saturn":
		// Saturn: orbital period 29.447498 Earth years
		age := ageOnEarth / 29.447498
		return age
	case "Uranus":
		// Uranus: orbital period 84.016846 Earth years
		age := ageOnEarth / 84.016846
		return age
	case "Neptune":
		// Neptune: orbital period 164.79132 Earth years
		age := ageOnEarth / 164.79132
		return age
	default:
		return ageOnEarth
	}
}
