package http

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/anoideaopen/channel-transfer/pkg/model"
	clihttp "github.com/anoideaopen/channel-transfer/test/integration/clihttp/client"
	"github.com/anoideaopen/channel-transfer/test/integration/clihttp/client/transfer"
	"github.com/anoideaopen/channel-transfer/test/integration/clihttp/models"
	pbfound "github.com/anoideaopen/foundation/proto"
	"github.com/anoideaopen/foundation/test/integration/cmn"
	"github.com/anoideaopen/foundation/test/integration/cmn/client"
	"github.com/btcsuite/btcutil/base58"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric/integration"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/metadata"
)

const (
	fnIndustrialBalanceOf            = "industrialBalanceOf"
	fnAllowedBalanceOf               = "allowedBalanceOf"
	fnChannelMultiTransferByAdmin    = "channelMultiTransferByAdmin"
	fnChannelMultiTransferByCustomer = "channelMultiTransferByCustomer"
)

var _ = Describe("Channel multi transfer HTTP tests", func() {
	var (
		ts client.TestSuite
	)

	BeforeEach(func() {
		ts = client.NewTestSuite(components)
	})
	AfterEach(func() {
		ts.ShutdownNetwork()
	})

	var (
		channels = []string{cmn.ChannelAcl, cmn.ChannelCC, cmn.ChannelIndustrial}
		user     *client.UserFoundation

		networkFound *cmn.NetworkFoundation

		clientCtx   context.Context
		transferCli *clihttp.CrossChanelTransfer
		auth        runtime.ClientAuthInfoWriter

		transferItems              []model.TransferItem
		initialBalances            []model.TransferItem
		expectedIndustrialBalances []model.TransferItem
	)
	BeforeEach(func() {
		By("start redis")
		ts.StartRedis()
	})
	BeforeEach(func() {
		ts.InitNetwork(channels, integration.GossipBasePort)
		ts.DeployChaincodes()
	})
	BeforeEach(func() {
		By("start robot")
		ts.StartRobot()
		By("start channel transfer")
		ts.StartChannelTransfer()
	})
	AfterEach(func() {
		By("stop robot")
		ts.StopRobot()
		By("stop channel transfer")
		ts.StopChannelTransfer()
		By("stop redis")
		ts.StopRedis()
	})

	BeforeEach(func() {
		networkFound = ts.NetworkFound()

		By("add admin to acl")
		ts.AddAdminToACL()

		By("add user to acl")
		var err error
		user, err = client.NewUserFoundation(pbfound.KeyType_ed25519)
		Expect(err).NotTo(HaveOccurred())

		ts.AddUser(user)

		By("emit tokens")
		ts.NBTxInvokeWithSign(
			cmn.ChannelIndustrial,
			cmn.ChannelIndustrial,
			ts.Admin(),
			"initialize",
			"",
			client.NewNonceByTime().Get(),
		).CheckErrorIsNil()

		// Here we use model.TransferItem to preserve the order of the fields
		// when marshalling to json string to sign the transaction arguments.
		// When sending a request to the http server, these structures will be
		// mapped to models.ChannelTransferTransferItem.
		// Using models.ChannelTransferTransferItem without mapping is very risky,
		// because this structure was generated by go-swagger tools, and we have
		// no control over the order of fields in it. When making http-requests,
		// the client should make sure that the order of fields is correct:
		// first Token, then Amount.
		initialBalances = []model.TransferItem{
			{
				Token:  "INDUSTRIAL_202009",
				Amount: "10000000000000",
			},
			{
				Token:  "INDUSTRIAL_202010",
				Amount: "100000000000000",
			},
			{
				Token:  "INDUSTRIAL_202011",
				Amount: "200000000000000",
			},
			{
				Token:  "INDUSTRIAL_202012",
				Amount: "50000000000000",
			},
		}

		transferItems = []model.TransferItem{
			{
				Token:  "INDUSTRIAL_202009",
				Amount: "1000000000000",
			},
			{
				Token:  "INDUSTRIAL_202010",
				Amount: "10000000000000",
			},
			{
				Token:  "INDUSTRIAL_202011",
				Amount: "20000000000000",
			},
			{
				Token:  "INDUSTRIAL_202012",
				Amount: "5000000000000",
			},
		}

		expectedIndustrialBalances = []model.TransferItem{
			{
				Token:  "INDUSTRIAL_202009",
				Amount: "9000000000000",
			},
			{
				Token:  "INDUSTRIAL_202010",
				Amount: "90000000000000",
			},
			{
				Token:  "INDUSTRIAL_202011",
				Amount: "180000000000000",
			},
			{
				Token:  "INDUSTRIAL_202012",
				Amount: "45000000000000",
			},
		}

		for _, initial := range initialBalances {
			group := strings.Split(initial.Token, "_")[1]

			ts.Query(
				cmn.ChannelIndustrial,
				cmn.ChannelIndustrial,
				fnIndustrialBalanceOf,
				ts.Admin().AddressBase58Check,
			).CheckIndustrialBalance(group, initial.Amount)

			ts.TxInvokeWithSign(
				cmn.ChannelIndustrial,
				cmn.ChannelIndustrial,
				ts.Admin(),
				"transferIndustrial",
				"",
				client.NewNonceByTime().Get(),
				user.AddressBase58Check,
				group,
				initial.Amount,
				"transfer industrial tokens",
			).CheckErrorIsNil()

			ts.Query(
				cmn.ChannelIndustrial,
				cmn.ChannelIndustrial,
				fnIndustrialBalanceOf,
				user.AddressBase58Check,
			).CheckIndustrialBalance(group, initial.Amount)
		}

		By("creating http connection")
		clientCtx = metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", networkFound.ChannelTransfer.AccessToken))

		httpAddress := networkFound.ChannelTransfer.HostAddress + ":" + strconv.FormatUint(uint64(networkFound.ChannelTransfer.Ports[cmn.HttpPort]), 10)
		transport := httptransport.New(httpAddress, "", nil)
		transferCli = clihttp.New(transport, strfmt.Default)

		auth = httptransport.APIKeyAuth("authorization", "header", networkFound.ChannelTransfer.AccessToken)
	})

	It("multi transfer by admin test", func() {
		authOpts := func(c *runtime.ClientOperation) {
			c.AuthInfo = auth
		}

		By("creating channel transfer request")
		items, err := json.Marshal(transferItems)
		Expect(err).NotTo(HaveOccurred())

		transferID := uuid.NewString()
		channelTransferArgs := []string{transferID, ccCCUpper, user.AddressBase58Check, string(items)}

		requestID := uuid.NewString()
		nonce := client.NewNonceByTime().Get()
		signArgs := append(append([]string{fnChannelMultiTransferByAdmin, requestID, cmn.ChannelIndustrial, cmn.ChannelIndustrial}, channelTransferArgs...), nonce)
		publicKey, sign, err := ts.Admin().Sign(signArgs...)
		Expect(err).NotTo(HaveOccurred())

		transferRequest := &models.ChannelTransferMultiTransferBeginAdminRequest{
			Generals: &models.ChannelTransferGeneralParams{
				MethodName: fnChannelMultiTransferByAdmin,
				RequestID:  requestID,
				Chaincode:  cmn.ChannelIndustrial,
				Channel:    cmn.ChannelIndustrial,
				Nonce:      nonce,
				PublicKey:  publicKey,
				Sign:       base58.Encode(sign),
			},
			IDTransfer: channelTransferArgs[0],
			ChannelTo:  channelTransferArgs[1],
			Address:    channelTransferArgs[2],
			Items:      mapTransferItems(transferItems),
		}

		By("sending transfer request")
		res, err := transferCli.Transfer.MultiTransferByAdmin(&transfer.MultiTransferByAdminParams{Body: transferRequest, Context: clientCtx}, authOpts)
		Expect(err).NotTo(HaveOccurred())

		err = checkResponseStatus(res.GetPayload(), models.ChannelTransferTransferStatusResponseStatusSTATUSINPROCESS, "")
		Expect(err).NotTo(HaveOccurred())

		By("awaiting for channel transfer to respond")
		err = waitForAnswerAndCheckStatus(clientCtx, transferCli, transferID, authOpts, models.ChannelTransferTransferStatusResponseStatusSTATUSCOMPLETED, "")
		Expect(err).NotTo(HaveOccurred())

		By("checking result balances")
		for i, expected := range expectedIndustrialBalances {
			group := strings.Split(expected.Token, "_")[1]

			ts.Query(
				cmn.ChannelIndustrial,
				cmn.ChannelIndustrial,
				fnIndustrialBalanceOf,
				user.AddressBase58Check,
			).CheckIndustrialBalance(group, expected.Amount)

			ts.Query(
				cmn.ChannelCC,
				cmn.ChannelCC,
				fnAllowedBalanceOf,
				user.AddressBase58Check,
				transferItems[i].Token,
			).CheckBalance(transferItems[i].Amount)
		}
	})

	It("multi transfer by customer test", func() {
		authOpts := func(c *runtime.ClientOperation) {
			c.AuthInfo = auth
		}

		By("creating channel transfer request")
		items, err := json.Marshal(transferItems)
		Expect(err).NotTo(HaveOccurred())

		transferID := uuid.NewString()
		channelTransferArgs := []string{transferID, ccCCUpper, string(items)}

		requestID := uuid.NewString()
		nonce := client.NewNonceByTime().Get()
		signArgs := append(append([]string{fnChannelMultiTransferByCustomer, requestID, cmn.ChannelIndustrial, cmn.ChannelIndustrial}, channelTransferArgs...), nonce)
		publicKey, sign, err := user.Sign(signArgs...)
		Expect(err).NotTo(HaveOccurred())

		transferRequest := &models.ChannelTransferMultiTransferBeginCustomerRequest{
			Generals: &models.ChannelTransferGeneralParams{
				MethodName: fnChannelMultiTransferByCustomer,
				RequestID:  requestID,
				Chaincode:  cmn.ChannelIndustrial,
				Channel:    cmn.ChannelIndustrial,
				Nonce:      nonce,
				PublicKey:  publicKey,
				Sign:       base58.Encode(sign),
			},
			IDTransfer: channelTransferArgs[0],
			ChannelTo:  channelTransferArgs[1],
			Items:      mapTransferItems(transferItems),
		}

		By("sending transfer request")
		res, err := transferCli.Transfer.MultiTransferByCustomer(&transfer.MultiTransferByCustomerParams{Body: transferRequest, Context: clientCtx}, authOpts)
		Expect(err).NotTo(HaveOccurred())

		err = checkResponseStatus(res.GetPayload(), models.ChannelTransferTransferStatusResponseStatusSTATUSINPROCESS, "")
		Expect(err).NotTo(HaveOccurred())

		By("awaiting for channel transfer to respond")
		err = waitForAnswerAndCheckStatus(clientCtx, transferCli, transferID, authOpts, models.ChannelTransferTransferStatusResponseStatusSTATUSCOMPLETED, "")
		Expect(err).NotTo(HaveOccurred())

		By("checking result balances")
		for i, expected := range expectedIndustrialBalances {
			group := strings.Split(expected.Token, "_")[1]

			ts.Query(
				cmn.ChannelIndustrial,
				cmn.ChannelIndustrial,
				fnIndustrialBalanceOf,
				user.AddressBase58Check,
			).CheckIndustrialBalance(group, expected.Amount)

			ts.Query(
				cmn.ChannelCC,
				cmn.ChannelCC,
				fnAllowedBalanceOf,
				user.AddressBase58Check,
				transferItems[i].Token,
			).CheckBalance(transferItems[i].Amount)
		}
	})
})

func mapTransferItems(transferItems []model.TransferItem) []*models.ChannelTransferTransferItem {
	mappedTransferItems := make([]*models.ChannelTransferTransferItem, len(transferItems), len(transferItems))
	for i, transferItem := range transferItems {
		mappedTransferItems[i] = &models.ChannelTransferTransferItem{
			Token:  transferItem.Token,
			Amount: transferItem.Amount,
		}
	}
	return mappedTransferItems
}
