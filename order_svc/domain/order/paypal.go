package order

import "time"

type PurchaseUnit struct {
	ReferenceID    string `json:"reference_id"`
	Amount         Amount `json:"amount"`
	Payee          Payee  `json:"payee"`
	SoftDescriptor string `json:"soft_descriptor"`
	Shipping       struct {
		Name    Name    `json:"name"`
		Address Address `json:"address"`
	} `json:"shipping"`
	Payments Payments `json:"payments"`
}

type Amount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

type Payee struct {
	EmailAddress string `json:"email_address"`
	MerchantID   string `json:"merchant_id"`
}

type Name struct {
	FullName string `json:"full_name"`
}

type Address struct {
	AddressLine1 string `json:"address_line_1"`
	AdminArea2   string `json:"admin_area_2"`
	AdminArea1   string `json:"admin_area_1"`
	PostalCode   string `json:"postal_code"`
	CountryCode  string `json:"country_code"`
}

type Capture struct {
	ID               string           `json:"id"`
	Status           string           `json:"status"`
	Amount           Amount           `json:"amount"`
	FinalCapture     bool             `json:"final_capture"`
	SellerProtection SellerProtection `json:"seller_protection"`
	CreateTime       time.Time        `json:"create_time"`
	UpdateTime       time.Time        `json:"update_time"`
}

type Payments struct {
	Captures []Capture `json:"captures"`
}

type SellerProtection struct {
	Status            string   `json:"status"`
	DisputeCategories []string `json:"dispute_categories"`
}

type Payer struct {
	Name         Name    `json:"name"`
	EmailAddress string  `json:"email_address"`
	PayerID      string  `json:"payer_id"`
	Address      Address `json:"address"`
}

type Link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

type PayPalPaymentDTO struct {
	ID            string         `json:"id"`
	Intent        string         `json:"intent"`
	Status        string         `json:"status"`
	PurchaseUnits []PurchaseUnit `json:"purchase_units"`
	Payer         Payer          `json:"payer"`
	CreateTime    time.Time      `json:"create_time"`
	UpdateTime    time.Time      `json:"update_time"`
	Links         []Link         `json:"links"`
}

type PayPalGetOrderDto struct {
	ID            string `json:"id"`
	Status        string `json:"status"`
	Intent        string `json:"intent"`
	PaymentSource struct {
		Paypal struct {
			Name struct {
				GivenName string `json:"given_name"`
				Surname   string `json:"surname"`
			} `json:"name"`
			EmailAddress string `json:"email_address"`
			AccountID    string `json:"account_id"`
		} `json:"paypal"`
	} `json:"payment_source"`
	PurchaseUnits []struct {
		ReferenceID string `json:"reference_id"`
		Amount      struct {
			CurrencyCode string `json:"currency_code"`
			Value        string `json:"value"`
		} `json:"amount"`
	} `json:"purchase_units"`
	Payer struct {
		Name struct {
			GivenName string `json:"given_name"`
			Surname   string `json:"surname"`
		} `json:"name"`
		EmailAddress string `json:"email_address"`
		PayerID      string `json:"payer_id"`
	} `json:"payer"`
	CreateTime time.Time `json:"create_time"`
	Links      []struct {
		Href   string `json:"href"`
		Rel    string `json:"rel"`
		Method string `json:"method"`
	} `json:"links"`
}

type PayPalAuthDTO struct {
	Scope       string `json:"scope"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	AppID       string `json:"app_id"`
	ExpiresIn   int    `json:"expires_in"`
	Nonce       string `json:"nonce"`
}
