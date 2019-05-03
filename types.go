package krakenapi

import (
	"encoding/json"
	"math/big"
	"reflect"
	"strconv"
)

const (
	// BCHEUR is a trading pair
	BCHEUR = "BCHEUR"
	// BCHUSD is a trading pair
	BCHUSD = "BCHUSD"
	// BCHXBT is a trading pair
	BCHXBT = "BCHXBT"
	// DASHEUR is a trading pair
	DASHEUR = "DASHEUR"
	// DASHUSD is a trading pair
	DASHUSD = "DASHUSD"
	// DASHXBT is a trading pair
	DASHXBT = "DASHXBT"
	// EOSETH is a trading pair
	EOSETH = "EOSETH"
	// EOSEUR is a trading pair
	EOSEUR = "EOSEUR"
	// EOSUSD is a trading pair
	EOSUSD = "EOSUSD"
	// EOSXBT is a trading pair
	EOSXBT = "EOSXBT"
	// GNOETH is a trading pair
	GNOETH = "GNOETH"
	// GNOEUR is a trading pair
	GNOEUR = "GNOEUR"
	// GNOUSD is a trading pair
	GNOUSD = "GNOUSD"
	// GNOXBT is a trading pair
	GNOXBT = "GNOXBT"
	// USDTZUSD is a trading pair
	USDTZUSD = "USDTZUSD"
	// XETCXETH is a trading pair
	XETCXETH = "XETCXETH"
	// XETCXXBT is a trading pair
	XETCXXBT = "XETCXXBT"
	// XETCZEUR is a trading pair
	XETCZEUR = "XETCZEUR"
	// XETCXUSD is a trading pair
	XETCXUSD = "XETCXUSD"
	// XETHXXBT is a trading pair
	XETHXXBT = "XETHXXBT"
	// XETHZCAD is a trading pair
	XETHZCAD = "XETHZCAD"
	// XETHZEUR is a trading pair
	XETHZEUR = "XETHZEUR"
	// XETHZGBP is a trading pair
	XETHZGBP = "XETHZGBP"
	// XETHZJPY is a trading pair
	XETHZJPY = "XETHZJPY"
	// XETHZUSD is a trading pair
	XETHZUSD = "XETHZUSD"
	// XICNXETH is a trading pair
	XICNXETH = "XICNXETH"
	// XICNXXBT is a trading pair
	XICNXXBT = "XICNXXBT"
	// XLTCXXBT is a trading pair
	XLTCXXBT = "XLTCXXBT"
	// XLTCZEUR is a trading pair
	XLTCZEUR = "XLTCZEUR"
	// XLTCZUSD is a trading pair
	XLTCZUSD = "XLTCZUSD"
	// XMLNXETH is a trading pair
	XMLNXETH = "XMLNXETH"
	// XMLNXXBT is a trading pair
	XMLNXXBT = "XMLNXXBT"
	// XREPXETH is a trading pair
	XREPXETH = "XREPXETH"
	// XREPXXBT is a trading pair
	XREPXXBT = "XREPXXBT"
	// XREPZEUR is a trading pair
	XREPZEUR = "XREPZEUR"
	// XREPZUSD is a trading pair
	XREPZUSD = "XREPZUSD"
	// XXBTZCAD is a trading pair
	XXBTZCAD = "XXBTZCAD"
	// XXBTZEUR is a trading pair
	XXBTZEUR = "XXBTZEUR"
	// XXBTZGBP is a trading pair
	XXBTZGBP = "XXBTZGBP"
	// XXBTZJPY is a trading pair
	XXBTZJPY = "XXBTZJPY"
	// XXBTZUSD is a trading pair
	XXBTZUSD = "XXBTZUSD"
	// XXDGXXBT is a trading pair
	XXDGXXBT = "XXDGXXBT"
	// XXLMXXBT is a trading pair
	XXLMXXBT = "XXLMXXBT"
	// XXLMZEUR is a trading pair
	XXLMZEUR = "XXLMZEUR"
	// XXLMZUSD is a trading pair
	XXLMZUSD = "XXLMZUSD"
	// XXMRXXBT is a trading pair
	XXMRXXBT = "XXMRXXBT"
	// XXMRZEUR is a trading pair
	XXMRZEUR = "XXMRZEUR"
	// XXMRZUSD is a trading pair
	XXMRZUSD = "XXMRZUSD"
	// XXRPXXBT is a trading pair
	XXRPXXBT = "XXRPXXBT"
	// XXRPZCAD is a trading pair
	XXRPZCAD = "XXRPZCAD"
	// XXRPZEUR is a trading pair
	XXRPZEUR = "XXRPZEUR"
	// XXRPZJPY is a trading pair
	XXRPZJPY = "XXRPZJPY"
	// XXRPZUSD is a trading pair
	XXRPZUSD = "XXRPZUSD"
	// XZECXXBT is a trading pair
	XZECXXBT = "XZECXXBT"
	// XZECZEUR is a trading pair
	XZECZEUR = "XZECZEUR"
	// XZECZUSD is a trading pair
	XZECZUSD = "XZECZUSD"

	// XBT is bitcoin
	XBT = "XBT"
)

