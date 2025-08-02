package params

import (
	"encoding/json"
	"os"

	pkgerrors "github.com/pkg/errors"

	wakutypes "github.com/status-im/status-go/waku/types"
)

// Define available fleets.
const (
	FleetUndefined     = ""
	FleetProd          = "eth.prod"
	FleetStatusStaging = "status.staging"
	FleetStatusProd    = "status.prod"
	FleetWakuSandbox   = "waku.sandbox"
	FleetWakuTest      = "waku.test"
)

type FleetInfo struct {
	WakuNodes            []string               `json:"wakuNodes"`
	DiscV5BootstrapNodes []string               `json:"discV5BootstrapNodes"`
	ClusterID            uint16                 `json:"clusterID"`
	StoreNodes           []wakutypes.Mailserver `json:"storeNodes"`
}
type FleetsMap map[string]FleetInfo

// DefaultWakuNodes is a list of "supported" fleets. This list is populated to clients UI settings.
var supportedFleets = FleetsMap{
	FleetStatusStaging: {
		ClusterID: 72,
		WakuNodes: []string{
			"enrtree://AIATD2HQJDPFVQ2SJB7HSNWBHQ3SEC3XCENFZYX4JMCPHFTEMDUQC@boot.ionalabs.ai",
		},
		DiscV5BootstrapNodes: []string{
			"enrtree://AIATD2HQJDPFVQ2SJB7HSNWBHQ3SEC3XCENFZYX4JMCPHFTEMDUQC@boot.ionalabs.ai",
			"enr:-L24QOcdmsEln9seXi4gMXW59Pi9nP9nG1dcj2TJapaKMZc6TEpLNumnZVGYRM_W9fQ6ZKXCCFf6DNspVxl8aGtUykkCgmlkgnY0gmlwhLM99gSKbXVsdGlhZGRyc5YACASzPfYEBiMxAAoEsz32BAYjMt0DgnJzhQBIAQAgiXNlY3AyNTZrMaECsq2h0_gpY49mMm6tp3j9awr_xGCbxjoFXHp0m2aG8MKDdGNwgiMxg3VkcIIjM4V3YWt1Mg8",
			"enr:-L24QKLgpWYrASXFkCHVq08LkXBjmI7G6VlrIzr1wVxOnXISDVQPNoQSxwizTD-ypogX6TYEx1auvEsotqlRxKyYFGoCgmlkgnY0gmlwhMMjAG2KbXVsdGlhZGRyc5YACATDIwBtBiMxAAoEwyMAbQYjMt0DgnJzhQBIAQAgiXNlY3AyNTZrMaEC1d2IxoXQ58u6XuypxWJoevt4T6Gyjz0x2VEEK3OjzF2DdGNwgiMxg3VkcIIjM4V3YWt1Mg8",
		},
		StoreNodes: []wakutypes.Mailserver{
			{
				ID:    "store-01.iona.status.staging",
				ENR:   wakutypes.MustDecodeENR("enr:-L24QOcdmsEln9seXi4gMXW59Pi9nP9nG1dcj2TJapaKMZc6TEpLNumnZVGYRM_W9fQ6ZKXCCFf6DNspVxl8aGtUykkCgmlkgnY0gmlwhLM99gSKbXVsdGlhZGRyc5YACASzPfYEBiMxAAoEsz32BAYjMt0DgnJzhQBIAQAgiXNlY3AyNTZrMaECsq2h0_gpY49mMm6tp3j9awr_xGCbxjoFXHp0m2aG8MKDdGNwgiMxg3VkcIIjM4V3YWt1Mg8"),
				Addr:  wakutypes.MustDecodeMultiaddress("/ip4/179.61.246.4/tcp/9009/p2p/16Uiu2HAm7TA7Qq13C7DkXno4eFZGpxbm3thJasvZpgP6SBGSQQJd"),
				Fleet: FleetStatusStaging,
			},
			{
				ID:    "store-02.iona.status.staging",
				ENR:   wakutypes.MustDecodeENR("enr:-L24QKLgpWYrASXFkCHVq08LkXBjmI7G6VlrIzr1wVxOnXISDVQPNoQSxwizTD-ypogX6TYEx1auvEsotqlRxKyYFGoCgmlkgnY0gmlwhMMjAG2KbXVsdGlhZGRyc5YACATDIwBtBiMxAAoEwyMAbQYjMt0DgnJzhQBIAQAgiXNlY3AyNTZrMaEC1d2IxoXQ58u6XuypxWJoevt4T6Gyjz0x2VEEK3OjzF2DdGNwgiMxg3VkcIIjM4V3YWt1Mg8"),
				Addr:  wakutypes.MustDecodeMultiaddress("/ip4/195.35.0.109/tcp/9009/p2p/16Uiu2HAm9pWkCfdapHawsYrfRFrA7FuHV9jegujTfhSFjNWmn5fA"),
				Fleet: FleetStatusStaging,
			},
		},
	},
	FleetStatusProd: {
   ClusterID: 72,
   WakuNodes: []string{
   	"enrtree://AIATD2HQJDPFVQ2SJB7HSNWBHQ3SEC3XCENFZYX4JMCPHFTEMDUQC@boot.ionalabs.ai",
   },
   DiscV5BootstrapNodes: []string{
   	"enrtree://AIATD2HQJDPFVQ2SJB7HSNWBHQ3SEC3XCENFZYX4JMCPHFTEMDUQC@boot.ionalabs.ai",
   	"enr:-L24QOcdmsEln9seXi4gMXW59Pi9nP9nG1dcj2TJapaKMZc6TEpLNumnZVGYRM_W9fQ6ZKXCCFf6DNspVxl8aGtUykkCgmlkgnY0gmlwhLM99gSKbXVsdGlhZGRyc5YACASzPfYEBiMxAAoEsz32BAYjMt0DgnJzhQBIAQAgiXNlY3AyNTZrMaECsq2h0_gpY49mMm6tp3j9awr_xGCbxjoFXHp0m2aG8MKDdGNwgiMxg3VkcIIjM4V3YWt1Mg8",
   	"enr:-L24QKLgpWYrASXFkCHVq08LkXBjmI7G6VlrIzr1wVxOnXISDVQPNoQSxwizTD-ypogX6TYEx1auvEsotqlRxKyYFGoCgmlkgnY0gmlwhMMjAG2KbXVsdGlhZGRyc5YACATDIwBtBiMxAAoEwyMAbQYjMt0DgnJzhQBIAQAgiXNlY3AyNTZrMaEC1d2IxoXQ58u6XuypxWJoevt4T6Gyjz0x2VEEK3OjzF2DdGNwgiMxg3VkcIIjM4V3YWt1Mg8",
   },
   StoreNodes: []wakutypes.Mailserver{
   	{
   		ID:    "store-01.iona.status.prod",
   		ENR:   wakutypes.MustDecodeENR("enr:-L24QOcdmsEln9seXi4gMXW59Pi9nP9nG1dcj2TJapaKMZc6TEpLNumnZVGYRM_W9fQ6ZKXCCFf6DNspVxl8aGtUykkCgmlkgnY0gmlwhLM99gSKbXVsdGlhZGRyc5YACASzPfYEBiMxAAoEsz32BAYjMt0DgnJzhQBIAQAgiXNlY3AyNTZrMaECsq2h0_gpY49mMm6tp3j9awr_xGCbxjoFXHp0m2aG8MKDdGNwgiMxg3VkcIIjM4V3YWt1Mg8"),
   		Addr:  wakutypes.MustDecodeMultiaddress("/ip4/179.61.246.4/tcp/9009/p2p/16Uiu2HAm7TA7Qq13C7DkXno4eFZGpxbm3thJasvZpgP6SBGSQQJd"),
   		Fleet: FleetStatusProd,
   	},
   	{
   		ID:    "store-02.iona.status.prod",
   		ENR:   wakutypes.MustDecodeENR("enr:-L24QKLgpWYrASXFkCHVq08LkXBjmI7G6VlrIzr1wVxOnXISDVQPNoQSxwizTD-ypogX6TYEx1auvEsotqlRxKyYFGoCgmlkgnY0gmlwhMMjAG2KbXVsdGlhZGRyc5YACATDIwBtBiMxAAoEwyMAbQYjMt0DgnJzhQBIAQAgiXNlY3AyNTZrMaEC1d2IxoXQ58u6XuypxWJoevt4T6Gyjz0x2VEEK3OjzF2DdGNwgiMxg3VkcIIjM4V3YWt1Mg8"),
   		Addr:  wakutypes.MustDecodeMultiaddress("/ip4/195.35.0.109/tcp/9009/p2p/16Uiu2HAm9pWkCfdapHawsYrfRFrA7FuHV9jegujTfhSFjNWmn5fA"),
   		Fleet: FleetStatusProd,
   	},
   },
},
	FleetWakuSandbox: {
   WakuNodes: []string{
   	"enrtree://AIATD2HQJDPFVQ2SJB7HSNWBHQ3SEC3XCENFZYX4JMCPHFTEMDUQC@boot.ionalabs.ai",
   },
   DiscV5BootstrapNodes: []string{
   	"enrtree://AIATD2HQJDPFVQ2SJB7HSNWBHQ3SEC3XCENFZYX4JMCPHFTEMDUQC@boot.ionalabs.ai",
   },
   StoreNodes: []wakutypes.Mailserver{
   	{
   		ID:    "node-01.iona.waku.sandbox",
   		ENR:   wakutypes.MustDecodeENR("enr:-L24QOcdmsEln9seXi4gMXW59Pi9nP9nG1dcj2TJapaKMZc6TEpLNumnZVGYRM_W9fQ6ZKXCCFf6DNspVxl8aGtUykkCgmlkgnY0gmlwhLM99gSKbXVsdGlhZGRyc5YACASzPfYEBiMxAAoEsz32BAYjMt0DgnJzhQBIAQAgiXNlY3AyNTZrMaECsq2h0_gpY49mMm6tp3j9awr_xGCbxjoFXHp0m2aG8MKDdGNwgiMxg3VkcIIjM4V3YWt1Mg8"),
   		Addr:  wakutypes.MustDecodeMultiaddress("/ip4/179.61.246.4/tcp/9009/p2p/16Uiu2HAm7TA7Qq13C7DkXno4eFZGpxbm3thJasvZpgP6SBGSQQJd"),
   		Fleet: FleetWakuSandbox,
   	},
   	{
   		ID:    "node-02.iona.waku.sandbox",
   		ENR:   wakutypes.MustDecodeENR("enr:-L24QKLgpWYrASXFkCHVq08LkXBjmI7G6VlrIzr1wVxOnXISDVQPNoQSxwizTD-ypogX6TYEx1auvEsotqlRxKyYFGoCgmlkgnY0gmlwhMMjAG2KbXVsdGlhZGRyc5YACATDIwBtBiMxAAoEwyMAbQYjMt0DgnJzhQBIAQAgiXNlY3AyNTZrMaEC1d2IxoXQ58u6XuypxWJoevt4T6Gyjz0x2VEEK3OjzF2DdGNwgiMxg3VkcIIjM4V3YWt1Mg8"),
   		Addr:  wakutypes.MustDecodeMultiaddress("/ip4/195.35.0.109/tcp/9009/p2p/16Uiu2HAm9pWkCfdapHawsYrfRFrA7FuHV9jegujTfhSFjNWmn5fA"),
   		Fleet: FleetWakuSandbox,
   	},
   },
},
	FleetWakuTest: {
		WakuNodes: []string{
			"enrtree://AOGYWMBYOUIMOENHXCHILPKY3ZRFEULMFI4DOM442QSZ73TT2A7VI@test.waku.nodes.status.im",
		},
		DiscV5BootstrapNodes: []string{
			"enrtree://AOGYWMBYOUIMOENHXCHILPKY3ZRFEULMFI4DOM442QSZ73TT2A7VI@test.waku.nodes.status.im",
		},
		StoreNodes: []wakutypes.Mailserver{
			{
				ID:    "node-01.ac-cn-hongkong-c.waku.test",
				ENR:   wakutypes.MustDecodeENR("enr:-QEeuECvvBe6kIzHgMv_mD1YWQ3yfOfid2MO9a_A6ZZmS7E0FmAfntz2ZixAnPXvLWDJ81ARp4oV9UM4WXyc5D5USdEPAYJpZIJ2NIJpcIQI2ttrim11bHRpYWRkcnO4aAAxNixub2RlLTAxLmFjLWNuLWhvbmdrb25nLWMud2FrdS50ZXN0LnN0YXR1cy5pbQZ2XwAzNixub2RlLTAxLmFjLWNuLWhvbmdrb25nLWMud2FrdS50ZXN0LnN0YXR1cy5pbQYfQN4DgnJzkwABCAAAAAEAAgADAAQABQAGAAeJc2VjcDI1NmsxoQJIN4qwz3v4r2Q8Bv8zZD0eqBcKw6bdLvdkV7-JLjqIj4N0Y3CCdl-DdWRwgiMohXdha3UyDw"),
				Addr:  wakutypes.MustDecodeMultiaddress("/dns4/node-01.ac-cn-hongkong-c.waku.test.statusim.net/tcp/30303/p2p/16Uiu2HAkzHaTP5JsUwfR9NR8Rj9HC24puS6ocaU8wze4QrXr9iXp"),
				Fleet: FleetWakuTest,
			},
			{
				ID:    "node-01.do-ams3.waku.test",
				ENR:   wakutypes.MustDecodeENR("enr:-QEMuEDbayK340kH24XzK5FPIYNzWNYuH01NASNIb1skZfe_6l4_JSsG-vZ0LgN4Cgzf455BaP5zrxMQADHL5OQpbW6OAYJpZIJ2NIJpcISygI2rim11bHRpYWRkcnO4VgAoNiNub2RlLTAxLmRvLWFtczMud2FrdS50ZXN0LnN0YXR1cy5pbQZ2XwAqNiNub2RlLTAxLmRvLWFtczMud2FrdS50ZXN0LnN0YXR1cy5pbQYfQN4DgnJzkwABCAAAAAEAAgADAAQABQAGAAeJc2VjcDI1NmsxoQJATXRSRSUyTw_QLB6H_U3oziVQgNRgrXpK7wp2AMyNxYN0Y3CCdl-DdWRwgiMohXdha3UyDw"),
				Addr:  wakutypes.MustDecodeMultiaddress("/dns4/node-01.do-ams3.waku.test.statusim.net/tcp/30303/p2p/16Uiu2HAkykgaECHswi3YKJ5dMLbq2kPVCo89fcyTd38UcQD6ej5W"),
				Fleet: FleetWakuTest,
			},
			{
				ID:    "node-01.gc-us-central1-a.waku.test",
				ENR:   wakutypes.MustDecodeENR("enr:-QEeuEBO08GSjWDOV13HTf6L7iFoPQhv4S0-_Bd7Of3lFCBNBmpB9j6pGLedkX88KAXm6BFCS4ViQ_rLeDQuzj9Q6fs9AYJpZIJ2NIJpcIQiEAFDim11bHRpYWRkcnO4aAAxNixub2RlLTAxLmdjLXVzLWNlbnRyYWwxLWEud2FrdS50ZXN0LnN0YXR1cy5pbQZ2XwAzNixub2RlLTAxLmdjLXVzLWNlbnRyYWwxLWEud2FrdS50ZXN0LnN0YXR1cy5pbQYfQN4DgnJzkwABCAAAAAEAAgADAAQABQAGAAeJc2VjcDI1NmsxoQMIJwesBVgUiBCi8yiXGx7RWylBQkYm1U9dvEy-neLG2YN0Y3CCdl-DdWRwgiMohXdha3UyDw"),
				Addr:  wakutypes.MustDecodeMultiaddress("/dns4/node-01.gc-us-central1-a.waku.test.statusim.net/tcp/30303/p2p/16Uiu2HAmDCp8XJ9z1ev18zuv8NHekAsjNyezAvmMfFEJkiharitG"),
				Fleet: FleetWakuTest,
			},
		},
	},
}

