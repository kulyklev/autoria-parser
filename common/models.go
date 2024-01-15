package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CarsSearchResult struct {
	AdditionalParams struct {
		LangID      int    `json:"lang_id"`
		Page        string `json:"page"`
		ViewTypeID  int    `json:"view_type_id"`
		Target      string `json:"target"`
		Section     string `json:"section"`
		CatalogName string `json:"catalog_name"`
		Elastica    bool   `json:"elastica"`
		Nodejs      bool   `json:"nodejs"`
	} `json:"additional_params"`
	Result struct {
		SearchResult struct {
			Ids    []string `json:"ids"`
			Count  int      `json:"count"`
			LastID int      `json:"last_id"`
		} `json:"search_result"`
		SearchResultCommon struct {
			Data []struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Count  int `json:"count"`
			LastID int `json:"last_id"`
		} `json:"search_result_common"`
		IsCommonSearch bool `json:"isCommonSearch"`
		ActiveMarka    struct {
			LangID       int    `json:"lang_id"`
			MarkaID      int    `json:"marka_id"`
			Name         string `json:"name"`
			SetCat       string `json:"set_cat"`
			MainCategory int    `json:"main_category"`
			Active       int    `json:"active"`
			CountryID    int    `json:"country_id"`
			Eng          string `json:"eng"`
			Count        int    `json:"count"`
			Fit          int    `json:"fit"`
		} `json:"active_marka"`
		ActiveModel struct {
			LangID   int    `json:"lang_id"`
			ModelID  int    `json:"model_id"`
			MarkaID  int    `json:"marka_id"`
			Name     string `json:"name"`
			SetCat   string `json:"set_cat"`
			Active   int    `json:"active"`
			IsGroup  int    `json:"is_group"`
			ParentID int    `json:"parent_id"`
			Eng      string `json:"eng"`
			Count    int    `json:"count"`
			Fit      int    `json:"fit"`
		} `json:"active_model"`
		ActiveCity  interface{} `json:"active_city"`
		ActiveState interface{} `json:"active_state"`
		Revies      interface{} `json:"revies"`
		Additional  struct {
			UserAutoPositions []interface{} `json:"user_auto_positions"`
			SearchParams      struct {
				All struct {
					MarkaID          []int         `json:"marka_id"`
					ModelID          []int         `json:"model_id"`
					SYers            []int         `json:"s_yers"`
					PoYers           []int         `json:"po_yers"`
					Type             []interface{} `json:"type"`
					EngineVolumeFrom int           `json:"engineVolumeFrom"`
					Countpage        int           `json:"countpage"`
					WithPhoto        string        `json:"with_photo"`
					Page             string        `json:"page"`
					SearchType       int           `json:"searchType"`
					Target           string        `json:"target"`
					Event            string        `json:"event"`
					LangID           int           `json:"lang_id"`
					LimitPage        interface{}   `json:"limit_page"`
					LastID           int           `json:"last_id"`
					SaledParam       int           `json:"saledParam"`
					MmMarkaFiltr     []interface{} `json:"mm_marka_filtr"`
					MmModelFiltr     []interface{} `json:"mm_model_filtr"`
					UseOrigAutoTable bool          `json:"useOrigAutoTable"`
					WithoutStatus    bool          `json:"withoutStatus"`
					WithVideo        bool          `json:"with_video"`
					UnderCredit      int           `json:"under_credit"`
					ConfiscatedCar   int           `json:"confiscated_car"`
					ExchangeFilter   []interface{} `json:"exchange_filter"`
					OldOnly          bool          `json:"old_only"`
					AutoOptions      []interface{} `json:"auto_options"`
					//UserID                int           `json:"user_id"` // Here is problem
					PersonID              int           `json:"person_id"`
					WithDiscount          bool          `json:"with_discount"`
					AutoIDStr             string        `json:"auto_id_str"`
					BlackUserID           int           `json:"black_user_id"`
					OrderBy               int           `json:"order_by"`
					IsOnline              bool          `json:"is_online"`
					WithoutCache          bool          `json:"withoutCache"`
					WithLastID            bool          `json:"with_last_id"`
					Top                   int           `json:"top"`
					Currency              int           `json:"currency"`
					CurrencyID            int           `json:"currency_id"`
					CurrenciesArr         []interface{} `json:"currencies_arr"`
					PowerName             int           `json:"power_name"`
					PowerFrom             int           `json:"powerFrom"`
					PowerTo               int           `json:"powerTo"`
					HideBlackList         []interface{} `json:"hide_black_list"`
					Custom                int           `json:"custom"`
					Abroad                bool          `json:"abroad"`
					Damage                int           `json:"damage"`
					StarAuto              int           `json:"star_auto"`
					PriceOt               int           `json:"price_ot"`
					PriceDo               int           `json:"price_do"`
					Year                  int           `json:"year"`
					AutoIdsSearchPosition int           `json:"auto_ids_search_position"`
					PrintQs               int           `json:"print_qs"`
					IsHot                 int           `json:"is_hot"`
					DeletedAutoSearch     int           `json:"deletedAutoSearch"`
					CanBeChecked          int           `json:"can_be_checked"`
					ExcludeMM             []int         `json:"excludeMM"`
					GenerationID          []interface{} `json:"generation_id"`
					ModificationID        []interface{} `json:"modification_id"`
					FuelID                []int         `json:"fuel_id"`
					Verified              struct {
						Vin int `json:"VIN"`
					} `json:"verified"`
					AuctionPossible int `json:"auctionPossible"`
				} `json:"all"`
				Cleaned struct {
					MarkaID          []int         `json:"marka_id"`
					ModelID          []int         `json:"model_id"`
					SYers            []int         `json:"s_yers"`
					PoYers           []int         `json:"po_yers"`
					Type             []interface{} `json:"type"`
					EngineVolumeFrom int           `json:"engineVolumeFrom"`
					Countpage        int           `json:"countpage"`
					WithPhoto        string        `json:"with_photo"`
					SearchType       int           `json:"searchType"`
					Target           string        `json:"target"`
					Event            string        `json:"event"`
					LangID           int           `json:"lang_id"`
					MmMarkaFiltr     []interface{} `json:"mm_marka_filtr"`
					MmModelFiltr     []interface{} `json:"mm_model_filtr"`
					ExchangeFilter   []interface{} `json:"exchange_filter"`
					AutoOptions      []interface{} `json:"auto_options"`
					Currency         int           `json:"currency"`
					CurrenciesArr    []interface{} `json:"currencies_arr"`
					HideBlackList    []interface{} `json:"hide_black_list"`
					ExcludeMM        []int         `json:"excludeMM"`
					GenerationID     []interface{} `json:"generation_id"`
					ModificationID   []interface{} `json:"modification_id"`
					FuelID           []int         `json:"fuel_id"`
					Verified         struct {
						Vin int `json:"VIN"`
					} `json:"verified"`
				} `json:"cleaned"`
			} `json:"search_params"`
			QueryString string `json:"query_string"`
		} `json:"additional"`
	} `json:"result"`
}

