package v2

import (
	"bytes"
	"embed"
	"log"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"golang.org/x/exp/slices"

	"github.com/smartcontractkit/chainlink/core/utils"
)

var (
	//go:embed defaults/*.toml
	defaultsFS   embed.FS
	fallback     Chain
	defaults     = map[string]Chain{}
	defaultNames = map[string]string{}

	// DefaultIDs is the set of chain ids which have defaults.
	DefaultIDs []*utils.Big
)

//TODO docs only?
func DefaultName(id *utils.Big) string {
	if id == nil {
		return ""
	}
	return defaultNames[id.String()]
}

func init() {
	fes, err := defaultsFS.ReadDir("defaults")
	if err != nil {
		log.Fatalf("failed to read defaults/: %v", err)
	}
	for _, fe := range fes {
		path := filepath.Join("defaults", fe.Name())
		b, err := defaultsFS.ReadFile(path)
		if err != nil {
			log.Fatalf("failed to read %q: %v", path, err)
		}
		var config = struct {
			ChainID *utils.Big
			Chain
		}{}
		d := toml.NewDecoder(bytes.NewReader(b)).DisallowUnknownFields()
		if err := d.Decode(&config); err != nil {
			log.Fatalf("failed to decode %q: %v", path, err)
		}
		if fe.Name() == "fallback.toml" {
			if config.ChainID != nil {
				log.Fatalf("fallback ChainID must be nil, not: %s", config.ChainID)
			}
			fallback = config.Chain
			continue
		}
		if config.ChainID == nil {
			log.Fatalf("missing ChainID: %s", path)
		}
		DefaultIDs = append(DefaultIDs, config.ChainID)
		id := config.ChainID.String()
		if _, ok := defaults[id]; ok {
			log.Fatalf("%q contains duplicate ChainID: %s", path, id)
		}
		defaults[id] = config.Chain
		defaultNames[id] = strings.ReplaceAll(strings.TrimSuffix(fe.Name(), ".toml"), "_", " ")
	}
	slices.SortFunc(DefaultIDs, func(a, b *utils.Big) bool {
		return a.Cmp(b) < 0
	})
}

// SetDefaults sets the Chain default values, optionally for a specific chain id.
func (c *Chain) SetDefaults(chainID *utils.Big) {
	c.SetFrom(&fallback)
	if chainID == nil {
		return
	}
	if d, ok := defaults[chainID.String()]; ok {
		c.SetFrom(&d)
	}
}