const (
	// BUY is a buy order type
	BUY = "buy"
	// SELL is a sell order type
	SELL = "sell"
	// MARKET is a market order type
	MARKET = "m"
	// LIMIT is a limit order type
	LIMIT = "l"
)

// KrakenResponse wraps the Kraken API JSON response
type KrakenResponse struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}

// TimeResponse represents the server's time
type TimeResponse struct {
	// Unix timestamp
	Unixtime int64
	// RFC 1123 time format
	Rfc1123 string
}

// AssetPairsResponse includes asset pair informations
type AssetPairsResponse struct {
	BCHEUR   AssetPairInfo
	BCHUSD   AssetPairInfo
	BCHXBT   AssetPairInfo
	DASHEUR  AssetPairInfo
	DASHUSD  AssetPairInfo
	DASHXBT  AssetPairInfo
	EOSETH   AssetPairInfo
	EOSEUR   AssetPairInfo
	EOSUSD   AssetPairInfo
	EOSXBT   AssetPairInfo
	GNOETH   AssetPairInfo
	GNOEUR   AssetPairInfo
	GNOUSD   AssetPairInfo
	GNOXBT   AssetPairInfo
	USDTZUSD AssetPairInfo
	XETCXETH AssetPairInfo
	XETCXXBT AssetPairInfo
	XETCZEUR AssetPairInfo
	XETCXUSD AssetPairInfo
	XETHXXBT AssetPairInfo
	XETHZCAD AssetPairInfo
	XETHZEUR AssetPairInfo
	XETHZGBP AssetPairInfo
	XETHZJPY AssetPairInfo
	XETHZUSD AssetPairInfo
	XICNXETH AssetPairInfo
	XICNXXBT AssetPairInfo
	XLTCXXBT AssetPairInfo
	XLTCZEUR AssetPairInfo
	XLTCZUSD AssetPairInfo
	XMLNXETH AssetPairInfo
	XMLNXXBT AssetPairInfo
	XREPXETH AssetPairInfo
	XREPXXBT AssetPairInfo
	XREPZEUR AssetPairInfo
	XREPZUSD AssetPairInfo
	XXBTZCAD AssetPairInfo
	XXBTZEUR AssetPairInfo
	XXBTZGBP AssetPairInfo
	XXBTZJPY AssetPairInfo
	XXBTZUSD AssetPairInfo
	XXDGXXBT AssetPairInfo
	XXLMXXBT AssetPairInfo
	XXLMZEUR AssetPairInfo
	XXLMZUSD AssetPairInfo
	XXMRXXBT AssetPairInfo
	XXMRZEUR AssetPairInfo
	XXMRZUSD AssetPairInfo
	XXRPXXBT AssetPairInfo
	XXRPZCAD AssetPairInfo
	XXRPZEUR AssetPairInfo
	XXRPZJPY AssetPairInfo
	XXRPZUSD AssetPairInfo
	XZECXXBT AssetPairInfo
	XZECZEUR AssetPairInfo
	XZECZUSD AssetPairInfo
}

