package opendata

type OpenData interface {
    GetURL()string
    GetPets()[]Pet
}
