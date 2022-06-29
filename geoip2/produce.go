package geoip2

const defaultLang = "en"

func getNameByLang(n map[string]string, l string) string {
	if len(n) == 0 {
		return ""
	}
	if val, ok := n[l]; ok {
		return val
	}
	if val, ok := n[defaultLang]; ok {
		return val
	}
	return ""
}

func ProduceCity(c *City, lang string) map[string]interface{} {
	out := make(map[string]interface{})

	out["city"] = getNameByLang(c.City.Names, lang)
	out["continent"] = getNameByLang(c.Continent.Names, lang)
	out["country"] = getNameByLang(c.Country.Names, lang)
	out["country_iso"] = c.Country.IsoCode

	out["latitude"] = c.Location.Latitude
	out["longitude"] = c.Location.Longitude
	out["accuracyRadius"] = c.Location.AccuracyRadius
	out["time_zone"] = c.Location.TimeZone
	out["metro_code"] = c.Location.MetroCode

	out["postCode"] = c.Postal.Code

	out["registeredCountry"] = getNameByLang(c.RegisteredCountry.Names, lang)
	out["registeredCountry_iso"] = c.RegisteredCountry.IsoCode

	out["representedCountry_iso"] = getNameByLang(c.RepresentedCountry.Names, lang)
	out["representedCountry_iso"] = c.RepresentedCountry.IsoCode
	out["representedCountry_type"] = c.RepresentedCountry.Type

	subdivisions := make([]map[string]string, len(c.Subdivisions))
	for n, sub := range c.Subdivisions {
		m := make(map[string]string)
		m["name"] = getNameByLang(sub.Names, lang)
		m["iso"] = sub.IsoCode
		subdivisions[n] = m
	}
	out["subdivisions"] = subdivisions

	out["is_anonymous_proxy"] = c.Traits.IsAnonymousProxy
	out["is_satellite_provider"] = c.Traits.IsSatelliteProvider

	return out
}
