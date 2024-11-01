package migration

import (
	"context"

	"github.com/funmi4194/instashop/database"
	"github.com/opensaucerer/barf"
)

var Table = []interface{}{
	// &repository.Factory{},
	// &userRepository.User{},
	// &txRepository.Transaction{},
	// &txRepository.Wallet{},
	// &txRepository.Card{},
	// &socialRepository.Social{},
	// &settingRepository.Setting{},
	// &settingRepository.PlatformSetting{},
	// &settingRepository.AssetCategory{},
	// &settingRepository.AssetType{},
	// &payoutRepository.Payout{},
	// &creationRepository.Asset{},
	// &creationRepository.Collection{},
	// &creationRepository.CollectionAsset{},
	// &creationRepository.CartItem{},
	// &creationRepository.Print{},
	// &creationRepository.Story{},
}

// CreateTables creates tables that do not already exist. Although we have connections to other DBs configure.Save should only handle migration for configure.Save DB.
func CreateTables() error {
	for _, m := range Table {
		_, err := database.PostgreSQLDB.NewCreateTable().
			IfNotExists().
			Model(m).Exec(context.TODO())
		if err != nil {
			barf.Logger().Warnf("failed to create %v table", m)
			return err
		}
	}
	return nil
}

// migrate effects any database schema migration
func Migrate() error {
	return nil
}
