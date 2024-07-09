package cache

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestMarketCache_Parse(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		m       *MarketCache
		args    args
		wantErr bool
	}{
		{
			name: "valid message",
			m: &MarketCache{Markets: map[string]*Market{"1.230452417": NewMarket("1.230452417")},
				HeartbeatThreshold: time.Duration(5000) * time.Millisecond},
			args: args{
				msg: fmt.Sprintf(`{"op":"mcm","id":2,"clk":"ANUIAPwJAPYJ","pt":%d,"mc":[{"id":"1.230452417","rc":[{"bdatl":[[1,2.96,33917.08],[3,3,29312.91],[5,3.1,7139.96],[6,3.15,20834.34]],"id":58805},{"bdatl":[[0,3.25,25573.93],[2,3.35,25993.63],[3,3.4,20385.34],[4,3.45,9726.3]],"id":24},{"bdatb":[[0,2.86,10459.4]],"id":22}]}]}`, time.Now().UnixNano()/int64(time.Millisecond)+500),
			},
			wantErr: false,
		},
		{
			name: "invalid JSON",
			m:    &MarketCache{},
			args: args{
				msg: `{"op":"mcm","id":2,"clk":"APIBAPIB,"mc":[{"id":"1.l":[[0,2.94,10852.98],"id":58805},{"bdatb":[[4,2.78,32537.13],[5,2.76,35723.67],[7,2.72,1774.4]],"id":22}]}]}`,
			},
			wantErr: true,
		},
		{
			name: "market not in cache",
			m:    &MarketCache{Markets: map[string]*Market{}},
			args: args{
				msg: fmt.Sprintf(`{"op":"mcm","id":2,"clk":"ANUIAPwJAPYJ","pt":%d,"mc":[{"id":"1.230452417","rc":[{"bdatl":[[1,2.96,33917.08],[3,3,29312.91],[5,3.1,7139.96],[6,3.15,20834.34]],"id":58805},{"bdatl":[[0,3.25,25573.93],[2,3.35,25993.63],[3,3.4,20385.34],[4,3.45,9726.3]],"id":24},{"bdatb":[[0,2.86,10459.4]],"id":22}]}]}`, time.Now().UnixNano()/int64(time.Millisecond)),
			},
			wantErr: true,
		},
		{
			name: "high latency",
			m: &MarketCache{Markets: map[string]*Market{"1.230452417": NewMarket("1.230452417")},
				HeartbeatThreshold: time.Duration(5000) * time.Millisecond},
			args: args{
				msg: fmt.Sprintf(`{"op":"mcm","id":2,"clk":"ANUIAPwJAPYJ","pt":%d,"mc":[{"id":"1.230452417","rc":[{"bdatl":[[1,2.96,33917.08],[3,3,29312.91],[5,3.1,7139.96],[6,3.15,20834.34]],"id":58805},{"bdatl":[[0,3.25,25573.93],[2,3.35,25993.63],[3,3.4,20385.34],[4,3.45,9726.3]],"id":24},{"bdatb":[[0,2.86,10459.4]],"id":22}]}]}`, time.Now().UnixNano()/int64(time.Millisecond)-1500),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Parse(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("MarketCache.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMarketCache_resetTimer(t *testing.T) {
	tests := []struct {
		name string
		m    *MarketCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.resetTimer()
		})
	}
}

func TestMarketCache_StopTimer(t *testing.T) {
	tests := []struct {
		name string
		m    *MarketCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.StopTimer()
		})
	}
}

func TestMarketCache_AddMarket(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		m    *MarketCache
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.AddMarket(tt.args.id)
		})
	}
}

func TestNewMarketCache(t *testing.T) {
	tests := []struct {
		name string
		want *MarketCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMarketCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMarketCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
