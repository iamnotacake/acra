// Copyright 2016, Cossack Labs Limited
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/cossacklabs/acra/keystore"
	"github.com/cossacklabs/acra/poison"
	"github.com/cossacklabs/acra/utils"
	"github.com/vharitonsky/iniflags"
	"os"
)

var DEFAULT_CONFIG_PATH = utils.GetConfigPathByName("acra_genpoisonrecord")

func main() {
	keys_dir := flag.String("keys_dir", keystore.DEFAULT_KEY_DIR_SHORT, "Folder from which will be loaded keys")
	data_length := flag.Int("data_length", poison.DEFAULT_DATA_LENGTH, fmt.Sprintf("Length of random data for data block in acrastruct. -1 is random in range 1..%v", poison.MAX_DATA_LENGTH))

	utils.LoadFromConfig(DEFAULT_CONFIG_PATH)
	iniflags.Parse()

	store, err := keystore.NewFilesystemKeyStore(*keys_dir)
	if err != nil {
		fmt.Printf("Error: %v\n", utils.ErrorMessage("can't initialize key store", err))
		os.Exit(1)
	}
	poison_record, err := poison.CreatePoisonRecord(store, *data_length)
	if err != nil {
		fmt.Printf("Error: %v\n", utils.ErrorMessage("can't create poison record", err))
	}
	fmt.Println(base64.StdEncoding.EncodeToString(poison_record))
}
