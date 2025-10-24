package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics holds all the application metrics
type Metrics struct {
	// Blockchain metrics
	BlocksMinedTotal          prometheus.Counter
	TransactionsTotal         prometheus.Counter
	BlockchainHeight          prometheus.Gauge
	Difficulty                prometheus.Gauge
	BlockProcessingTime       prometheus.Histogram
	TransactionProcessingTime prometheus.Histogram

	// Network metrics
	PeersConnected        prometheus.Gauge
	MessagesSentTotal     prometheus.Counter
	MessagesReceivedTotal prometheus.Counter
	BytesSentTotal        prometheus.Counter
	BytesReceivedTotal    prometheus.Counter

	// API metrics
	APIRequestsTotal   *prometheus.CounterVec
	APIRequestDuration *prometheus.HistogramVec
	APIErrorsTotal     prometheus.Counter

	// Wallet metrics
	WalletTransactionsTotal prometheus.Counter
	WalletBalance           prometheus.Gauge
}

// NewMetrics creates a new metrics instance
func NewMetrics() *Metrics {
	return &Metrics{
		// Blockchain metrics
		BlocksMinedTotal: promauto.NewCounter(prometheus.CounterOpts{
			Name: "coubcore_blocks_mined_total",
			Help: "Total number of blocks mined",
		}),
		TransactionsTotal: promauto.NewCounter(prometheus.CounterOpts{
			Name: "coubcore_transactions_total",
			Help: "Total number of transactions processed",
		}),
		BlockchainHeight: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "coubcore_blockchain_height",
			Help: "Current height of the blockchain",
		}),
		Difficulty: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "coubcore_difficulty",
			Help: "Current mining difficulty",
		}),
		BlockProcessingTime: promauto.NewHistogram(prometheus.HistogramOpts{
			Name:    "coubcore_block_processing_duration_seconds",
			Help:    "Time spent processing blocks",
			Buckets: prometheus.ExponentialBuckets(0.001, 2, 15),
		}),
		TransactionProcessingTime: promauto.NewHistogram(prometheus.HistogramOpts{
			Name:    "coubcore_transaction_processing_duration_seconds",
			Help:    "Time spent processing transactions",
			Buckets: prometheus.ExponentialBuckets(0.0001, 2, 15),
		}),

		// Network metrics
		PeersConnected: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "coubcore_peers_connected",
			Help: "Number of connected peers",
		}),
		MessagesSentTotal: promauto.NewCounter(prometheus.CounterOpts{
			Name: "coubcore_messages_sent_total",
			Help: "Total number of messages sent",
		}),
		MessagesReceivedTotal: promauto.NewCounter(prometheus.CounterOpts{
			Name: "coubcore_messages_received_total",
			Help: "Total number of messages received",
		}),
		BytesSentTotal: promauto.NewCounter(prometheus.CounterOpts{
			Name: "coubcore_bytes_sent_total",
			Help: "Total number of bytes sent",
		}),
		BytesReceivedTotal: promauto.NewCounter(prometheus.CounterOpts{
			Name: "coubcore_bytes_received_total",
			Help: "Total number of bytes received",
		}),

		// API metrics
		APIRequestsTotal: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "coubcore_api_requests_total",
			Help: "Total number of API requests",
		}, []string{"method", "endpoint", "status_code"}),
		APIRequestDuration: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "coubcore_api_request_duration_seconds",
			Help:    "API request duration in seconds",
			Buckets: prometheus.ExponentialBuckets(0.001, 2, 15),
		}, []string{"method", "endpoint"}),
		APIErrorsTotal: promauto.NewCounter(prometheus.CounterOpts{
			Name: "coubcore_api_errors_total",
			Help: "Total number of API errors",
		}),

		// Wallet metrics
		WalletTransactionsTotal: promauto.NewCounter(prometheus.CounterOpts{
			Name: "coubcore_wallet_transactions_total",
			Help: "Total number of wallet transactions",
		}),
		WalletBalance: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "coubcore_wallet_balance",
			Help: "Current wallet balance",
		}),
	}
}

// RecordBlockProcessingTime records the time taken to process a block
func (m *Metrics) RecordBlockProcessingTime(duration float64) {
	m.BlockProcessingTime.Observe(duration)
}

// RecordTransactionProcessingTime records the time taken to process a transaction
func (m *Metrics) RecordTransactionProcessingTime(duration float64) {
	m.TransactionProcessingTime.Observe(duration)
}

// RecordAPIRequest records an API request
func (m *Metrics) RecordAPIRequest(method, endpoint, statusCode string, duration float64) {
	m.APIRequestsTotal.WithLabelValues(method, endpoint, statusCode).Inc()
	m.APIRequestDuration.WithLabelValues(method, endpoint).Observe(duration)
}

// RecordAPIError records an API error
func (m *Metrics) RecordAPIError() {
	m.APIErrorsTotal.Inc()
}

// SetBlockchainHeight sets the current blockchain height
func (m *Metrics) SetBlockchainHeight(height float64) {
	m.BlockchainHeight.Set(height)
}

// SetDifficulty sets the current mining difficulty
func (m *Metrics) SetDifficulty(difficulty float64) {
	m.Difficulty.Set(difficulty)
}

// SetPeersConnected sets the number of connected peers
func (m *Metrics) SetPeersConnected(count float64) {
	m.PeersConnected.Set(count)
}

// SetWalletBalance sets the current wallet balance
func (m *Metrics) SetWalletBalance(balance float64) {
	m.WalletBalance.Set(balance)
}
