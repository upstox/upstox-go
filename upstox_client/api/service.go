// Code generated from the swagger-codegen output; DO NOT EDIT.

package api

import "github.com/upstox/upstox-go/upstox_client"

type service struct {
	client *upstox_client.APIClient
}

// NewChargeApi returns the Charge API service bound to client.
func NewChargeApi(client *upstox_client.APIClient) *ChargeApiService {
	return &ChargeApiService{client: client}
}

// NewExpiredInstrumentApi returns the ExpiredInstrument API service bound to client.
func NewExpiredInstrumentApi(client *upstox_client.APIClient) *ExpiredInstrumentApiService {
	return &ExpiredInstrumentApiService{client: client}
}

// NewFundamentalsApi returns the Fundamentals API service bound to client.
func NewFundamentalsApi(client *upstox_client.APIClient) *FundamentalsApiService {
	return &FundamentalsApiService{client: client}
}

// NewHistoryApi returns the History API service bound to client.
func NewHistoryApi(client *upstox_client.APIClient) *HistoryApiService {
	return &HistoryApiService{client: client}
}

// NewHistoryV3Api returns the HistoryV3 API service bound to client.
func NewHistoryV3Api(client *upstox_client.APIClient) *HistoryV3ApiService {
	return &HistoryV3ApiService{client: client}
}

// NewIPOApi returns the IPO API service bound to client.
func NewIPOApi(client *upstox_client.APIClient) *IPOApiService {
	return &IPOApiService{client: client}
}

// NewInstrumentsApi returns the Instruments API service bound to client.
func NewInstrumentsApi(client *upstox_client.APIClient) *InstrumentsApiService {
	return &InstrumentsApiService{client: client}
}

// NewLoginApi returns the Login API service bound to client.
func NewLoginApi(client *upstox_client.APIClient) *LoginApiService {
	return &LoginApiService{client: client}
}

// NewMarketHolidaysAndTimingsApi returns the MarketHolidaysAndTimings API service bound to client.
func NewMarketHolidaysAndTimingsApi(client *upstox_client.APIClient) *MarketHolidaysAndTimingsApiService {
	return &MarketHolidaysAndTimingsApiService{client: client}
}

// NewMarketQuoteApi returns the MarketQuote API service bound to client.
func NewMarketQuoteApi(client *upstox_client.APIClient) *MarketQuoteApiService {
	return &MarketQuoteApiService{client: client}
}

// NewMarketQuoteV3Api returns the MarketQuoteV3 API service bound to client.
func NewMarketQuoteV3Api(client *upstox_client.APIClient) *MarketQuoteV3ApiService {
	return &MarketQuoteV3ApiService{client: client}
}

// NewMutualFundApi returns the MutualFund API service bound to client.
func NewMutualFundApi(client *upstox_client.APIClient) *MutualFundApiService {
	return &MutualFundApiService{client: client}
}

// NewNewsApi returns the News API service bound to client.
func NewNewsApi(client *upstox_client.APIClient) *NewsApiService {
	return &NewsApiService{client: client}
}

// NewOptionsApi returns the Options API service bound to client.
func NewOptionsApi(client *upstox_client.APIClient) *OptionsApiService {
	return &OptionsApiService{client: client}
}

// NewOrderApi returns the Order API service bound to client.
func NewOrderApi(client *upstox_client.APIClient) *OrderApiService {
	return &OrderApiService{client: client}
}

// NewOrderV3Api returns the OrderV3 API service bound to client.
func NewOrderV3Api(client *upstox_client.APIClient) *OrderV3ApiService {
	return &OrderV3ApiService{client: client}
}

// NewPortfolioApi returns the Portfolio API service bound to client.
func NewPortfolioApi(client *upstox_client.APIClient) *PortfolioApiService {
	return &PortfolioApiService{client: client}
}

// NewPostTradeApi returns the PostTrade API service bound to client.
func NewPostTradeApi(client *upstox_client.APIClient) *PostTradeApiService {
	return &PostTradeApiService{client: client}
}

// NewTradeProfitAndLossApi returns the TradeProfitAndLoss API service bound to client.
func NewTradeProfitAndLossApi(client *upstox_client.APIClient) *TradeProfitAndLossApiService {
	return &TradeProfitAndLossApiService{client: client}
}

// NewUSMarketHolidaysAndTimingsApi returns the USMarketHolidaysAndTimings API service bound to client.
func NewUSMarketHolidaysAndTimingsApi(client *upstox_client.APIClient) *USMarketHolidaysAndTimingsApiService {
	return &USMarketHolidaysAndTimingsApiService{client: client}
}

// NewUSStocksChargesApi returns the USStocksCharges API service bound to client.
func NewUSStocksChargesApi(client *upstox_client.APIClient) *USStocksChargesApiService {
	return &USStocksChargesApiService{client: client}
}

// NewUSStocksFundsApi returns the USStocksFunds API service bound to client.
func NewUSStocksFundsApi(client *upstox_client.APIClient) *USStocksFundsApiService {
	return &USStocksFundsApiService{client: client}
}

// NewUSStocksOrderApi returns the USStocksOrder API service bound to client.
func NewUSStocksOrderApi(client *upstox_client.APIClient) *USStocksOrderApiService {
	return &USStocksOrderApiService{client: client}
}

// NewUSStocksPortfolioApi returns the USStocksPortfolio API service bound to client.
func NewUSStocksPortfolioApi(client *upstox_client.APIClient) *USStocksPortfolioApiService {
	return &USStocksPortfolioApiService{client: client}
}

// NewUserApi returns the User API service bound to client.
func NewUserApi(client *upstox_client.APIClient) *UserApiService {
	return &UserApiService{client: client}
}

// NewWebsocketApi returns the Websocket API service bound to client.
func NewWebsocketApi(client *upstox_client.APIClient) *WebsocketApiService {
	return &WebsocketApiService{client: client}
}
