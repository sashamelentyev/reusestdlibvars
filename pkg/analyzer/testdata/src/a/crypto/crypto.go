package crypto

import "fmt"

var (
	_ = "MD4"         // want `"MD4" can be replaced by crypto.MD4.String\(\)`
	_ = "MD5"         // want `"MD5" can be replaced by crypto.MD5.String\(\)`
	_ = "SHA-1"       // want `"SHA-1" can be replaced by crypto.SHA1.String\(\)`
	_ = "SHA-224"     // want `"SHA-224" can be replaced by crypto.SHA224.String\(\)`
	_ = "SHA-256"     // want `"SHA-256" can be replaced by crypto.SHA256.String\(\)`
	_ = "SHA-384"     // want `"SHA-384" can be replaced by crypto.SHA384.String\(\)`
	_ = "SHA-512"     // want `"SHA-512" can be replaced by crypto.SHA512.String\(\)`
	_ = "MD5+SHA1"    // want `"MD5\+SHA1" can be replaced by crypto.MD5SHA1.String\(\)`
	_ = "RIPEMD-160"  // want `"RIPEMD-160" can be replaced by crypto.RIPEMD160.String\(\)`
	_ = "SHA3-224"    // want `"SHA3-224" can be replaced by crypto.SHA3_224.String\(\)`
	_ = "SHA3-256"    // want `"SHA3-256" can be replaced by crypto.SHA3_256.String\(\)`
	_ = "SHA3-384"    // want `"SHA3-384" can be replaced by crypto.SHA3_384.String\(\)`
	_ = "SHA3-512"    // want `"SHA3-512" can be replaced by crypto.SHA3_512.String\(\)`
	_ = "SHA-512/224" // want `"SHA-512/224" can be replaced by crypto.SHA512_224.String\(\)`
	_ = "SHA-512/256" // want `"SHA-512/256" can be replaced by crypto.SHA512_256.String\(\)`
	_ = "BLAKE2s-256" // want `"BLAKE2s-256" can be replaced by crypto.BLAKE2s_256.String\(\)`
	_ = "BLAKE2b-256" // want `"BLAKE2b-256" can be replaced by crypto.BLAKE2b_256.String\(\)`
	_ = "BLAKE2b-384" // want `"BLAKE2b-384" can be replaced by crypto.BLAKE2b_384.String\(\)`
	_ = "BLAKE2b-512" // want `"BLAKE2b-512" can be replaced by crypto.BLAKE2b_512.String\(\)`
)

const (
	_ = "MD4"         // want `"MD4" can be replaced by crypto.MD4.String\(\)`
	_ = "MD5"         // want `"MD5" can be replaced by crypto.MD5.String\(\)`
	_ = "SHA-1"       // want `"SHA-1" can be replaced by crypto.SHA1.String\(\)`
	_ = "SHA-224"     // want `"SHA-224" can be replaced by crypto.SHA224.String\(\)`
	_ = "SHA-256"     // want `"SHA-256" can be replaced by crypto.SHA256.String\(\)`
	_ = "SHA-384"     // want `"SHA-384" can be replaced by crypto.SHA384.String\(\)`
	_ = "SHA-512"     // want `"SHA-512" can be replaced by crypto.SHA512.String\(\)`
	_ = "MD5+SHA1"    // want `"MD5\+SHA1" can be replaced by crypto.MD5SHA1.String\(\)`
	_ = "RIPEMD-160"  // want `"RIPEMD-160" can be replaced by crypto.RIPEMD160.String\(\)`
	_ = "SHA3-224"    // want `"SHA3-224" can be replaced by crypto.SHA3_224.String\(\)`
	_ = "SHA3-256"    // want `"SHA3-256" can be replaced by crypto.SHA3_256.String\(\)`
	_ = "SHA3-384"    // want `"SHA3-384" can be replaced by crypto.SHA3_384.String\(\)`
	_ = "SHA3-512"    // want `"SHA3-512" can be replaced by crypto.SHA3_512.String\(\)`
	_ = "SHA-512/224" // want `"SHA-512/224" can be replaced by crypto.SHA512_224.String\(\)`
	_ = "SHA-512/256" // want `"SHA-512/256" can be replaced by crypto.SHA512_256.String\(\)`
	_ = "BLAKE2s-256" // want `"BLAKE2s-256" can be replaced by crypto.BLAKE2s_256.String\(\)`
	_ = "BLAKE2b-256" // want `"BLAKE2b-256" can be replaced by crypto.BLAKE2b_256.String\(\)`
	_ = "BLAKE2b-384" // want `"BLAKE2b-384" can be replaced by crypto.BLAKE2b_384.String\(\)`
	_ = "BLAKE2b-512" // want `"BLAKE2b-512" can be replaced by crypto.BLAKE2b_512.String\(\)`
)

func _() {
	_ = fmt.Sprint("MD4") // want `"MD4" can be replaced by crypto.MD4.String\(\)`
	_ = fmt.Sprint("MD4 is cryptographic hash function")
}