type CarAd struct {
	Id     string `json:"id"`
	UserID int    `json:"userId"`
	/*UserBlocked []struct {
		ProjectIds struct {
			Num2 string `json:"2"`
		} `json:"project_ids,omitempty"`
		Reason []interface{} `json:"reason,omitempty"`
	} `json:"userBlocked,omitempty"`*/
	ChipsCount        int    `json:"chipsCount"`
	LocationCityName  string `json:"locationCityName"`
	CityLocative      string `json:"cityLocative"`
	AuctionPossible   bool   `json:"auctionPossible"`
	ExchangePossible  bool   `json:"exchangePossible"`
	RealtyExchange    bool   `json:"realtyExchange"`
	ExchangeType      string `json:"exchangeType"`
	ExchangeTypeID    int    `json:"exchangeTypeId"`
	AddDate           string `json:"addDate"`
	UpdateDate        string `json:"updateDate"`
	ExpireDate        string `json:"expireDate"`
	SoldDate          string `json:"soldDate"`
	UserHideADSStatus bool   `json:"userHideADSStatus"`
	UserPhoneData     struct {
		PhoneID string `json:"phoneId"`
		Phone   string `json:"phone"`
	} `json:"userPhoneData"`
	Uah                  int  `json:"UAH"`
	Usd                  int  `json:"USD"`
	Eur                  int  `json:"EUR"`
	IsAutoAddedByPartner bool `json:"isAutoAddedByPartner"`
	PartnerID            int  `json:"partnerId"`
	LevelData            struct {
		Level      int    `json:"level"`
		Label      int    `json:"label"`
		Period     int    `json:"period"`
		HotType    string `json:"hotType"`
		ExpireDate string `json:"expireDate"`
	} `json:"levelData"`
	Color struct {
		Name string `json:"name"`
		Eng  string `json:"eng"`
		Hex  string `json:"hex"`
	} `json:"color"`
	AutoData struct {
		Active             bool   `json:"active"`
		Vat                bool   `json:"vat"`
		Description        string `json:"description"`
		Version            string `json:"version"`
		GenerationID       int    `json:"generationId"`
		GenerationName     string `json:"generationName"`
		ModificationID     int    `json:"modificationId"`
		ModificationName   string `json:"modificationName"`
		EquipmentID        int    `json:"equipmentId"`
		EquipmentName      string `json:"equipmentName"`
		OnModeration       bool   `json:"onModeration"`
		Year               int    `json:"year"`
		AutoID             int    `json:"autoId"`
		BodyID             int    `json:"bodyId"`
		StatusID           int    `json:"statusId"`
		WithVideo          bool   `json:"withVideo"`
		Race               string `json:"race"`
		RaceInt            int    `json:"raceInt"`
		FuelID             int    `json:"fuelId"`
		FuelName           string `json:"fuelName"`
		FuelNameEng        string `json:"fuelNameEng"`
		GearBoxID          int    `json:"gearBoxId"`
		GearboxName        string `json:"gearboxName"`
		DriveID            int    `json:"driveId"`
		DriveName          string `json:"driveName"`
		IsSold             bool   `json:"isSold"`
		MainCurrency       string `json:"mainCurrency"`
		FromArchive        bool   `json:"fromArchive"`
		CategoryID         int    `json:"categoryId"`
		CategoryNameEng    string `json:"categoryNameEng"`
		SubCategoryNameEng string `json:"subCategoryNameEng"`
		Custom             int    `json:"custom"`
		WithVideoMessages  bool   `json:"withVideoMessages"`
	} `json:"autoData"`
	AutoInfoBar struct {
		Custom         bool `json:"custom"`
		Abroad         bool `json:"abroad"`
		OnRepairParts  bool `json:"onRepairParts"`
		Damage         bool `json:"damage"`
		UnderCredit    bool `json:"underCredit"`
		ConfiscatedCar bool `json:"confiscatedCar"`
	} `json:"autoInfoBar"`
	MarkName        string `json:"markName"`
	MarkNameEng     string `json:"markNameEng"`
	MarkID          int    `json:"markId"`
	ModelName       string `json:"modelName"`
	ModelNameEng    string `json:"modelNameEng"`
	ModelID         int    `json:"modelId"`
	SubCategoryName string `json:"subCategoryName"`
	PhotoData       struct {
		All       []int  `json:"all"`
		Count     int    `json:"count"`
		SeoLinkM  string `json:"seoLinkM"`
		SeoLinkSX string `json:"seoLinkSX"`
		SeoLinkB  string `json:"seoLinkB"`
		SeoLinkF  string `json:"seoLinkF"`
	} `json:"photoData"`
	LinkToView string `json:"linkToView"`
	Title      string `json:"title"`
	StateData  struct {
		Name          string `json:"name"`
		RegionName    string `json:"regionName"`
		RegionNameEng string `json:"regionNameEng"`
		LinkToCatalog string `json:"linkToCatalog"`
		Title         string `json:"title"`
		StateID       int    `json:"stateId"`
		CityID        int    `json:"cityId"`
	} `json:"stateData"`
	CanSetSpecificPhoneToAdvert bool          `json:"canSetSpecificPhoneToAdvert"`
	DontComment                 int           `json:"dontComment"`
	SendComments                int           `json:"sendComments"`
	Badges                      []interface{} `json:"badges"`
	Vin                         string        `json:"VIN"`
	VinSvg                      string        `json:"vinSvg"`
	CodedVin                    struct {
		Text string `json:"text"`
		Iv   string `json:"iv"`
	} `json:"codedVin"`
	HaveInfotechReport bool `json:"haveInfotechReport"`
	CheckedVin         struct {
		OrderID int    `json:"orderId"`
		Vin     string `json:"vin"`
		IsShow  bool   `json:"isShow"`
	} `json:"checkedVin"`
	InfotechReport struct {
		Vin string `json:"vin"`
	} `json:"infotechReport"`
	ModeratedAbroad  bool   `json:"moderatedAbroad"`
	SecureKey        string `json:"secureKey"`
	FirstTime        bool   `json:"firstTime"`
	TechnicalChecked bool   `json:"technicalChecked"`
	VideoMessageID   string `json:"videoMessageID"`
	Prices           []struct {
		Uah string `json:"UAH"`
		Usd string `json:"USD"`
		Eur string `json:"EUR"`
	} `json:"prices"`
	IsLeasing int `json:"isLeasing"`
	Dealer    struct {
		Link       string `json:"link"`
		Logo       string `json:"logo"`
		Type       string `json:"type"`
		ID         int    `json:"id"`
		Name       string `json:"name"`
		PackageID  int    `json:"packageId"`
		TypeID     int    `json:"typeId"`
		IsReliable bool   `json:"isReliable"`
		Verified   bool   `json:"verified"`
	} `json:"dealer"`
	WithInfoBar  bool          `json:"withInfoBar"`
	InfoBarText  string        `json:"infoBarText"`
	OptionStyles []interface{} `json:"optionStyles"`
}