// AssetPairInfo represents asset pair information
type AssetPairInfo struct {
	// Alternate pair name
	Altname string `json:"altname"`
	// Asset class of base component
	AssetClassBase string `json:"aclass_base"`
	// Asset id of base component
	Base string `json:"base"`
	// Asset class of quote component
	AssetClassQuote string `json:"aclass_quote"`
	// Asset id of quote component
	Quote string `json:"quote"`
	// Volume lot size
	Lot string `json:"lot"`
	// Scaling decimal places for pair
	PairDecimals int `json:"pair_decimals"`
	// Scaling decimal places for volume
	LotDecimals int `json:"lot_decimals"`
	// Amount to multiply lot volume by to get currency volume
	LotMultiplier int `json:"lot_multiplier"`
	// Array of leverage amounts available when buying
	LeverageBuy []float64 `json:"leverage_buy"`
	// Array of leverage amounts available when selling
	LeverageSell []float64 `json:"leverage_sell"`
	// Fee schedule array in [volume, percent fee] tuples
	Fees [][]float64 `json:"fees"`
	// // Maker fee schedule array in [volume, percent fee] tuples (if on maker/taker)
	FeesMaker [][]float64 `json:"fees_maker"`
	// // Volume discount currency
	FeeVolumeCurrency string `json:"fee_volume_currency"`
	// Margin call level
	MarginCall int `json:"margin_call"`
	// Stop-out/Liquidation margin level
	MarginStop int `json:"margin_stop"`
}

// AssetsResponse includes asset informations
type AssetsResponse struct {
	BCH  AssetInfo
	DASH AssetInfo
	EOS  AssetInfo
	GNO  AssetInfo
	KFEE AssetInfo
	USDT AssetInfo
	XDAO AssetInfo
	XETC AssetInfo
	XETH AssetInfo
	XICN AssetInfo
	XLTC AssetInfo
	XMLN AssetInfo
	XNMC AssetInfo
	XREP AssetInfo
	XXBT AssetInfo
	XXDG AssetInfo
	XXLM AssetInfo
	XXMR AssetInfo
	XXRP AssetInfo
	XXVN AssetInfo
	XZEC AssetInfo
	ZCAD AssetInfo
	ZEUR AssetInfo
	ZGBP AssetInfo
	ZJPY AssetInfo
	ZKRW AssetInfo
	ZUSD AssetInfo
}

// AssetInfo represents an asset information
type AssetInfo struct {
	// Alternate name
	Altname string
	// Asset class
	AssetClass string `json:"aclass"`
	// Scaling decimal places for record keeping
	Decimals int
	// Scaling decimal places for output display
	DisplayDecimals int `json:"display_decimals"`
}

// BalanceResponse is a list of balances
type BalanceResponse struct {
	BCH  float64 `json:"BCH,string"`
	DASH float64 `json:"DASH,string"`
	EOS  float64 `json:"EOS,string"`
	GNO  float64 `json:"GNO,string"`
	KFEE float64 `json:"KFEE,string"`
	USDT float64 `json:"USDT,string"`
	XDAO float64 `json:"XDAO,string"`
	XETC float64 `json:"XETC,string"`
	XETH float64 `json:"XETH,string"`
	XICN float64 `json:"XICN,string"`
	XLTC float64 `json:"XLTC,string"`
	XMLN float64 `json:"XMLN,string"`
	XNMC float64 `json:"XNMC,string"`
	XREP float64 `json:"XREP,string"`
	XXBT float64 `json:"XXBT,string"`
	XXDG float64 `json:"XXDG,string"`
	XXLM float64 `json:"XXLM,string"`
	XXMR float64 `json:"XXMR,string"`
	XXRP float64 `json:"XXRP,string"`
	XXVN float64 `json:"XXVN,string"`
	XZEC float64 `json:"XZEC,string"`
	ZCAD float64 `json:"ZCAD,string"`
	ZEUR float64 `json:"ZEUR,string"`
	ZGBP float64 `json:"ZGBP,string"`
	ZJPY float64 `json:"ZJPY,string"`
	ZKRW float64 `json:"ZKRW,string"`
	ZUSD float64 `json:"ZUSD,string"`
}

