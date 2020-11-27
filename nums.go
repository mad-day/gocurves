// CC0 1.0 Universal

package gocurves

import "crypto/elliptic"
import "math/big"

func strbig(s string) (i *big.Int) {
	i = new(big.Int)
	i.SetString(s,0)
	return
}

var numsp256d1 = &elliptic.CurveParams{
	P: strbig("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff43"),
	N: strbig("0xffffffffffffffffffffffffffffffffe43c8275ea265c6020ab20294751a825"),
	B: strbig("0x25581"),
	Gx: strbig("0x01"),
	Gy: strbig("0x696f1853c1e466d7fc82c96cceeedd6bd02c2f9375894ec10bf46306c2b56c77"),
	BitSize: 256,
	Name: "numsp256d1",
}
var numsp384d1 = &elliptic.CurveParams{
	P: strbig("0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec3"),
	N: strbig("0xffffffffffffffffffffffffffffffffffffffffffffffffd61eaf1eeb5d6881beda9d3d4c37e27a604d81f67b0e61b9"),
	B: strbig("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff77bb"),
	Gx: strbig("0x02"),
	Gy: strbig("0x3c9f82cb4b87b4dc71e763e0663e5dbd8034ed422f04f82673330dc58d15ffa2b4a3d0bad5d30f865bcbbf503ea66f43"),
	BitSize: 384,
	Name: "numsp384d1",
}
var numsp512d1 = &elliptic.CurveParams{
	P: strbig("0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc7"),
	N: strbig("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b3ca4fb94e7831b4fc258ed97d0bdc63b568b36607cd243ce153f390433555d"),
	B: strbig("0x1d99b"),
	Gx: strbig("0x02"),
	Gy: strbig("0x1c282eb23327f9711952c250ea61ad53fcc13031cf6dd336e0b9328433afbdd8cc5a1c1f0c716fdc724dde537c2b0adb00bb3d08dc83755b205cc30d7f83cf28"),
	BitSize: 512,
	Name: "numsp512d1",
}

// Nums256() returns a Curve which implements nums256d1 of Microsoft's Nothing Up My Sleeve (NUMS)
func Nums256() elliptic.Curve { return numsp256d1 }

// Nums384() returns a Curve which implements nums384d1 of Microsoft's Nothing Up My Sleeve (NUMS)
func Nums384() elliptic.Curve { return numsp384d1 }

// Nums512() returns a Curve which implements nums512d1 of Microsoft's Nothing Up My Sleeve (NUMS)
func Nums512() elliptic.Curve { return numsp512d1 }