var defaultPushNotificationServers = []string{
	"401ba5eda402678dc78a0a40fd0795f4ea8b1e34972c4d15cf33ac01292341c89f0cbc637fa9f7a3ffe0b9dfe90e9cdae7a14925500ab01b6a91c67bae42a97a",
	"181141b1d111908aaf05f4788e6778ec07073a1d4e1ce43c73815c40ee4e7345a1cbf5a90a45f601bf3763f12be63b01624ba1f36eeb9572455e7034b8f9f2c4",
	"5ffc34d5ffda180d94cd3974d9ed2bb082ede68f342babdbe801ceffb7da902087d43f9aa961c7b85029358874c08ef04ecad9f1d95a1f0e448cbdd5d04350c7",
}

func loadFleetsFromFile(filepath string) (FleetsMap, error) {
	// Read the JSON file to populate the supportedFleets map
	file, err := os.Open(filepath)
	if err != nil {
		err = pkgerrors.Wrap(err, "failed to open fleets json file")
		return nil, err
	}

	defer file.Close()

	var overrideFleets FleetsMap
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&overrideFleets)
	if err != nil {
		err = pkgerrors.Wrap(err, "failed to decode fleets json file")
		return nil, err
	}

	return overrideFleets, nil
}

func LoadFleetsFromFile(filepath string) error {
	fleetsMap, err := loadFleetsFromFile(filepath)
	if err != nil {
		return err
	}

	supportedFleets = fleetsMap
	return nil
}

func DefaultWakuNodes(fleet string) []string {
	return supportedFleets[fleet].WakuNodes
}

func DefaultDiscV5Nodes(fleet string) []string {
	return supportedFleets[fleet].DiscV5BootstrapNodes
}

func DefaultClusterID(fleet string) uint16 {
	return supportedFleets[fleet].ClusterID
}

func IsFleetSupported(fleet string) bool {
	_, ok := supportedFleets[fleet]
	return ok
}

func GetSupportedFleets() FleetsMap {
	return supportedFleets
}

func DefaultStoreNodes(fleet string) []wakutypes.Mailserver {
	return supportedFleets[fleet].StoreNodes
}

func DefaultPushNotificationServers() []string {
	return defaultPushNotificationServers
}

func DefaultClusterConfig(fleet string) ClusterConfig {
	return ClusterConfig{
		Enabled:                  true,
		Fleet:                    fleet,
		WakuNodes:                DefaultWakuNodes(fleet),
		DiscV5BootstrapNodes:     DefaultDiscV5Nodes(fleet),
		ClusterID:                DefaultClusterID(fleet),
		PushNotificationsServers: DefaultPushNotificationServers(),
	}
}
