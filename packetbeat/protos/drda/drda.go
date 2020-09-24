package drda

import (
	"fmt"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/go-ucfg"
	"strings"
	"time"

	"github.com/Intermernet/ebcdic"
	"github.com/elastic/beats/packetbeat/procs"
	"github.com/elastic/beats/packetbeat/protos"
	"github.com/elastic/beats/packetbeat/protos/tcp"
	"github.com/elastic/beats/packetbeat/publish"
)

func init() {
	protos.Register("drda", New)
}

func New(
	testMode bool,
	results publish.TransactionPublisher,
	cfg *common.Config,
) (protos.Plugin, error) {
	p := &drdaPlugin{}
	config := defaultConfig
	if !testMode {
		if err := cfg.Unpack(&config); err != nil {
			return nil, err
		}
	}

	if err := p.init(results, &config); err != nil {
		return nil, err
	}
	return p, nil
}

func (drda *drdaPlugin) init(results publish.TransactionPublisher, config *drdaConfig) error {
	drda.setFromConfig(config)

	drda.transactions = common.NewCache(
		drda.transactionTimeout,
		protos.DefaultTransactionHashSize)
	drda.transactions.StartJanitor(drda.transactionTimeout)
	drda.results = results

	return nil
}

func (drda *drdaPlugin) setFromConfig(config *drdaConfig) {
	drda.Ports = config.Ports
	drda.Send_request = config.SendRequest
	drda.Send_response = config.SendResponse
	drda.transactionTimeout = config.TransactionTimeout
}

func (drda *drdaPlugin) GetPorts() []int {
	return drda.Ports
}
