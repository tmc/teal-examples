package globals

var MinTxnFee uint64            // micro Algos
var MinBalance uint64           // micro Algos
var MaxTxnLife uint64           // rounds
var ZeroAddress []byte          // 32 byte address of all zero bytes
var GroupSize uint64            // Number of transactions in this atomic transaction group. At least 1
var LogicSigVersion uint64      // Maximum supported TEAL version. LogicSigVersion >= 2.
var Round uint64                // Current round number. LogicSigVersion >= 2.
var LatestTimestamp uint64      // Last confirmed block UNIX timestamp. Fails if negative. LogicSigVersion >= 2.
var CurrentApplicationID uint64 // ID of current application executing. Fails if no such application is executing. LogicSigVersion >= 2.
var CreatorAddress []byte       // Address of the creator of the current application. Fails if no such application is executing. LogicSigVersion >= 3.