func (c *Chain) SetFrom(f *Chain) {
	if v := f.BalanceMonitorEnabled; v != nil {
		c.BalanceMonitorEnabled = v
	}
	if v := f.BlockBackfillDepth; v != nil {
		c.BlockBackfillDepth = v
	}
	if v := f.BlockBackfillSkip; v != nil {
		c.BlockBackfillSkip = v
	}
	if v := f.ChainType; v != nil {
		c.ChainType = v
	}
	if v := f.EIP1559DynamicFees; v != nil {
		c.EIP1559DynamicFees = v
	}
	if v := f.FinalityDepth; v != nil {
		c.FinalityDepth = v
	}
	if v := f.FlagsContractAddress; v != nil {
		c.FlagsContractAddress = v
	}
	if v := f.GasBumpPercent; v != nil {
		c.GasBumpPercent = v
	}
	if v := f.GasBumpThreshold; v != nil {
		c.GasBumpThreshold = v
	}
	if v := f.GasBumpTxDepth; v != nil {
		c.GasBumpTxDepth = v
	}
	if v := f.GasBumpWei; v != nil {
		c.GasBumpWei = v
	}
	if v := f.GasEstimatorMode; v != nil {
		c.GasEstimatorMode = v
	}
	if v := f.GasFeeCapDefault; v != nil {
		c.GasFeeCapDefault = v
	}
	if v := f.GasLimitDefault; v != nil {
		c.GasLimitDefault = v
	}
	if v := f.GasLimitMultiplier; v != nil {
		c.GasLimitMultiplier = v
	}
	if v := f.GasLimitTransfer; v != nil {
		c.GasLimitTransfer = v
	}
	if v := f.GasPriceDefault; v != nil {
		c.GasPriceDefault = v
	}
	if v := f.GasTipCapDefault; v != nil {
		c.GasTipCapDefault = v
	}
	if v := f.GasTipCapMinimum; v != nil {
		c.GasTipCapMinimum = v
	}
	if v := f.LinkContractAddress; v != nil {
		c.LinkContractAddress = v
	}
	if v := f.LogBackfillBatchSize; v != nil {
		c.LogBackfillBatchSize = v
	}
	if v := f.LogPollInterval; v != nil {
		c.LogPollInterval = v
	}
	if v := f.MaxGasPriceWei; v != nil {
		c.MaxGasPriceWei = v
	}
	if v := f.MaxInFlightTransactions; v != nil {
		c.MaxInFlightTransactions = v
	}
	if v := f.MaxQueuedTransactions; v != nil {
		c.MaxQueuedTransactions = v
	}
	if v := f.MinGasPriceWei; v != nil {
		c.MinGasPriceWei = v
	}
	if v := f.MinIncomingConfirmations; v != nil {
		c.MinIncomingConfirmations = v
	}
	if v := f.MinimumContractPayment; v != nil {
		c.MinimumContractPayment = v
	}
	if v := f.NonceAutoSync; v != nil {
		c.NonceAutoSync = v
	}
	if v := f.OCRContractConfirmations; v != nil {
		c.OCRContractConfirmations = v
	}
	if v := f.OCRContractTransmitterTransmitTimeout; v != nil {
		c.OCRContractTransmitterTransmitTimeout = v
	}
	if v := f.OCRDatabaseTimeout; v != nil {
		c.OCRDatabaseTimeout = v
	}
	if v := f.OCRObservationTimeout; v != nil {
		c.OCRObservationTimeout = v
	}
	if v := f.OCRObservationGracePeriod; v != nil {
		c.OCRObservationGracePeriod = v
	}
	if v := f.OCR2ContractConfirmations; v != nil {
		c.OCR2ContractConfirmations = v
	}
	if v := f.OperatorFactoryAddress; v != nil {
		c.OperatorFactoryAddress = v
	}
	if v := f.RPCDefaultBatchSize; v != nil {
		c.RPCDefaultBatchSize = v
	}
	if v := f.TxReaperInterval; v != nil {
		c.TxReaperInterval = v
	}
	if v := f.TxReaperThreshold; v != nil {
		c.TxReaperThreshold = v
	}
	if v := f.TxResendAfterThreshold; v != nil {
		c.TxResendAfterThreshold = v
	}
	if v := f.UseForwarders; v != nil {
		c.UseForwarders = v
	}
	if b := f.BlockHistoryEstimator; b != nil {
		if c.BlockHistoryEstimator == nil {
			c.BlockHistoryEstimator = &BlockHistoryEstimator{}
		}
		if v := b.BatchSize; v != nil {
			c.BlockHistoryEstimator.BatchSize = v
		}
		if v := b.BlockDelay; v != nil {
			c.BlockHistoryEstimator.BlockDelay = v
		}
		if v := b.BlockHistorySize; v != nil {
			c.BlockHistoryEstimator.BlockHistorySize = v
		}
		if v := b.EIP1559FeeCapBufferBlocks; v != nil {
			c.BlockHistoryEstimator.EIP1559FeeCapBufferBlocks = v
		}
		if v := b.TransactionPercentile; v != nil {
			c.BlockHistoryEstimator.TransactionPercentile = v
		}
	}
	// skip KeySpecific
	if h := f.HeadTracker; h != nil {
		if c.HeadTracker == nil {
			c.HeadTracker = &HeadTracker{}
		}
		if v := h.BlockEmissionIdleWarningThreshold; v != nil {
			c.HeadTracker.BlockEmissionIdleWarningThreshold = v
		}
		if v := h.HistoryDepth; v != nil {
			c.HeadTracker.HistoryDepth = v
		}
		if v := h.MaxBufferSize; v != nil {
			c.HeadTracker.MaxBufferSize = v
		}
		if v := h.SamplingInterval; v != nil {
			c.HeadTracker.SamplingInterval = v
		}
	}
	if n := f.NodePool; n != nil {
		if c.NodePool == nil {
			c.NodePool = &NodePool{}
		}
		if v := n.NoNewHeadsThreshold; v != nil {
			c.NodePool.NoNewHeadsThreshold = v
		}
		if v := n.PollFailureThreshold; v != nil {
			c.NodePool.PollFailureThreshold = v
		}
		if v := n.PollInterval; v != nil {
			c.NodePool.PollInterval = v
		}
	}
}