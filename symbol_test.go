package kucoin

import (
	"testing"
	"time"
)

func TestApiService_Symbols(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Symbols()
	if err != nil {
		t.Fatal(err)
	}
	l := SymbolsModel{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}
	for _, c := range l {
		t.Log(ToJsonString(c))
		switch {
		case c.Name == "":
			t.Error("Empty key 'name'")
		case c.Symbol == "":
			t.Error("Empty key 'symbol'")
		case c.BaseCurrency == "":
			t.Error("Empty key 'baseCurrency'")
		case c.QuoteCurrency == "":
			t.Error("Empty key 'quoteCurrency'")
		case c.BaseMinSize == "":
			t.Error("Empty key 'baseMinSize'")
		case c.QuoteMinSize == "":
			t.Error("Empty key 'quoteMinSize'")
		case c.BaseMaxSize == "":
			t.Error("Empty key 'baseMaxSize'")
		case c.QuoteMaxSize == "":
			t.Error("Empty key 'quoteMaxSize'")
		case c.BaseIncrement == "":
			t.Error("Empty key 'baseIncrement'")
		case c.QuoteIncrement == "":
			t.Error("Empty key 'quoteIncrement'")
		case c.PriceIncrement == "":
			t.Error("Empty key 'priceIncrement'")
		}
	}
}

func TestApiService_Ticker(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Ticker("ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	tk := &TickerModel{}
	if err := rsp.ReadData(tk); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(tk))
	switch {
	case tk.Sequence == "":
		t.Error("Empty key 'sequence'")
	case tk.Price == "":
		t.Error("Empty key 'price'")
	case tk.Size == "":
		t.Error("Empty key 'size'")
	case tk.BestBid == "":
		t.Error("Empty key 'bestBid'")
	case tk.BestBidSize == "":
		t.Error("Empty key 'bestBidSize'")
	case tk.BestAsk == "":
		t.Error("Empty key 'bestAsk'")
	case tk.BestAskSize == "":
		t.Error("Empty key 'bestAskSize'")
	}
}

func TestApiService_MarketTickers(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.MarketTickers()
	if err != nil {
		t.Fatal(err)
	}
	ts := MarketTickersModel{}
	if err := rsp.ReadData(&ts); err != nil {
		t.Fatal(err)
	}
	for _, tk := range ts {
		switch {
		case tk.Symbol == "":
			t.Error("Empty key 'symbol'")
		case tk.High == "":
			t.Error("Empty key 'high'")
		case tk.Vol == "":
			t.Error("Empty key 'vol'")
		case tk.Low == "":
			t.Error("Empty key 'low'")
		//case tk.ChangePrice == "":
		//	t.Error("Empty key 'changePrice'")
		case tk.ChangeRate == "":
			t.Error("Empty key 'changeRate'")
		case tk.Close == "":
			t.Error("Empty key 'close'")
		case tk.VolValue == "":
			t.Error("Empty key 'volValue'")
		case tk.Open == "":
			t.Error("Empty key 'open'")
		}
	}
}

func TestApiService_Stats24hr(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Stats24hr("ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	st := &Stats24hrModel{}
	if err := rsp.ReadData(st); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(st))
	switch {
	case st.Symbol == "":
		t.Error("Empty key 'symbol'")
	case st.ChangeRate == "":
		t.Error("Empty key 'changRate'")
	case st.ChangePrice == "":
		t.Error("Empty key 'changPrice'")
	case st.High == "":
		t.Error("Empty key 'high'")
	case st.Low == "":
		t.Error("Empty key 'low'")
	case st.Vol == "":
		t.Error("Empty key 'vol'")
	case st.VolValue == "":
		t.Error("Empty key 'volValue'")
	}
}

func TestApiService_Markets(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Markets()
	if err != nil {
		t.Fatal(err)
	}
	ms := make([]string, 0)
	if err := rsp.ReadData(&ms); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(ms))
	if len(ms) == 0 {
		t.Error("Empty markets")
	}
}

func TestApiService_PartOrderBook(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.PartOrderBook("ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	c := &PartOrderBookModel{}
	if err := rsp.ReadData(c); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(c))
	switch {
	case c.Sequence == "":
		t.Error("Empty key 'sequence'")
	case len(c.Asks) == 0:
		t.Error("Empty key 'asks'")
	case len(c.Asks[0]) != 2:
		t.Error("Invalid ask length")
	case len(c.Bids) == 0:
		t.Error("Empty key 'bids'")
	case len(c.Bids[0]) != 2:
		t.Error("Invalid bid length")
	}
}

func TestApiService_AggregatedFullOrderBook(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.AggregatedFullOrderBook("ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	c := &FullOrderBookModel{}
	if err := rsp.ReadData(c); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(c))
	switch {
	case c.Sequence == "":
		t.Error("Empty key 'sequence'")
	case len(c.Asks) == 0:
		t.Error("Empty key 'asks'")
	case len(c.Asks[0]) != 2:
		t.Error("Invalid ask length")
	case len(c.Bids) == 0:
		t.Error("Empty key 'bids'")
	case len(c.Bids[0]) != 2:
		t.Error("Invalid bid length")
	}
}

func TestApiService_AtomicFullOrderBook(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.AtomicFullOrderBook("ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	c := &FullOrderBookModel{}
	if err := rsp.ReadData(c); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(c))
	switch {
	case c.Sequence == "":
		t.Error("Empty key 'sequence'")
	case len(c.Asks) == 0:
		t.Error("Empty key 'asks'")
	case len(c.Asks[0]) != 3:
		t.Error("Invalid ask length")
	case len(c.Bids) == 0:
		t.Error("Empty key 'bids'")
	case len(c.Bids[0]) != 3:
		t.Error("Invalid bid length")
	}
}

func TestApiService_TradeHistories(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.TradeHistories("ETH-BTC")
	if err != nil {
		t.Fatal(err)
	}
	l := TradeHistoriesModel{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}
	for _, c := range l {
		t.Log(ToJsonString(c))
		switch {
		case c.Sequence == "":
			t.Error("Empty key 'sequence'")
		case c.Price == "":
			t.Error("Empty key 'price'")
		case c.Size == "":
			t.Error("Empty key 'size'")
		case c.Side == "":
			t.Error("Empty key 'side'")
		case c.Time == 0:
			t.Error("Empty key 'time'")
		}
	}
}

func TestApiService_HistoricRates(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.HistoricRates("ETH-BTC", time.Now().Unix()-7*24*3600, time.Now().Unix(), "30min")
	if err != nil {
		t.Fatal(err)
	}
	l := HistoricRatesModel{}
	if err := rsp.ReadData(&l); err != nil {
		t.Fatal(err)
	}
	for _, c := range l {
		t.Log(ToJsonString(c))
		if len(*c) != 7 {
			t.Error("Invalid length of rate")
		}
	}
}