// TickerResponse includes the requested ticker pairs
type TickerResponse struct {
	BCHEUR   PairTickerInfo
	BCHUSD   PairTickerInfo
	BCHXBT   PairTickerInfo
	DASHEUR  PairTickerInfo
	DASHUSD  PairTickerInfo
	DASHXBT  PairTickerInfo
	EOSETH   PairTickerInfo
	EOSEUR   PairTickerInfo
	EOSUSD   PairTickerInfo
	EOSXBT   PairTickerInfo
	GNOETH   PairTickerInfo
	GNOEUR   PairTickerInfo
	GNOUSD   PairTickerInfo
	GNOXBT   PairTickerInfo
	USDTZUSD PairTickerInfo
	XETCXETH PairTickerInfo
	XETCXXBT PairTickerInfo
	XETCZEUR PairTickerInfo
	XETCXUSD PairTickerInfo
	XETHXXBT PairTickerInfo
	XETHZCAD PairTickerInfo
	XETHZEUR PairTickerInfo
	XETHZGBP PairTickerInfo
	XETHZJPY PairTickerInfo
	XETHZUSD PairTickerInfo
	XICNXETH PairTickerInfo
	XICNXXBT PairTickerInfo
	XLTCXXBT PairTickerInfo
	XLTCZEUR PairTickerInfo
	XLTCZUSD PairTickerInfo
	XMLNXETH PairTickerInfo
	XMLNXXBT PairTickerInfo
	XREPXETH PairTickerInfo
	XREPXXBT PairTickerInfo
	XREPZEUR PairTickerInfo
	XREPZUSD PairTickerInfo
	XXBTZCAD PairTickerInfo
	XXBTZEUR PairTickerInfo
	XXBTZGBP PairTickerInfo
	XXBTZJPY PairTickerInfo
	XXBTZUSD PairTickerInfo
	XXDGXXBT PairTickerInfo
	XXLMXXBT PairTickerInfo
	XXLMZEUR PairTickerInfo
	XXLMZUSD PairTickerInfo
	XXMRXXBT PairTickerInfo
	XXMRZEUR PairTickerInfo
	XXMRZUSD PairTickerInfo
	XXRPXXBT PairTickerInfo
	XXRPZCAD PairTickerInfo
	XXRPZEUR PairTickerInfo
	XXRPZJPY PairTickerInfo
	XXRPZUSD PairTickerInfo
	XZECXXBT PairTickerInfo
	XZECZEUR PairTickerInfo
	XZECZUSD PairTickerInfo
}

// DepositMethodsResponse is the response type of a DepositMethods query to the Kraken API.
type DepositMethodsResponse []struct {
	Method     string  `json:"method"`
	Limit      bool    `json:"limit"`
	Fee        float64 `json:"fee,string"`
	GenAddress bool    `json:"gen-address"`
}

// DepositAddressesResponse is the response type of a DepositAddresses query to the Kraken API.
type DepositAddressesResponse []struct {
	Address  string `json:"address"`
	Expiretm string `json:"expiretm"`
	New      bool   `json:"new,omitempty"`
}

// WithdrawResponse is the response type of a Withdraw query to the Kraken API.
type WithdrawResponse struct {
	RefID string `json:"refid"`
}

// WithdrawInfoResponse is the response type showing withdrawal information for a selected withdrawal method.
type WithdrawInfoResponse struct {
	Method string    `json:"method"`
	Limit  big.Float `json:"limit"`
	Amount big.Float `json:"amount"`
	Fee    big.Float `json:"fee"`
}

// GetPairTickerInfo is a helper method that returns given `pair`'s `PairTickerInfo`
func (v *TickerResponse) GetPairTickerInfo(pair string) PairTickerInfo {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(pair)

	return f.Interface().(PairTickerInfo)
}

// PairTickerInfo represents ticker information for a pair
type PairTickerInfo struct {
	// Ask array(<price>, <whole lot volume>, <lot volume>)
	Ask []string `json:"a"`
	// Bid array(<price>, <whole lot volume>, <lot volume>)
	Bid []string `json:"b"`
	// Last trade closed array(<price>, <lot volume>)
	Close []string `json:"c"`
	// Volume array(<today>, <last 24 hours>)
	Volume []string `json:"v"`
	// Volume weighted average price array(<today>, <last 24 hours>)
	VolumeAveragePrice []string `json:"p"`
	// Number of trades array(<today>, <last 24 hours>)
	Trades []int `json:"t"`
	// Low array(<today>, <last 24 hours>)
	Low []string `json:"l"`
	// High array(<today>, <last 24 hours>)
	High []string `json:"h"`
	// Today's opening price
	OpeningPrice float64 `json:"o,string"`
}

// TradesResponse represents a list of the last trades
type TradesResponse struct {
	Last   int64
	Trades []TradeInfo
}

// TradeInfo represents a trades information
type TradeInfo struct {
	Price         string
	PriceFloat    float64
	Volume        string
	VolumeFloat   float64
	Time          int64
	Buy           bool
	Sell          bool
	Market        bool
	Limit         bool
	Miscellaneous string
}

