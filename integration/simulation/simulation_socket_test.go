package simulation

import (
	"testing"
	"time"

	"github.com/obscuronet/obscuro-playground/integration"

	"github.com/obscuronet/obscuro-playground/integration/datagenerator"

	ethereum_mock "github.com/obscuronet/obscuro-playground/integration/ethereummock"

	"github.com/obscuronet/obscuro-playground/integration/simulation/params"

	"github.com/obscuronet/obscuro-playground/integration/simulation/network"
)

// This test creates a network of L2 nodes, then injects transactions, and finally checks the resulting output blockchain
// The L2 nodes communicate with each other via sockets, and with their enclave servers via RPC.
// All nodes and enclaves live in the same process, and the Ethereum nodes are mocked out.
func TestSocketNodesMonteCarloSimulation(t *testing.T) {
	setupTestLog("socket")

	simParams := &params.SimParams{
		NumberOfNodes:             7,
		AvgBlockDuration:          250 * time.Millisecond,
		SimulationTime:            25 * time.Second,
		L1EfficiencyThreshold:     0.2,
		L2EfficiencyThreshold:     0.33,
		L2ToL1EfficiencyThreshold: 0.4,
		MgmtContractLib:           ethereum_mock.NewMgmtContractLibMock(),
		ERC20ContractLib:          ethereum_mock.NewERC20ContractLibMock(),
		StartPort:                 integration.StartPortSimulationSocket,
	}
	simParams.AvgNetworkLatency = simParams.AvgBlockDuration / 15
	simParams.AvgGossipPeriod = simParams.AvgBlockDuration / 4

	for i := 0; i < simParams.NumberOfNodes+1; i++ {
		simParams.NodeEthWallets = append(simParams.NodeEthWallets, datagenerator.RandomWallet(integration.EthereumChainID))
		simParams.SimEthWallets = append(simParams.SimEthWallets, datagenerator.RandomWallet(integration.EthereumChainID))
	}

	testSimulation(t, network.NewNetworkOfSocketNodes(), simParams)
}
