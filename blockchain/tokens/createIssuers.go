package tokens

import (
	"log"

	"github.com/GeldNetworkMVP/GeldMVPBackend/commons"
	"github.com/GeldNetworkMVP/GeldMVPBackend/constants"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/txnbuild"
)

func CreateIssuer() (string, []byte, error) {

	pair, err := keypair.Random()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	// Token issuer keys
	tokenIssuerPK := pair.Address()
	tokenIssuerSK := pair.Seed()

	//Geld Main Account
	mainPK := commons.GoDotEnvVariable("STELLARPK")
	mainSK := commons.GoDotEnvVariable("STELLARSK")
	request := horizonclient.AccountRequest{AccountID: mainPK}
	issuerAccount, err := commons.GetHorizonClient().AccountDetail(request)
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	issuerSign, err := keypair.ParseFull(mainSK)
	if err != nil {
		return "", nil, err
	}
	CreateAccount := []txnbuild.Operation{
		&txnbuild.CreateAccount{
			Destination:   tokenIssuerPK,
			Amount:        "10",
			SourceAccount: mainPK,
		},
	}
	tx, err := txnbuild.NewTransaction(txnbuild.TransactionParams{
		SourceAccount:        &issuerAccount,
		IncrementSequenceNum: true,
		Operations:           CreateAccount,
		BaseFee:              constants.MinBaseFee,
		Memo:                 nil,
		Preconditions:        txnbuild.Preconditions{TimeBounds: constants.TransactionTimeOut},
	})
	if err != nil {
		log.Println("Error when build transaction : ", err)
	}

	signedTx, err := tx.Sign(commons.GetStellarNetwork(), issuerSign)
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	// submit transaction
	respn, err := commons.GetHorizonClient().SubmitTransaction(signedTx)
	if err != nil {
		log.Println("Error submitting transaction:", err)
	}
	// encrypt the issuer secret key
	encryptedSK := commons.Encrypt(tokenIssuerSK)

	log.Println("Transaction Hash for new Account creation: ", respn.Hash)
	return tokenIssuerPK, encryptedSK, err
}
