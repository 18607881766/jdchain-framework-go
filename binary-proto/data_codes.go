package binary_proto

/*
 * Author: imuge
 * Date: 2020/5/26 上午8:51
 */

const (
	MERKLE_SNAPSHOT = 0x070

	MERKLE_DATA = 0x071

	MERKLE_LEAF = 0x072

	MERKLE_PATH = 0x073

	BYTES_VALUE = 0x080

	BYTES_VALUE_LIST = 0x081

	BLOCK_CHAIN_IDENTITY = 0x090

	BLOCK = 0x100

	BLOCK_BODY = 0x110

	BLOCK_GENESIS = 0x120

	DATA_SNAPSHOT = 0x130

	LEDGER_ADMIN_INFO = 0x131

	TX = 0x200

	TX_LEDGER = 0x201

	TX_CONTENT = 0x210

	TX_CONTENT_BODY = 0x220

	TX_RETURN_MESSAGE = 0x230

	TX_SET = 0x240

	TX_OP = 0x300

	TX_OP_LEDGER_INIT = 0x301

	TX_OP_USER_REG         = 0x310
	TX_OP_USER_INFO_SET    = 0x311
	TX_OP_USER_INFO_SET_KV = 0x312

	TX_OP_DATA_ACC_REG    = 0x320
	TX_OP_DATA_ACC_SET    = 0x321
	TX_OP_DATA_ACC_SET_KV = 0x322

	TX_OP_CONTRACT_DEPLOY = 0x330
	TX_OP_CONTRACT_UPDATE = 0x331

	TX_OP_CONTRACT_EVENT_SEND = 0x340

	TX_OP_PARTICIPANT_REG          = 0x350
	TX_OP_PARTICIPANT_STATE_UPDATE = 0x351

	TX_RESPONSE = 0x360

	TX_OP_RESULT = 0x370

	TX_OP_ROLE_CONFIGURE = 0x371

	TX_OP_ROLE_CONFIGURE_ENTRY = 0x372

	TX_OP_USER_ROLES_AUTHORIZE = 0x373

	TX_OP_USER_ROLE_AUTHORIZE_ENTRY = 0x374

	// enum types of permissions;
	ENUM_TX_PERMISSION      = 0x401
	ENUM_LEDGER_PERMISSION  = 0x402
	ENUM_MULTI_ROLES_POLICY = 0x403
	PRIVILEGE_SET = 0x410
	ROLE_SET = 0x411
	SECURITY_INIT_SETTING = 0x420
	SECURITY_ROLE_INIT_SETTING = 0x421
	SECURITY_USER_AUTH_INIT_SETTING = 0x422

	// contract types of metadata;
	METADATA    = 0x600
	METADATA_V2 = 0x601
	METADATA_INIT_SETTING = 0x610
	METADATA_INIT_PROPOSAL = 0x611
	METADATA_INIT_DECISION = 0x612
	METADATA_LEDGER_SETTING = 0x620
	METADATA_CONSENSUS_PARTICIPANT = 0x621
	METADATA_CONSENSUS_SETTING = 0x631
	METADATA_PARTICIPANT_INFO = 0x640
	METADATA_PARTICIPANT_STATE_INFO = 0x641
	METADATA_CRYPTO_SETTING = 0x642
	METADATA_CRYPTO_SETTING_PROVIDER = 0x643

	//	ACCOUNT = 0x700;

	ACCOUNT_HEADER = 0x710
	USER_ACCOUNT_HEADER = 0x800
	USER_INFO = 0x801
	DATA = 0x900

	// contract related;
	CONTRACT_ACCOUNT_HEADER = 0xA00
	CONTRACT_INFO = 0xA01

	// ...0xA19
	HASH = 0xB00
	HASH_OBJECT = 0xB10
	ENUM_TYPE = 0xB20
	CRYPTO_ALGORITHM = 0xB21
	ENUM_TYPE_TRANSACTION_STATE = 0xB22
	ENUM_TYPE_BYTES_VALUE_TYPE = 0xB23
	ENUM_TYPE_PARTICIPANT_NODE_STATE = 0xB24
	DIGITALSIGNATURE = 0xB30
	DIGITALSIGNATURE_BODY = 0xB31
	CLIENT_IDENTIFICATION = 0xC00
	CLIENT_IDENTIFICATIONS = 0xC10
	REQUEST = 0xD00
	REQUEST_NODE = 0xD10
	REQUEST_ENDPOINT = 0xD20

	// ------------------ 共识相关 ----------------

	CONSENSUS = 0x1000
	CONSENSUS_ACTION_REQUEST = CONSENSUS | 0x01
	CONSENSUS_ACTION_RESPONSE = CONSENSUS | 0x02
	CONSENSUS_SETTINGS = CONSENSUS | 0x03
	CONSENSUS_NODE_SETTINGS = CONSENSUS | 0x04
	CONSENSUS_CLI_INCOMING_SETTINGS = CONSENSUS | 0x05

	// ------------------ 共识相关（BFTSMART） ----------------
	CONSENSUS_BFTSMART = 0x1100
	CONSENSUS_BFTSMART_SETTINGS = CONSENSUS_BFTSMART | 0x01
	CONSENSUS_BFTSMART_NODE_SETTINGS = CONSENSUS_BFTSMART | 0x02
	CONSENSUS_BFTSMART_CLI_INCOMING_SETTINGS = CONSENSUS_BFTSMART | 0x03
	CONSENSUS_BFTSMART_BLOCK_SETTINGS = CONSENSUS_BFTSMART | 0x04

	// ------------------ 共识相关（MSGQUEUE） ----------------
	CONSENSUS_MSGQUEUE = 0x1200
	CONSENSUS_MSGQUEUE_SETTINGS = CONSENSUS_MSGQUEUE | 0x01
	CONSENSUS_MSGQUEUE_NODE_SETTINGS = CONSENSUS_MSGQUEUE | 0x02
	CONSENSUS_MSGQUEUE_CLI_INCOMING_SETTINGS = CONSENSUS_MSGQUEUE | 0x03
	CONSENSUS_MSGQUEUE_NETWORK_SETTINGS = CONSENSUS_MSGQUEUE | 0x04
	CONSENSUS_MSGQUEUE_BLOCK_SETTINGS = CONSENSUS_MSGQUEUE | 0x05
)
