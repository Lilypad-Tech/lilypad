package collectpowsubmission

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/pow"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type PowValidPOWSubmitted struct {
	Id                string `bun:"id,pk"`
	Challenge         string `bun:"challenge"`
	WalletAddress     string `bun:"wallet_address,"`
	NodeId            string `bun:"node_id"`
	Nonce             string `bun:"nonce"`
	StartTimestamp    int64  `bun:"start_timestamp"`
	CompleteTimestamp int64  `bun:"complete_timestamp"`
	Difficulty        string `bun:"difficulty"`
}

func StartCollectPowSubmission(ctx context.Context, we3Sdk *web3.Web3SDK, pgConnectionString string, cm *system.CleanupManager) error {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(pgConnectionString)))
	db := bun.NewDB(sqldb, pgdialect.New())

	_, err := db.NewCreateTable().IfNotExists().Model((*PowValidPOWSubmitted)(nil)).Exec(ctx)
	if err != nil {
		return err
	}

	_, err = db.NewCreateIndex().Model((*PowValidPOWSubmitted)(nil)).IfNotExists().Index("address_index").Column("wallet_address").Exec(ctx)
	if err != nil {
		return err
	}

	log.Info().Msg("Start to fix missing pow submission event, this may take some time")
	err = fixMissingPowSubmissionEvent(ctx, we3Sdk, db)
	if err != nil {
		return err
	}
	log.Info().Msg("Finish fix missing pow submission event")

	channel := web3.NewValidPOWSubmittedEventChannels()
	go channel.Start(we3Sdk, ctx, cm)

	channel.SubscribePowValidPOWSubmitted(func(pvp pow.PowValidPOWSubmitted) {
		id := submissionId(pvp.WalletAddress, pvp.Challenge, pvp.NodeId)
		dbModel := PowValidPOWSubmitted{
			Id:                id,
			WalletAddress:     pvp.WalletAddress.String(),
			NodeId:            pvp.NodeId,
			Nonce:             pvp.Nonce.String(),
			StartTimestamp:    pvp.StartTimestamp.Int64(),
			CompleteTimestamp: pvp.CompleteTimestamp.Int64(),
			Challenge:         hex.EncodeToString(pvp.Challenge[:]),
			Difficulty:        pvp.Difficulty.String(),
		}

		//todo 1 write log to efs?
		//todo 2 need another command to check no event missing
		_, err = db.NewInsert().Model(&dbModel).Ignore().Exec(ctx)
		if err != nil {
			log.Err(err).Msgf("Unable to insert data to database")
			return
		}
		log.Info().
			Str("address", dbModel.WalletAddress).
			Str("nodeId", dbModel.NodeId).
			Str("nonce", dbModel.Nonce).
			Int64("start_timestamp", dbModel.StartTimestamp).
			Int64("complete_timestamp", dbModel.CompleteTimestamp).
			Str("challenge", dbModel.Challenge).
			Str("difficulty", dbModel.Difficulty).
			Msgf("Receive a valid pow submission")
	})

	<-ctx.Done()
	return nil
}

func fixMissingPowSubmissionEvent(ctx context.Context, we3Sdk *web3.Web3SDK, db *bun.DB) error {
	bf := bloom.NewWithEstimates(1000000, 0.0001) //arbitary parameter, need to update if not enough
	pageSize := 100
	offset := 0

	//collect all challenges and store the in bloomfilter
	for {
		var ids []string
		err := db.NewSelect().Model((*PowValidPOWSubmitted)(nil)).Column("id").Offset(offset).Limit(pageSize).Scan(ctx, &ids)
		if err != nil {
			return err
		}
		for _, id := range ids {
			pk, err := hex.DecodeString(id)
			if err != nil {
				return err
			}
			bf.Add(pk)
		}

		if len(ids) < pageSize {
			break
		}

		offset += pageSize
	}

	// Get submissions from sc and check them with bloomfilter and database. if missing, record it
	addrs, err := we3Sdk.Contracts.Pow.GetMiners(we3Sdk.CallOpts)
	if err != nil {
		return err
	}
	for _, addr := range addrs {
		submissions, err := we3Sdk.Contracts.Pow.GetMinerPosSubmissions(we3Sdk.CallOpts, addr)
		if err != nil {
			return err
		}
		for _, submission := range submissions {
			id := submissionId(submission.WalletAddress, submission.Challenge, submission.NodeId)
			count, err := db.NewSelect().Model((*PowValidPOWSubmitted)(nil)).Where("id = ?", id).Count(ctx)
			if err != nil {
				return err
			}
			if count == 0 {
				//missing value
				dbModel := PowValidPOWSubmitted{
					Id:                id,
					WalletAddress:     submission.WalletAddress.String(),
					NodeId:            submission.NodeId,
					Nonce:             submission.Nonce.String(),
					StartTimestamp:    submission.StartTimestamp.Int64(),
					CompleteTimestamp: submission.CompleteTimestamp.Int64(),
					Challenge:         hex.EncodeToString(submission.Challenge[:]),
					Difficulty:        submission.Difficulty.String(),
				}
				_, err = db.NewInsert().Model(&dbModel).Ignore().Exec(ctx)
				if err != nil {
					return nil
				}
				bf.Add(submission.Challenge[:])
				log.Info().
					Str("address", dbModel.WalletAddress).
					Str("nodeId", dbModel.NodeId).
					Str("nonce", dbModel.Nonce).
					Int64("start_timestamp", dbModel.StartTimestamp).
					Int64("complete_timestamp", dbModel.CompleteTimestamp).
					Str("challenge", dbModel.Challenge).
					Str("difficulty", dbModel.Difficulty).
					Msgf("Fix a missing value")
			} else {
				log.Debug().Msg("A value missed in bloomfilter, consider to increase bloomfilter argument")
			}
		}
	}
	return nil
}

func submissionId(addr common.Address, challenge [32]byte, nodeID string) string {
	hasher := sha256.New()
	hasher.Write(addr.Bytes())
	hasher.Write(challenge[:])
	hasher.Write([]byte(nodeID))
	return hex.EncodeToString(hasher.Sum(nil))
}