// OrderTypes for AddOrder
const (
	OTMarket              = "market"
	OTLimit               = "limit"                  // (price = limit price)
	OTStopLoss            = "stop-loss"              // (price = stop loss price)
	OTTakeProfi           = "take-profit"            // (price = take profit price)
	OTStopLossProfit      = "stop-loss-profit"       // (price = stop loss price, price2 = take profit price)
	OTStopLossProfitLimit = "stop-loss-profit-limit" // (price = stop loss price, price2 = take profit price)
	OTStopLossLimit       = "stop-loss-limit"        // (price = stop loss trigger price, price2 = triggered limit price)
	OTTakeProfitLimit     = "take-profit-limit"      // (price = take profit trigger price, price2 = triggered limit price)
	OTTrailingStop        = "trailing-stop"          // (price = trailing stop offset)
	OTTrailingStopLimit   = "trailing-stop-limit"    // (price = trailing stop offset, price2 = triggered limit offset)
	OTStopLossAndLimit    = "stop-loss-and-limit"    // (price = stop loss price, price2 = limit price)
	OTSettlePosition      = "settle-position"
)

// OrderDescription represents an orders description
type OrderDescription struct {
	AssetPair      string `json:"pair"`
	Close          string `json:"close"`
	Leverage       string `json:"leverage"`
	Order          string `json:"order"`
	OrderType      string `json:"ordertype"`
	PrimaryPrice   string `json:"price"`
	SecondaryPrice string `json:"price2"`
	Type           string `json:"type"`
}

// Order represents a single order
type Order struct {
	TransactionID  string           `json:"-"`
	ReferenceID    string           `json:"refid"`
	UserRef        int              `json:"userref"`
	Status         string           `json:"status"`
	OpenTime       float64          `json:"opentm"`
	StartTime      float64          `json:"starttm"`
	ExpireTime     float64          `json:"expiretm"`
	Description    OrderDescription `json:"descr"`
	Volume         string           `json:"vol"`
	VolumeExecuted float64          `json:"vol_exec,string"`
	Cost           float64          `json:"cost,string"`
	Fee            float64          `json:"fee,string"`
	Price          float64          `json:"price,string"`
	StopPrice      float64          `json:"stopprice.string"`
	LimitPrice     float64          `json:"limitprice,string"`
	Misc           string           `json:"misc"`
	OrderFlags     string           `json:"oflags"`
	CloseTime      float64          `json:"closetm"`
	Reason         string           `json:"reason"`
}

// ClosedOrdersResponse represents a list of closed orders, indexed by id
type ClosedOrdersResponse struct {
	Closed map[string]Order `json:"closed"`
	Count  int              `json:"count"`
}

// OrderBookItem is a piece of information about an order.
type OrderBookItem struct {
	Price  float64
	Amount float64
	Ts     int64
}

// UnmarshalJSON takes a json array from kraken and converts it into an OrderBookItem.
func (o *OrderBookItem) UnmarshalJSON(data []byte) error {
	tmpStruct := struct {
		price  string
		amount string
		ts     int64
	}{}
	tmpArr := []interface{}{&tmpStruct.price, &tmpStruct.amount, &tmpStruct.ts}
	err := json.Unmarshal(data, &tmpArr)
	if err != nil {
		return err
	}

	o.Price, err = strconv.ParseFloat(tmpStruct.price, 64)
	if err != nil {
		return err
	}
	o.Amount, err = strconv.ParseFloat(tmpStruct.amount, 64)
	if err != nil {
		return err
	}
	o.Ts = tmpStruct.ts
	return nil
}

// DepthResponse is a response from kraken to Depth request.
type DepthResponse map[string]OrderBook

// OrderBook contains top asks and bids.
type OrderBook struct {
	Asks []OrderBookItem
	Bids []OrderBookItem
}

// OpenOrdersResponse is a list of open orders
type OpenOrdersResponse struct {
	Open  map[string]Order `json:"open"`
	Count int              `json:"count"`
}

// AddOrderResponse is the response from adding an order
type AddOrderResponse struct {
	Description    OrderDescription `json:"descr"`
	TransactionIds []string         `json:"txid"`
}

// CancelOrderResponse is the respsonse from cancelling an order
type CancelOrderResponse struct {
	Count   int  `json:"count"`
	Pending bool `json:"pending"`
}

// QueryOrdersResponse is a map of orders
type QueryOrdersResponse map[string]Order
