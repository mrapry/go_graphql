package helper

// PlanktonVariant data structure
type PlanktonVariant struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes Attr   `json:"attributes"`
	} `json:"data"`
	Included []struct {
		OffersProducts
	} `json:"included"`
}

// Attr data structure
type Attr struct {
	Name          string    `json:"name"`
	Type          []string  `json:"type"`
	SKU           string    `json:"skuNo"`
	RegisterNo    *string   `json:"registerNo"`
	Status        string    `json:"status"`
	Note          *string   `json:"note"`
	Cogs          int       `json:"cogs"`
	VideoURL      *[]string `json:"videoUrl"`
	AttachmentURL *[]string `json:"attachmentUrl"`
	Image         []Images  `json:"images"`
	Stock         []Stocks  `json:"stock"`
	Creator       Creators  `json:"creator"`
	CreatedAt     string    `json:"createdAt"`
	Editor        Editors   `json:"editor"`
	ModifiedAt    string    `json:"modifiedAt"`
}

// Images data structure
type Images struct {
	Default bool        `json:"default"`
	Order   int         `json:"order"`
	Variety []Varieties `json:"variety"`
}

// Varieties data structure
type Varieties struct {
	Type string `json:"type"`
	Size string `json:"size"`
	URL  string `json:"url"`
}

// Stocks data structure
type Stocks struct {
	LocationCode string `json:"locationCode"`
	Name         string `json:"name"`
	OnHand       int    `json:"onHand"`
	OnReserve    int    `json:"onReserve"`
	Available    int    `json:"available"`
}

// Creators data structure
type Creators struct {
	Email    string `json:"email"`
	ID       string `json:"id"`
	IP       string `json:"ip"`
	JobTitle string `json:"jobTitle"`
	Name     string `json:"name"`
}

// Editors data structure
type Editors struct {
	Email    string `json:"email"`
	ID       string `json:"id"`
	IP       string `json:"ip"`
	JobTitle string `json:"jobTitle"`
	Name     string `json:"name"`
}

// OffersProducts data structure
type OffersProducts struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes Attrs  `json:"attributes"`
}

// Attrs data structure
type Attrs struct {
	SkuNo                 string               `json:"skuNo,omitempty"`
	VendorSku             *string              `json:"vendorSku,omitempty"`
	MerchantID            *string              `json:"merchantId,omitempty"`
	ConditionID           int                  `json:"conditionId,omitempty"`
	ConditionName         string               `json:"conditionName,omitempty"`
	ConditionNote         *string              `json:"conditionNote,omitempty"`
	ShippingLength        *string              `json:"shippingLength,omitempty"`
	ShippingWidth         *string              `json:"shippingWidth,omitempty"`
	ShippingHeight        *string              `json:"shippingHeight,omitempty"`
	ShippingWeight        float32              `json:"shippingWeight,omitempty"`
	VplPrice              int                  `json:"vplPrice,omitempty"`
	VplSuggestedPrice     int                  `json:"vplSuggestedPrice,omitempty"`
	NormalPrice           int                  `json:"normalPrice,omitempty"`
	SpecialPrice          int                  `json:"specialPrice,omitempty"`
	SpecialPriceStartDate string               `json:"specialPriceStartDate,omitempty"`
	SpecialPriceEndDate   string               `json:"specialPriceEndDate,omitempty"`
	WarrantyTypeName      *string              `json:"warrantyTypeName,omitempty"`
	WarrantyPeriodID      int                  `json:"warrantyPeriodId,omitempty"`
	WarrantyPeriodName    *string              `json:"warrantyPeriodName,omitempty"`
	WarrantyNote          *string              `json:"warrantyNote,omitempty"`
	OfferInfo             *string              `json:"offerInfo,omitempty"`
	OfferStatus           string               `json:"offerStatus,omitempty"`
	HandlingTime          int                  `json:"handlingTime,omitempty"`
	ShippingNote          string               `json:"shippingNote,omitempty"`
	Name                  string               `json:"name,omitempty"`
	Model                 string               `json:"model,omitempty"`
	Status                string               `json:"status,omitempty"`
	RegisterNoTypeID      *int                 `json:"registerNoTypeId,omitempty"`
	RegisterNoTypeName    *string              `json:"registerNoTypeName,omitempty"`
	ExemptionTypeID       int                  `json:"exemptionTypeId,omitempty"`
	ExemptionTypeName     *string              `json:"exemptionTypeName,omitempty"`
	CategoryID            string               `json:"categoryId,omitempty"`
	CategoryName          *string              `json:"categoryName,omitempty"`
	CategoryStructure     []CategoryStructures `json:"categoryStructure,omitempty"`
	BrandID               string               `json:"brandId,omitempty"`
	BrandName             string               `json:"brandName,omitempty"`
	Description           []string             `json:"description,omitempty"`
	InTheBox              []string             `json:"inTheBox,omitempty"`
	KeyFeatures           *string              `json:"keyFeatures,omitempty"`
	SearchTerms           *string              `json:"searchTerms,omitempty"`
	IntendedUse           *string              `json:"intendedUse,omitempty"`
	TargetAudience        *string              `json:"targetAudience,omitempty"`
	RealLength            *string              `json:"realLength,omitempty"`
	RealWidth             *string              `json:"realWidth,omitempty"`
	RealHeight            *string              `json:"realHeight,omitempty"`
	RealWeight            float64              `json:"realWeight,omitempty"`
	Specs                 []Specs              `json:"specs,omitempty"`
	VariantTypeID         int                  `json:"variantTypeId,omitempty"`
	VariantTypeName       *string              `json:"variantTypeName,omitempty"`
	VariantType           *string              `json:"variantType,omitempty"`
	RichContent           string               `json:"richContent,omitempty"`
	Creator               Creators             `json:"creator"`
	CreatedAt             string               `json:"createdAt"`
	Editor                Editors              `json:"editor"`
	ModifiedAt            string               `json:"modifiedAt"`
}

// CategoryStructures data structure
type CategoryStructures struct {
	ID           string `json:"id"`
	OldID        string `json:"oldId"`
	Name         string `json:"name"`
	CurrentLevel int    `json:"currentLevel"`
}

// Specs data structure
type Specs struct {
	ID    string `json:"id"`
	OldID int    `json:"oldId"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