var bp160t1 = &elliptic.CurveParams{
	P: strbig("0xe95e4a5f737059dc60dfc7ad95b3d8139515620f"),
	N: strbig("0xe95e4a5f737059dc60df5991d45029409e60fc09"),
	B: strbig("0x7a556b6dae535b7b51ed2c4d7daa7a0b5c55f380"),
	Gx: strbig("0xb199b13b9b34efc1397e64baeb05acc265ff2378"),
	Gy: strbig("0xadd6718b7c7c1961f0991b842443772152c9e0ad"),
	BitSize: 160,
	Name: "brainpoolP160t1",
}
var bp192t1 = &elliptic.CurveParams{
	P: strbig("0xc302f41d932a36cda7a3463093d18db78fce476de1a86297"),
	N: strbig("0xc302f41d932a36cda7a3462f9e9e916b5be8f1029ac4acc1"),
	B: strbig("0x13d56ffaec78681e68f9deb43b35bec2fb68542e27897b79"),
	Gx: strbig("0x3ae9e58c82f63c30282e1fe7bbf43fa72c446af6f4618129"),
	Gy: strbig("0x97e2c5667c2223a902ab5ca449d0084b7e5b3de7ccc01c9"),
	BitSize: 192,
	Name: "brainpoolP192t1",
}
var bp224t1 = &elliptic.CurveParams{
	P: strbig("0xd7c134aa264366862a18302575d1d787b09f075797da89f57ec8c0ff"),
	N: strbig("0xd7c134aa264366862a18302575d0fb98d116bc4b6ddebca3a5a7939f"),
	B: strbig("0x4b337d934104cd7bef271bf60ced1ed20da14c08b3bb64f18a60888d"),
	Gx: strbig("0x6ab1e344ce25ff3896424e7ffe14762ecb49f8928ac0c76029b4d580"),
	Gy: strbig("0x374e9f5143e568cd23f3f4d7c0d4b1e41c8cc0d1c6abd5f1a46db4c"),
	BitSize: 224,
	Name: "brainpoolP224t1",
}
var bp256t1 = &elliptic.CurveParams{
	P: strbig("0xa9fb57dba1eea9bc3e660a909d838d726e3bf623d52620282013481d1f6e5377"),
	N: strbig("0xa9fb57dba1eea9bc3e660a909d838d718c397aa3b561a6f7901e0e82974856a7"),
	B: strbig("0x662c61c430d84ea4fe66a7733d0b76b7bf93ebc4af2f49256ae58101fee92b04"),
	Gx: strbig("0xa3e8eb3cc1cfe7b7732213b23a656149afa142c47aafbc2b79a191562e1305f4"),
	Gy: strbig("0x2d996c823439c56d7f7b22e14644417e69bcb6de39d027001dabe8f35b25c9be"),
	BitSize: 256,
	Name: "brainpoolP256t1",
}
var bp320t1 = &elliptic.CurveParams{
	P: strbig("0xd35e472036bc4fb7e13c785ed201e065f98fcfa6f6f40def4f92b9ec7893ec28fcd412b1f1b32e27"),
	N: strbig("0xd35e472036bc4fb7e13c785ed201e065f98fcfa5b68f12a32d482ec7ee8658e98691555b44c59311"),
	B: strbig("0xa7f561e038eb1ed560b3d147db782013064c19f27ed27c6780aaf77fb8a547ceb5b4fef422340353"),
	Gx: strbig("0x925be9fb01afc6fb4d3e7d4990010f813408ab106c4f09cb7ee07868cc136fff3357f624a21bed52"),
	Gy: strbig("0x63ba3a7a27483ebf6671dbef7abb30ebee084e58a0b077ad42a5a0989d1ee71b1b9bc0455fb0d2c3"),
	BitSize: 320,
	Name: "brainpoolP320t1",
}
var bp384t1 = &elliptic.CurveParams{
	P: strbig("0x8cb91e82a3386d280f5d6f7e50e641df152f7109ed5456b412b1da197fb71123acd3a729901d1a71874700133107ec53"),
	N: strbig("0x8cb91e82a3386d280f5d6f7e50e641df152f7109ed5456b31f166e6cac0425a7cf3ab6af6b7fc3103b883202e9046565"),
	B: strbig("0x7f519eada7bda81bd826dba647910f8c4b9346ed8ccdc64e4b1abd11756dce1d2074aa263b88805ced70355a33b471ee"),
	Gx: strbig("0x18de98b02db9a306f2afcd7235f72a819b80ab12ebd653172476fecd462aabffc4ff191b946a5f54d8d0aa2f418808cc"),
	Gy: strbig("0x25ab056962d30651a114afd2755ad336747f93475b7a1fca3b88f2b6a208ccfe469408584dc2b2912675bf5b9e582928"),
	BitSize: 384,
	Name: "brainpoolP384t1",
}
var bp512t1 = &elliptic.CurveParams{
	P: strbig("0xaadd9db8dbe9c48b3fd4e6ae33c9fc07cb308db3b3c9d20ed6639cca703308717d4d9b009bc66842aecda12ae6a380e62881ff2f2d82c68528aa6056583a48f3"),
	N: strbig("0xaadd9db8dbe9c48b3fd4e6ae33c9fc07cb308db3b3c9d20ed6639cca70330870553e5c414ca92619418661197fac10471db1d381085ddaddb58796829ca90069"),
	B: strbig("0x7cbbbcf9441cfab76e1890e46884eae321f70c0bcb4981527897504bec3e36a62bcdfa2304976540f6450085f2dae145c22553b465763689180ea2571867423e"),
	Gx: strbig("0x640ece5c12788717b9c1ba06cbc2a6feba85842458c56dde9db1758d39c0313d82ba51735cdb3ea499aa77a7d6943a64f7a3f25fe26f06b51baa2696fa9035da"),
	Gy: strbig("0x5b534bd595f5af0fa2c892376c84ace1bb4e3019b71634c01131159cae03cee9d9932184beef216bd71df2dadf86a627306ecff96dbb8bace198b61e00f8b332"),
	BitSize: 512,
	Name: "brainpoolP512t1",
}


// Bp160() returns a Curve which implements brainpoolP160t1 of Brainpool Standard
func Bp160() elliptic.Curve { return bp160t1 }

// Bp192() returns a Curve which implements brainpoolP192t1 of Brainpool Standard
func Bp192() elliptic.Curve { return bp192t1 }

// Bp224() returns a Curve which implements brainpoolP224t1 of Brainpool Standard
func Bp224() elliptic.Curve { return bp224t1 }

// Bp256() returns a Curve which implements brainpoolP256t1 of Brainpool Standard
func Bp256() elliptic.Curve { return bp256t1 }

// Bp320() returns a Curve which implements brainpoolP320t1 of Brainpool Standard
func Bp320() elliptic.Curve { return bp320t1 }

// Bp384() returns a Curve which implements brainpoolP384t1 of Brainpool Standard
func Bp384() elliptic.Curve { return bp384t1 }

// Bp512() returns a Curve which implements brainpoolP512t1 of Brainpool Standard
func Bp512() elliptic.Curve { return bp512t1 }




