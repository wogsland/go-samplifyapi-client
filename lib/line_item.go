package samplify

// Action ...
type Action string

// Action values for changing LineItem state
const (
	ActionLaunched Action = "launch"
	ActionPaused   Action = "pause"
	ActionClosed   Action = "close"
)

// FeasibilityStatus ...
type FeasibilityStatus string

// FeasibilityStatus values
const (
	FeasibilityStatusReady      FeasibilityStatus = "READY"
	FeasibilityStatusProcessing FeasibilityStatus = "PROCESSING"
)

//QuotaPlan ...
type QuotaPlan struct {
	Filters     []*QuotaFilters `json:"filters" valid:"required"`
	QuotaGroups []*QuotaGroup   `json:"quotaGroups" valid:"required"`
}

// QuotaFilters ...
type QuotaFilters struct {
	AttributeID string   `json:"attributeId"`
	Options     []string `json:"options"`
}

// QuotaGroup ...
type QuotaGroup struct {
	Name   string   `json:"name"`
	Quotas []*Quota `json:"quotas"`
}

// Quota ...
type Quota struct {
	AttributeID string         `json:"attributeId"`
	Options     []*QuotaOption `json:"options"`
}

// QuotaOption ...
type QuotaOption struct {
	Option []string `json:"option"`
	Perc   float64  `json:"perc"`
}

// EndLinks ...
type EndLinks struct {
	Complete  string `json:"complete"`
	Screenout string `json:"screenout"`
	OverQuota string `json:"overquota"`
}

// LineItemHeader ...
type LineItemHeader struct {
	Model
	ExtLineItemID string      `json:"extLineItemId"`
	State         State       `json:"state"`
	StateReason   string      `json:"stateReason"`
	LaunchedAt    *CustomTime `json:"launchedAt"`
}

// LineItem ...
type LineItem struct {
	LineItemHeader
	Title               string     `json:"title"`
	CountryISOCode      string     `json:"countryISOCode"`
	LanguageISOCode     string     `json:"languageISOCode"`
	SurveyURL           string     `json:"surveyURL"`
	SurveyTestURL       string     `json:"surveyTestURL"`
	IndicativeIncidence float64    `json:"indicativeIncidence"`
	DaysInField         int64      `json:"daysInField"`
	LengthOfInterview   int64      `json:"lengthOfInterview"`
	RequiredCompletes   int64      `json:"requiredCompletes"`
	QuotaPlan           *QuotaPlan `json:"quotaPlan"`
	EndLinks            *EndLinks  `json:"endLinks"`
}

// LineItemCriteria has the fields to create or update a LineItem
type LineItemCriteria struct {
	ExtLineItemID       string     `json:"extLineItemId" valid:"required"`
	Title               string     `json:"title" valid:"required"`
	CountryISOCode      string     `json:"countryISOCode" valid:"required,ISO3166Alpha2"`
	LanguageISOCode     string     `json:"languageISOCode" valid:"required,languageISOCode"`
	SurveyURL           string     `json:"surveyURL" valid:"optional,url"`
	SurveyTestURL       string     `json:"surveyTestURL" valid:"optional,url"`
	IndicativeIncidence float64    `json:"indicativeIncidence" valid:"required"`
	DaysInField         int64      `json:"daysInField" valid:"required"`
	LengthOfInterview   int64      `json:"lengthOfInterview" valid:"required"`
	RequiredCompletes   int64      `json:"requiredCompletes" valid:"required"`
	QuotaPlan           *QuotaPlan `json:"quotaPlan" valid:"required"`
}

// BuyProjectLineItem ...
type BuyProjectLineItem struct {
	ExtLineItemID string `json:"extLineItemId"`
	State         State  `json:"state"`
}

// LineItemReport ...
type LineItemReport struct {
	ExtLineItemID      string  `json:"extLineItemId"`
	State              State   `json:"state"`
	Attempts           int64   `json:"attempts"`
	Completes          int64   `json:"completes"`
	Overquotas         int64   `json:"overquotas"`
	Screenouts         int64   `json:"screenouts"`
	Starts             int64   `json:"starts"`
	Conversion         float64 `json:"conversion"`
	RemainingCompletes int64   `json:"remainingCompletes"`
	ActualMedianLOI    int64   `json:"actualMedianLOI"`
	IncurredCost       float64 `json:"incurredCost"`
	EstimatedCost      float64 `json:"estimatedCost"`
}

// Feasibility ...
type Feasibility struct {
	Status           FeasibilityStatus `json:"status"`
	CostPerInterview float64           `json:"costPerInterview"`
	Expiry           CustomTime        `json:"expiry"`
	Currency         string            `json:"currency"`
	Feasible         bool              `json:"feasible"`
}

// Attribute ... Supported attribute for a country and language. Required to build up the Quota Plan
type Attribute struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Text    string             `json:"text"`
	Type    string             `json:"type"`
	Options []*AttributeOption `json:"options"`
}

// AttributeOption ...
type AttributeOption struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
