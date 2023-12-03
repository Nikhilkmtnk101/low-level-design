package prototype

import "errors"

type ShirtCloner interface {
	GetClone(color byte) (ItemInfoGetter, error)
}

type ShirtCacheCloner struct{}

func GetShirtCloner() ShirtCloner {
	return &ShirtCacheCloner{}
}

func (s *ShirtCacheCloner) GetClone(color byte) (ItemInfoGetter, error) {
	switch color {
	case White:
		res := *whiteShirtCache
		return &res, nil
	case Black:
		res := *blackShirtCache
		return &res, nil
	case Blue:
		res := *blueShirtCache
		return &res, nil
	}
	return nil, errors.New("not valid color of shirt")
}
