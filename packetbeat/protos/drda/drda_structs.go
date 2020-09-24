package drda

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/packetbeat/publish"
	"time"
)

type Ddm struct {
	Length    uint16
	Cor       uint16
	Format    uint8
	Length2   uint16
	Codepoint uint16
	DSSType   uint8
	DSSFlags  uint8
}

type Parameter struct {
	Length     uint16
	Codepoint  uint16
	ASCIIData  string
	EBCDICData string
}

type DrdaMessage struct {
	start           int
	end             int
	ddm             Ddm
	parameters      map[uint16]Parameter
	RemainingLength int
	Direction       uint8
	TcpTuple        common.TCPTuple
	//CmdlineTuple    *common.BaseTuple
	//Raw          []byte
	Notes []string
}

type DrdaTransaction struct {
	Type         string
	tuple        common.TCPTuple
	Src          common.Endpoint
	Dst          common.Endpoint
	ResponseTime int32
	TsStart      time.Time
	TsEnd        time.Time
	BytesOut     uint64
	BytesIn      uint64
	Notes        []string
	Requests     common.MapStr
	Responses    common.MapStr
	Query        string
	//Request_raw  string
	//Response_raw string
}

type DrdaStream struct {
	tcptuple    *common.TCPTuple
	data        []byte
	parseOffset int
	//parseState  parsestate
	message *DrdaMessage
}

type drdaPlugin struct {
	// config
	Ports         []int
	Send_request  bool
	Send_response bool

	transactions       *common.Cache
	transactionTimeout time.Duration

	results publish.TransactionPublisher

	// function pointer for mocking
	handleDrda func(drda *drdaPlugin, m *DrdaMessage, tcp *common.TCPTuple,
		dir uint8, raw_msg []byte)
}
