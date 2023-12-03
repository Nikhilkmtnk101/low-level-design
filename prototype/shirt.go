package prototype

import "fmt"

type ItemInfoGetter interface {
	GetInfo() string
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

var whiteShirtCache = &Shirt{
	Color: White,
	Price: 15,
	SKU:   "",
}

var blackShirtCache = &Shirt{
	Color: Black,
	Price: 18,
	SKU:   "",
}

var blueShirtCache = &Shirt{
	Color: Blue,
	Price: 20,
	SKU:   "",
}

type Shirt struct {
	Color byte
	Price float64
	SKU   string
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt Info is -> color: %b, price: %f, SKU: %s", s.Color, s.Price, s.SKU)
}
