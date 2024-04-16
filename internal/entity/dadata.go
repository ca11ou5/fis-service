package entity

type DadataRequest struct {
	Query string `json:"query"`
	Count int    `json:"count,omitempty"`
}

type DadataAddressResponse struct {
	Suggestions []Suggestion `json:"suggestions,omitempty"`
}

type Suggestion struct {
	Value             string `json:"value,omitempty"`
	UnrestrictedValue string `json:"unrestricted_value,omitempty"`
	Data              *Data  `json:"data,omitempty"`
}

type Data struct {
	PostalCode           string   `json:"postal_code,omitempty"`
	Country              string   `json:"country,omitempty"`
	CountryISOCode       string   `json:"country_iso_code,omitempty"`
	FederalDistrict      string   `json:"federal_district,omitempty"`
	RegionFiasID         string   `json:"region_fias_id,omitempty"`
	RegionKladrID        string   `json:"region_kladr_id,omitempty"`
	RegionISOCode        string   `json:"region_iso_code,omitempty"`
	RegionWithType       string   `json:"region_with_type,omitempty"`
	RegionType           string   `json:"region_type,omitempty"`
	RegionTypeFull       string   `json:"region_type_full,omitempty"`
	Region               string   `json:"region,omitempty"`
	AreaFiasID           string   `json:"area_fias_id,omitempty"`
	AreaKladrID          string   `json:"area_kladr_id,omitempty"`
	AreaWithType         string   `json:"area_with_type,omitempty"`
	AreaType             string   `json:"area_type,omitempty"`
	AreaTypeFull         string   `json:"area_type_full,omitempty"`
	Area                 string   `json:"area,omitempty"`
	CityFiasID           string   `json:"city_fias_id,omitempty"`
	CityKladrID          string   `json:"city_kladr_id,omitempty"`
	CityWithType         string   `json:"city_with_type,omitempty"`
	CityType             string   `json:"city_type,omitempty"`
	CityTypeFull         string   `json:"city_type_full,omitempty"`
	City                 string   `json:"city,omitempty"`
	CityArea             string   `json:"city_area,omitempty"`
	CityDistrictFiasID   string   `json:"city_district_fias_id,omitempty"`
	CityDistrictKladrID  string   `json:"city_district_kladr_id,omitempty"`
	CityDistrictWithType string   `json:"city_district_with_type,omitempty"`
	CityDistrictType     string   `json:"city_district_type,omitempty"`
	CityDistrictTypeFull string   `json:"city_district_type_full,omitempty"`
	CityDistrict         string   `json:"city_district,omitempty"`
	SettlementFiasID     string   `json:"settlement_fias_id,omitempty"`
	SettlementKladrID    string   `json:"settlement_kladr_id,omitempty"`
	SettlementWithType   string   `json:"settlement_with_type,omitempty"`
	SettlementType       string   `json:"settlement_type,omitempty"`
	SettlementTypeFull   string   `json:"settlement_type_full,omitempty"`
	Settlement           string   `json:"settlement,omitempty"`
	StreetFiasID         string   `json:"street_fias_id,omitempty"`
	StreetKladrID        string   `json:"street_kladr_id,omitempty"`
	StreetWithType       string   `json:"street_with_type,omitempty"`
	StreetType           string   `json:"street_type,omitempty"`
	StreetTypeFull       string   `json:"street_type_full,omitempty"`
	Street               string   `json:"street,omitempty"`
	SteadFiasID          string   `json:"stead_fias_id,omitempty"`
	SteadCadnum          string   `json:"stead_cadnum,omitempty"`
	SteadType            string   `json:"stead_type,omitempty"`
	SteadTypeFull        string   `json:"stead_type_full,omitempty"`
	Stead                string   `json:"stead,omitempty"`
	HouseFiasID          string   `json:"house_fias_id,omitempty"`
	HouseKladrID         string   `json:"house_kladr_id,omitempty"`
	HouseCadnum          string   `json:"house_cadnum,omitempty"`
	HouseType            string   `json:"house_type,omitempty"`
	HouseTypeFull        string   `json:"house_type_full,omitempty"`
	House                string   `json:"house,omitempty"`
	BlockType            string   `json:"block_type,omitempty"`
	BlockTypeFull        string   `json:"block_type_full,omitempty"`
	Block                string   `json:"block,omitempty"`
	Entrance             string   `json:"entrance,omitempty"`
	Floor                string   `json:"floor,omitempty"`
	FlatFiasID           string   `json:"flat_fias_id,omitempty"`
	FlatCadnum           string   `json:"flat_cadnum,omitempty"`
	FlatType             string   `json:"flat_type,omitempty"`
	FlatTypeFull         string   `json:"flat_type_full,omitempty"`
	Flat                 string   `json:"flat,omitempty"`
	FlatArea             string   `json:"flat_area,omitempty"`
	SquareMeterPrice     string   `json:"square_meter_price,omitempty"`
	FlatPrice            string   `json:"flat_price,omitempty"`
	RoomFiasID           string   `json:"room_fias_id,omitempty"`
	RoomCadnum           string   `json:"room_cadnum,omitempty"`
	RoomType             string   `json:"room_type,omitempty"`
	RoomTypeFull         string   `json:"room_type_full,omitempty"`
	Room                 string   `json:"room,omitempty"`
	PostalBox            string   `json:"postal_box,omitempty"`
	FiasID               string   `json:"fias_id,omitempty"`
	FiasCode             string   `json:"fias_code,omitempty"`
	FiasLevel            string   `json:"fias_level,omitempty"`
	FiasActualityState   string   `json:"fias_actuality_state,omitempty"`
	KladrID              string   `json:"kladr_id,omitempty"`
	GeonameID            string   `json:"geoname_id,omitempty"`
	CapitalMarker        string   `json:"capital_marker,omitempty"`
	Okato                string   `json:"okato,omitempty"`
	Oktmo                string   `json:"oktmo,omitempty"`
	TaxOffice            string   `json:"tax_office,omitempty"`
	TaxOfficeLegal       string   `json:"tax_office_legal,omitempty"`
	Timezone             string   `json:"timezone,omitempty"`
	GeoLat               string   `json:"geo_lat,omitempty"`
	GeoLon               string   `json:"geo_lon,omitempty"`
	BeltwayHit           string   `json:"beltway_hit,omitempty"`
	BeltwayDistance      string   `json:"beltway_distance,omitempty"`
	Metro                [3]Metro `json:"metro,omitempty"`
	Divisions            string   `json:"divisions,omitempty"`
	QcGeo                string   `json:"qc_geo,omitempty"`
	QcComplete           string   `json:"qc_complete,omitempty"`
	QcHouse              string   `json:"qc_house,omitempty"`
	HistoryValues        []string `json:"history_values,omitempty"`
	UnparsedParts        string   `json:"unparsed_parts,omitempty"`
	Source               string   `json:"source,omitempty"`
	Qc                   string   `json:"qc,omitempty"`
}

type Metro struct {
	Name     string `json:"name,omitempty"`
	Line     string `json:"line,omitempty"`
	Distance string `json:"distance,omitempty"`
}
