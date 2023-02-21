I denne oppgaven skal jeg implementere applikasjonen funtemps, som tar inn data ved hjelp av flagg, og skriver ut svar til stdout

difference := math.Abs(a - b)

func withinTolerance(a, b, error float64) bool {
// Først sjekk om tallene er nøyaktig like
if a == b {
return true
}

difference := math.Abs(a - b)

// Siden vi skal dele med b, må vi sjekke om den er 0
// Hvis b er 0, returner avgjørelsen om d er mindre enn feilmarginen
// som vi aksepterer
if b == 0 {
return difference < error
}

// Tilslutt sjekk den relative differanse mot feilmargin
return (difference/math.Abs(b)) < error
}

for _, tc := range tests {
got := FarhenheitToCelsius(tc.input)
if !reflect.DeepEqual(tc.want, got) {
t.Errorf("expected: %v, got: %v", tc.want, got)
}
}

for _, tc := range tests {
got := FarhenheitToCelsius(tc.input)
if !withinTolerance(tc.want, got, 1e-12) {
t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
}
}
