package tokens

import (
	"log"
	"strconv"

	"github.com/GeldNetworkMVP/GeldMVPBackend/businessFacade"
	"github.com/GeldNetworkMVP/GeldMVPBackend/commons"
	"github.com/GeldNetworkMVP/GeldMVPBackend/constants"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/txnbuild"
	"github.com/stellar/go/xdr"
)

func PutTokenOnSaleToken(TokenIssuerPK string, assetcode string, cid string, htmlhash string, price string) (string, error) {
	data, err := businessFacade.GetStellarKeys(TokenIssuerPK)
	if err != nil {
		log.Println(err)
		return "", err
	} else {
		issuerSK := data.SK
		DecryptedSK := commons.Decrypt(issuerSK)

		request := horizonclient.AccountRequest{AccountID: TokenIssuerPK}
		issuerAccount, err := commons.GetHorizonClient().AccountDetail(request)
		if err != nil {
			return "", err
		}
		issuerSign, err := keypair.ParseFull(DecryptedSK)
		if err != nil {
			return "", err
		}

		var homeDomain string = commons.GoDotEnvVariable("HOMEDOMAIN")
		var distributerPK string = commons.GoDotEnvVariable("STELLARPK")
		var distributorSK string = commons.GoDotEnvVariable("STELLARSK")

		mainSign, err := keypair.ParseFull(distributorSK)
		if err != nil {
			return "", err
		}
		var cotenthash = []byte(htmlhash)
		var contentCID = []byte(cid)
		unitprice, err := (strconv.Atoi(price))
		unitpriceInt := int32(unitprice)

		payments := []txnbuild.Operation{
			&txnbuild.ChangeTrust{
				Line: txnbuild.ChangeTrustAssetWrapper{
					Asset: txnbuild.CreditAsset{Code: assetcode,
						Issuer: TokenIssuerPK},
				},
				Limit:         "1",
				SourceAccount: distributerPK,
			},
			&txnbuild.Payment{Destination: distributerPK, Asset: txnbuild.CreditAsset{Code: assetcode,
				Issuer: TokenIssuerPK}, Amount: "1"},
			&txnbuild.ManageData{
				Name:  assetcode,
				Value: cotenthash,
			},
			&txnbuild.ManageData{
				Name:  "CID",
				Value: contentCID,
			},
			&txnbuild.ManageSellOffer{
				Selling: txnbuild.CreditAsset{
					Code:   assetcode,
					Issuer: TokenIssuerPK,
				},
				Buying: txnbuild.NativeAsset{},
				Amount: "1",
				Price: xdr.Price{
					N: xdr.Int32(unitpriceInt),
					D: 1,
				},
				OfferID:       0,
				SourceAccount: distributerPK,
			},
			&txnbuild.SetOptions{
				MasterWeight: txnbuild.NewThreshold(0),
				HomeDomain:   &homeDomain,
			},
		}
		tx, err := txnbuild.NewTransaction(
			txnbuild.TransactionParams{
				SourceAccount:        &issuerAccount,
				IncrementSequenceNum: true,
				Operations:           payments,
				BaseFee:              constants.MinBaseFee,
				Memo:                 nil,
				Preconditions:        txnbuild.Preconditions{TimeBounds: constants.TransactionTimeOut},
			},
		)
		if err != nil {
			log.Println(err)
			return "", err
		}

		signedTx, err := tx.Sign(commons.GetStellarNetwork(), issuerSign, mainSign)
		if err != nil {
			return "", err
		}
		// submit transaction
		respn, err := commons.GetHorizonClient().SubmitTransaction(signedTx)
		if err != nil {
			log.Println("Error submitting transaction:", err)
		}
		return respn.Hash, nil
	}

}
