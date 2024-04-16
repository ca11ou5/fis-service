package entity

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"unicode/utf8"
)

var errUnmarshalFail = errors.New("failed to unmarshal into Address type")

// Application Entity
type Application struct {
	ID                          primitive.ObjectID `bson:"_id,omitempty"`
	Name                        string             `json:"name"`
	Surname                     string             `json:"surname"`
	Patronymic                  string             `json:"patronymic"`
	Phone                       string             `json:"phone"`
	BirthDate                   string             `json:"birthDate"`
	PassportSeries              string             `json:"passportSeries"`
	PassportNumber              string             `json:"passportNumber"`
	PassportDate                string             `json:"passportDate"`
	Sum                         int                `json:"sum"`
	UtmMedium                   string             `json:"utm_medium"`
	UtmSource                   string             `json:"utm_source"`
	UtmCampaign                 string             `json:"utm_campaign"`
	UtmContent                  string             `json:"utm_content"`
	UtmTerm                     string             `json:"utm_term"`
	IPAddressClient             string             `json:"ipAddressClient"`
	Term                        int                `json:"term"`
	Gender                      string             `json:"gender"`
	ValidPhone                  string             `json:"validPhone"`
	Agree                       bool               `json:"agree"`
	BirthPlace                  string             `json:"birthPlace"`
	BirthCountry                string             `json:"birthCountry"`
	Citizenship                 *string            `json:"citizenship"`
	PassportCode                string             `json:"passportCode"`
	PassportOrgan               string             `json:"passportOrgan"`
	RegistrationDate            string             `json:"registrationDate"`
	RegistrationAddress         string             `json:"registrationAddress"`
	LifeAddress                 string             `json:"lifeAddress"`
	LifeAddressDate             string             `json:"lifeAddressDate"`
	DifferentRegistration       bool               `json:"differentRegistration"`
	Salary                      int                `json:"salary"`
	OtherSalary                 int                `json:"otherSalary"`
	MaritalStatus               string             `json:"maritalStatus"`
	SpouseName                  string             `json:"spouseName"`
	SpouseSurname               string             `json:"spouseSurname"`
	SpousePatronymic            string             `json:"spousePatronymic"`
	SpouseBirthDate             string             `json:"spouseBirthDate"`
	SpousePhone                 string             `json:"spousePhone"`
	SocialStatus                string             `json:"socialStatus"`
	Education                   string             `json:"education"`
	AdditionalContactPhone      string             `json:"additionalContactPhone"`
	AdditionalContact           string             `json:"additionalContact"`
	AdditionalContactSurname    string             `json:"additionalContactSurname"`
	AdditionalContactName       string             `json:"additionalContactName"`
	AdditionalContactPatronymic string             `json:"additionalContactPatronymic"`
	AdditionalContactAgreement  bool               `json:"additionalContactAgreement"`
	HasOldPassport              bool               `json:"hasOldPassport"`
	OldPassportSeries           string             `json:"oldPassportSeries"`
	OldPassportNumber           string             `json:"oldPassportNumber"`
	OldPassportCode             string             `json:"oldPassportCode"`
	OldPassportDate             string             `json:"oldPassportDate"`
	OldPassportOrgan            string             `json:"oldPassportOrgan"`
	ChangedCredentials          bool               `json:"changedCredentials"`
	OldSurname                  string             `json:"oldSurname"`
	OldName                     string             `json:"oldName"`
	OldPatronymic               string             `json:"oldPatronymic"`
	LastNameChangeReason        string             `json:"lastNameChangeReason"`
	ChangeCredentialsYear       int                `json:"changeCredentialsYear"`
	LivingType                  string             `json:"livingType"`
	WorkName                    string             `json:"workName"`
	WorkAge                     string             `json:"workAge"`
	WorkDate                    string             `json:"workDate"`
	WorkEmployerIndustry        string             `json:"workEmployerIndustry"`
	WorkEmployerType            string             `json:"workEmployerType"`
	WorkPhone                   string             `json:"workPhone"`
	Experience                  string             `json:"experience"`
	EmployerType                string             `json:"employerType"`
	WorkPosition                string             `json:"workPosition"`
	WorkSector                  string             `json:"workSector"`
	WorkAddress                 string             `json:"workAddress"`
	HasCar                      bool               `json:"hasCar"`
	HasCarNumber                bool               `json:"hasCarNumber"`
	CarMark                     string             `json:"carMark"`
	CarModel                    string             `json:"carModel"`
	CarNumber                   string             `json:"carNumber"`
	CarYear                     string             `json:"carYear"`
	HasRealEstate               bool               `json:"hasRealEstate"`
	RealEstatePropertyType      string             `json:"realEstatePropertyType"`
	RealEstateType              string             `json:"realEstateType"`
	RealEstateRooms             string             `json:"realEstateRooms"`
	RealEstateWayToGet          string             `json:"realEstateWayToGet"`
	RealEstateAddress           string             `json:"realEstateAddress"`
	ProductName                 string             `json:"productName"`
	Email                       string             `json:"email"`
	Region                      string             `json:"region"`
	AdditionalData              AdditionalData     `json:"additionalData"`
	OpenApiSocialStatus         string             `json:"openApiSocialStatus" bson:"openApiSocialStatus"`
	OnlineDecision              string             `json:"onlineDecision" bson:"onlineDecision"`
	OrderID                     string             `json:"orderID" bson:"orderID"`
	// TODO: check OrderID type
	LastScoringStatus ScoringStatus `json:"lastScoringStatus" bson:"lastScoringStatus"`
	CreditScore       string        `json:"creditScore" bson:"creditScore"`
	PartnerName       string        `json:"partnerName" bson:"partnerName"`
}

type AdditionalData struct {
	QsParams      QsParams      `json:"qsParams"`
	Cookie        string        `json:"cookie"`
	URL           string        `json:"url"`
	KameleoonData KameleoonData `json:"kameleoonData"`
	TZ            string        `json:"tz"`
}

type QsParams struct{}
type KameleoonData struct{}

// TODO: Check the fields
type UserApplication struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Products []string           `json:"products" bson:"products"`
	Scope    []string           `json:"scope" bson:"scope"`
	Username string             `json:"username" bson:"username"`
	Token    string             `json:"token" bson:"token"`
}

type Address struct {
	StructValue *Suggestion
	StringValue string
}

func (a *Address) UnmarshalBSON(data []byte) error {
	var address Suggestion

	err := bson.Unmarshal(data, &address)
	if err == nil {
		a.StructValue = &address
		return nil
	}

	if utf8.Valid(data) {
		a.StringValue = string(data)
		return nil
	}

	return errUnmarshalFail
}

func (a *Address) MarshalBSON() ([]byte, error) {
	if a.StructValue != nil {
		data, err := bson.Marshal(a.StructValue)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	if a.StringValue != "" {
		data, err := bson.Marshal(a.StringValue)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	return []byte{}, nil
}