type DbCar struct {
	Id           string
	UserId       int       `db:"user_id"`
	AutoId       int       `db:"auto_id"`
	Manufacturer string    `db:"manufacturer"`
	Model        string    `db:"model"`
	AddDate      time.Time `db:"add_date"`
	UpdateDate   time.Time `db:"update_date"`
	ExpireDate   time.Time `db:"expire_date"`
	SoldDate     time.Time `db:"sold_date"`
	Year         int       `db:"year"`
	BodyStyle    string    `db:"body_style"`
	FuelType     string    `db:"fuel_type"`
	GearboxType  string    `db:"gearbox_type"`
	Drive        string    `db:"drive"`
	MainCurrency string    `db:"main_currency"`
	Vin          string    `db:"vin"`
	VinSvg       string    `db:"vin_svg"`
	Url          string    `db:"url"`
	IsActive     bool      `db:"is_active"`
	ParsedAt     time.Time `db:"parsed_at"`
	Attitude     string    `db:"attitude"`
	Location     location  `db:"location"`
}

type location struct {
	Name          string
	RegionName    string
	RegionNameEng string
	LinkToCatalog string
	Title         string
	StateID       int
	CityID        int
}

// Value makes the "location" struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a location) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan makes the "location" struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a location) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

func ToDBCar(carAd CarAd) DbCar {
	addDateParse, err := time.Parse("2006-01-02 15:04:05", carAd.AddDate)
	if err != nil {
		fmt.Printf("addTime: %s\n", carAd.AddDate)
		fmt.Printf("error: %s\n", err)
	}

	updateDateParse, err := time.Parse("2006-01-02 15:04:05", carAd.UpdateDate)
	if err != nil {
		fmt.Printf("updateTime: %s\n", carAd.UpdateDate)
		fmt.Printf("error: %s\n", err)
	}

	expireDate, err := time.Parse("2006-01-02 15:04:05", carAd.ExpireDate)
	if err != nil {
		fmt.Printf("expireTime: %s\n", carAd.ExpireDate)
		fmt.Printf("error: %s\n", err)
	}

	var soldDate time.Time
	if carAd.SoldDate != "" {
		soldDate, err = time.Parse("2006-01-02 15:04:05", carAd.SoldDate)
		if err != nil {
			fmt.Printf("soldTime: %s\n", carAd.SoldDate)
			fmt.Printf("error: %s\n", err)
		}
	}

	return DbCar{
		Id:           uuid.NewString(),
		UserId:       carAd.UserID,
		AutoId:       carAd.AutoData.AutoID,
		Manufacturer: carAd.MarkName,
		Model:        carAd.ModelName,
		AddDate:      addDateParse,
		UpdateDate:   updateDateParse,
		ExpireDate:   expireDate,
		SoldDate:     soldDate,
		Year:         carAd.AutoData.Year,
		BodyStyle:    parseBodyStyle(carAd.AutoData.BodyID),
		FuelType:     parseFuelType(carAd.AutoData.FuelID),
		GearboxType:  parseGearboxType(carAd.AutoData.GearBoxID),
		Drive:        parseDrive(carAd.AutoData.DriveID),
		MainCurrency: carAd.AutoData.MainCurrency,
		Vin:          carAd.Vin,
		VinSvg:       carAd.VinSvg,
		Url:          "https://auto.ria.com" + carAd.LinkToView,
		ParsedAt:     time.Now(),
		Location: location{
			Name:          carAd.StateData.Name,
			RegionName:    carAd.StateData.RegionName,
			RegionNameEng: carAd.StateData.RegionNameEng,
			LinkToCatalog: carAd.StateData.LinkToCatalog,
			Title:         carAd.StateData.Title,
			StateID:       carAd.StateData.StateID,
			CityID:        carAd.StateData.CityID,
		},
	}
}

