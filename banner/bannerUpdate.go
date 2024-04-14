package banner

// аналогично исходному баннеру, но в виде интерфейсов
type BannerUpdater struct {
	Data      interface{} // json
	FeatureID interface{} // номер фичи
	TagsIDs   interface{} // номера тегов
	IsActive  interface{} // активность банера
}

// можно конкретизировать ошибки для каждого типа данных
func (b *Banner) UpdateBanner(bu BannerUpdater) error {
	d, ok := bu.Data.([]byte)
	if !ok {
		return ErrIncorrectData
	}
	b.Data = d

	fID, ok := bu.FeatureID.(int)
	if !ok {
		return ErrIncorrectData
	}
	b.FeatureID = fID

	tID, ok := bu.TagsIDs.([]int)
	if !ok {
		return ErrIncorrectData
	}
	b.TagsIDs = tID

	iA, ok := bu.IsActive.(bool)
	if !ok {
		return ErrIncorrectData
	}
	b.IsActive = iA

	return nil
}
