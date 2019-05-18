package factory

import (
	"github.com/aliforever/gista/errs"
	storage_handler "github.com/aliforever/gista/settings/storage-handler"
	storage_file "github.com/aliforever/gista/settings/storage/storage-file"
)

type Factory struct {
}

func (f *Factory) GetUserConfig(settingsName string, storageConfig map[string]string) *string {
	// To-Do
	//defaultConfig := "storage-file"
	return nil
}

func CreateHandler(storageConfig *map[string]string) (sh *storage_handler.StorageHandler, err error) {
	f := Factory{}
	if storageConfig == nil {
		st := "file"
		storage := &st
		storageConfig = &map[string]string{"storage": *storage}
	} else if _, ok := (*storageConfig)["storage"]; !ok {
		storage := f.GetUserConfig("storage", *storageConfig)
		if storage == nil {
			st := "file"
			storage = &st
		}
		storageConfig = &map[string]string{"storage": *storage}
	}
	locationConfig := map[string]string{}
	var sInstance storage_handler.StorageInterface
	switch (*storageConfig)["storage"] {
	case "file":
		baseFolder := f.GetUserConfig("basefolder", *storageConfig)
		bF := ""
		if baseFolder != nil {
			bF = *baseFolder
		}
		locationConfig = map[string]string{"baseFolder": bF}
		sInstance = &storage_file.File{}
	default:
		err = errs.UnknownSettingsStorageType((*storageConfig)["storage"])
	}
	sh, err = storage_handler.NewStorageHandler(sInstance, locationConfig)
	return
}