func parseBodyStyle(bodyId int) string {
	switch bodyId {
	case 3:
		return "Sedan"
	case 5:
		return "Crossover"
	case 8:
		return "Minivan"
	case 449:
		return "Microvan"
	case 4:
		return "Hatchback"
	case 2:
		return "Universal"
	case 6:
		return "Coup"
	case 7:
		return "Cabrio"
	case 9:
		return "Pickup"
	case 307:
		return "Liftback"
	case 448:
		return "Fastback"
	default:
		return ""
	}
}

func parseFuelType(fuelID int) string {
	switch fuelID {
	case 1:
		return "Petrol"
	case 2:
		return "Diesel"
	case 3:
		return "LPG"
	case 4:
		return "Petrol/LPG"
	case 5:
		return "Hybrid"
	case 6:
		return "Electro"
	case 7:
		return "Other"
	case 8:
		return "Metan"
	case 9:
		return "Propan"
	default:
		return ""
	}
}

func parseGearboxType(gearboxId int) string {
	switch gearboxId {
	case 1:
		return "Manual"
	case 2:
		return "Auto"
	case 3:
		return "Tiptronic"
	case 4:
		return "Robot"
	case 5:
		return "Variator"
	default:
		return ""
	}
}

func parseDrive(driveId int) string {
	switch driveId {
	case 1:
		return "AWD"
	case 2:
		return "FWD"
	case 3:
		return "RWD"
	default:
		return ""
	}
}

type dbPrice struct {
	Id       string
	CarId    string    `db:"car_id"`
	Uah      int       `db:"uah"`
	Usd      int       `db:"usd"`
	Eur      int       `db:"eur"`
	ParsedAt time.Time `db:"parsed_at"`
}

func ToDBPrice(carAd CarAd) dbPrice {
	return dbPrice{
		Id:       uuid.NewString(),
		Uah:      carAd.Uah,
		Usd:      carAd.Usd,
		Eur:      carAd.Eur,
		ParsedAt: time.Now(),
	}
}

type dbError struct {
	Id        string    `db:"id"`
	AutoId    string    `db:"auto_id"`
	Err       string    `db:"err"`
	CreatedAt time.Time `db:"created_at"`
}

func ToDBError(grErr ErrorFromGoroutine) dbError {
	return dbError{
		Id:        uuid.NewString(),
		AutoId:    grErr.FailedAutoId,
		Err:       grErr.Err.Error(),
		CreatedAt: time.Now(),
	}
}
