package blockchain

import (
	"log"

	"github.com/GeldNetworkMVP/GeldMVPBackend/commons"
	"github.com/sirupsen/logrus"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/txnbuild"
)

/*
ActivateAccount
@desc - Activate account for newly created keys
@params - None
*/
func ActivateAccount(publickey string) (string, error) {
	client := commons.GetHorizonClient()
	mainPK := commons.GoDotEnvVariable("STELLARPK")
	mainSK := commons.GoDotEnvVariable("STELLARSK")
	accountRequest := horizonclient.AccountRequest{AccountID: mainPK}
	sourceAccount, err := client.AccountDetail(accountRequest)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	mainKeypair, _ := keypair.ParseFull(mainSK)
	CreateAccount := []txnbuild.Operation{
		&txnbuild.CreateAccount{
			Destination:   publickey,
			Amount:        "10",
			SourceAccount: mainPK,
		},
	}
	tx, err := txnbuild.NewTransaction(txnbuild.TransactionParams{
		SourceAccount:        &sourceAccount,
		IncrementSequenceNum: true,
		Operations:           CreateAccount,
		BaseFee:              txnbuild.MinBaseFee,
		Memo:                 nil,
		Preconditions:        txnbuild.Preconditions{TimeBounds: txnbuild.NewInfiniteTimeout()},
	})
	if err != nil {
		log.Println("Error when build transaction : ", err)
		return "", err
	}

	signedTx, err := tx.Sign(commons.GetStellarNetwork(), mainKeypair)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	// submit transaction
	respn, err := commons.GetHorizonClient().SubmitTransaction(signedTx)
	if err != nil {
		log.Println("Error submitting transaction:", err)
		return "", err
	}

	log.Println("Transaction Hash for new Account creation: ", respn.Hash)
	return respn.Hash, err
}

func FundAccount(publickey string) (string, error) {
	client := commons.GetHorizonClient()
	mainPK := commons.GoDotEnvVariable("STELLARPK")
	mainSK := commons.GoDotEnvVariable("STELLARSK")
	accountRequest := horizonclient.AccountRequest{AccountID: mainPK}
	sourceAccount, err := client.AccountDetail(accountRequest)
	if err != nil {
		log.Fatal(err)
	}
	mainKeypair, _ := keypair.ParseFull(mainSK)
	Payment := []txnbuild.Operation{
		&txnbuild.Payment{
			Destination:   publickey,
			Amount:        "10",
			Asset:         txnbuild.NativeAsset{},
			SourceAccount: mainPK,
		},
	}
	tx, err := txnbuild.NewTransaction(txnbuild.TransactionParams{
		SourceAccount:        &sourceAccount,
		IncrementSequenceNum: true,
		Operations:           Payment,
		BaseFee:              txnbuild.MinBaseFee,
		Memo:                 nil,
		Preconditions:        txnbuild.Preconditions{TimeBounds: txnbuild.NewInfiniteTimeout()},
	})
	if err != nil {
		log.Println("Error when build transaction : ", err)
	}

	signedTx, err := tx.Sign(commons.GetStellarNetwork(), mainKeypair)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	// submit transaction
	respn, err := commons.GetHorizonClient().SubmitTransaction(signedTx)
	if err != nil {
		log.Println("Error submitting transaction:", err)
	}

	log.Println("Transaction Hash for funding Accounts: ", respn.Hash)
	return respn.Hash, err
}
