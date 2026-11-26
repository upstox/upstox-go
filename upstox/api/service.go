// Code generated from the swagger-codegen output; DO NOT EDIT.

package api

import "github.com/upstox/upstox-go/upstox"

type service struct {
	client *upstox.APIClient
}

// NewChargeApi returns the Charge API service bound to client.
func NewChargeApi(client *upstox.APIClient) *ChargeApiService {
	return &ChargeApiService{client: client}
}

// NewExpiredInstrumentApi returns the ExpiredInstrument API service bound to client.
func NewExpiredInstrumentApi(client *upstox.APIClient) *ExpiredInstrumentApiService {
	return &ExpiredInstrumentApiService{client: client}
}

// NewFundamentalsApi returns the Fundamentals API service bound to client.
func NewFundamentalsApi(client *upstox.APIClient) *FundamentalsApiService {
	return &FundamentalsApiService{client: client}
}

// NewHistoryApi returns the History API service bound to client.
func NewHistoryApi(client *upstox.APIClient) *HistoryApiService {
	return &HistoryApiService{client: client}
}

// NewHistoryV3Api returns the HistoryV3 API service bound to client.
func NewHistoryV3Api(client *upstox.APIClient) *HistoryV3ApiService {
	return &HistoryV3ApiService{client: client}
}

// NewIPOApi returns the IPO API service bound to client.
func NewIPOApi(client *upstox.APIClient) *IPOApiService {
	return &IPOApiService{client: client}
}

// NewInstrumentsApi returns the Instruments API service bound to client.
func NewInstrumentsApi(client *upstox.APIClient) *InstrumentsApiService {
	return &InstrumentsApiService{client: client}
}

// NewLoginApi returns the Login API service bound to client.
func NewLoginApi(client *upstox.APIClient) *LoginApiService {
	return &LoginApiService{client: client}
}

// NewMarketHolidaysAndTimingsApi returns the MarketHolidaysAndTimings API service bound to client.
func NewMarketHolidaysAndTimingsApi(client *upstox.APIClient) *MarketHolidaysAndTimingsApiService {
	return &MarketHolidaysAndTimingsApiService{client: client}
}

// NewMarketQuoteApi returns the MarketQuote API service bound to client.
func NewMarketQuoteApi(client *upstox.APIClient) *MarketQuoteApiService {
	return &MarketQuoteApiService{client: client}
}

// NewMarketQuoteV3Api returns the MarketQuoteV3 API service bound to client.
func NewMarketQuoteV3Api(client *upstox.APIClient) *MarketQuoteV3ApiService {
	return &MarketQuoteV3ApiService{client: client}
}

// NewMutualFundApi returns the MutualFund API service bound to client.
func NewMutualFundApi(client *upstox.APIClient) *MutualFundApiService {
	return &MutualFundApiService{client: client}
}

// NewNewsApi returns the News API service bound to client.
func NewNewsApi(client *upstox.APIClient) *NewsApiService {
	return &NewsApiService{client: client}
}

// NewOptionsApi returns the Options API service bound to client.
func NewOptionsApi(client *upstox.APIClient) *OptionsApiService {
	return &OptionsApiService{client: client}
}

// NewOrderApi returns the Order API service bound to client.
func NewOrderApi(client *upstox.APIClient) *OrderApiService {
	return &OrderApiService{client: client}
}

// NewOrderV3Api returns the OrderV3 API service bound to client.
func NewOrderV3Api(client *upstox.APIClient) *OrderV3ApiService {
	return &OrderV3ApiService{client: client}
}

// NewPortfolioApi returns the Portfolio API service bound to client.
func NewPortfolioApi(client *upstox.APIClient) *PortfolioApiService {
	return &PortfolioApiService{client: client}
}

// NewPostTradeApi returns the PostTrade API service bound to client.
func NewPostTradeApi(client *upstox.APIClient) *PostTradeApiService {
	return &PostTradeApiService{client: client}
}

// NewTradeProfitAndLossApi returns the TradeProfitAndLoss API service bound to client.
func NewTradeProfitAndLossApi(client *upstox.APIClient) *TradeProfitAndLossApiService {
	return &TradeProfitAndLossApiService{client: client}
}

// NewUSMarketHolidaysAndTimingsApi returns the USMarketHolidaysAndTimings API service bound to client.
func NewUSMarketHolidaysAndTimingsApi(client *upstox.APIClient) *USMarketHolidaysAndTimingsApiService {
	return &USMarketHolidaysAndTimingsApiService{client: client}
}

// NewUSStocksChargesApi returns the USStocksCharges API service bound to client.
func NewUSStocksChargesApi(client *upstox.APIClient) *USStocksChargesApiService {
	return &USStocksChargesApiService{client: client}
}

// NewUSStocksFundsApi returns the USStocksFunds API service bound to client.
func NewUSStocksFundsApi(client *upstox.APIClient) *USStocksFundsApiService {
	return &USStocksFundsApiService{client: client}
}

// NewUSStocksOrderApi returns the USStocksOrder API service bound to client.
func NewUSStocksOrderApi(client *upstox.APIClient) *USStocksOrderApiService {
	return &USStocksOrderApiService{client: client}
}

// NewUSStocksPortfolioApi returns the USStocksPortfolio API service bound to client.
func NewUSStocksPortfolioApi(client *upstox.APIClient) *USStocksPortfolioApiService {
	return &USStocksPortfolioApiService{client: client}
}

// NewUserApi returns the User API service bound to client.
func NewUserApi(client *upstox.APIClient) *UserApiService {
	return &UserApiService{client: client}
}

// NewWebsocketApi returns the Websocket API service bound to client.
func NewWebsocketApi(client *upstox.APIClient) *WebsocketApiService {
	return &WebsocketApiService{client: client}
}
