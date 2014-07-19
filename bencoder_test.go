package bencodius

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestDecodingInt(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Decode("i2e"), BencodeInt(2), "it should decode ints")
	assert.Equal(Decode("i-4e"), BencodeInt(-4), "it should decode negatives")
}

func TestDecodingString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Decode("4:girl"), BencodeString("girl"), "it should decode strings")
	assert.Equal(Decode("0:"), BencodeString(""), "it should decode empty strings")
}

func TestDecodingList(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Decode("l4:girle"), BencodeList([]BencodeValue{BencodeString("girl")}), "it should decode lists")
	assert.Equal(Decode("le"), BencodeList([]BencodeValue{}), "it should decode empty lists")
}

func TestDecodingDict(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Decode("d2:lai2e1:zd1:aleee"), BencodeDict(
		map[BencodeString]BencodeValue{
			BencodeString("la"): BencodeInt(2),
			BencodeString("z"): BencodeDict(
				map[BencodeString]BencodeValue{BencodeString("a"): BencodeList([]BencodeValue{})})}),
		"it should decode dicts")
	assert.Equal(Decode("de"), BencodeDict(map[BencodeString]BencodeValue{}))
}

func TestDecoding(t *testing.T) {
	assert := assert.New(t)
	data, _ := ioutil.ReadFile("test/leaves.torrent")
	assert.Equal(Decode(string(data)), BencodeDict(map[BencodeString]BencodeValue{
		BencodeString("announce"): BencodeString("http://tracker.thepiratebay.org/announce"),
		BencodeString("announce-list"): BencodeList([]BencodeValue{
			BencodeList([]BencodeValue{BencodeString("http://tracker.thepiratebay.org/announce")}),
			BencodeList([]BencodeValue{BencodeString("udp://tracker.openbittorrent.com:80")}),
			BencodeList([]BencodeValue{BencodeString("udp://tracker.ccc.de:80")}),
			BencodeList([]BencodeValue{BencodeString("udp://tracker.publicbt.com:80")}),
			BencodeList([]BencodeValue{BencodeString("udp://fr33domtracker.h33t.com:3310/announce"),
				BencodeString("http://tracker.bittorrent.am/announce")})}),
		BencodeString("comment"):       BencodeString("Downloaded from http://TheTorrent.org"),
		BencodeString("created by"):    BencodeString("uTorrent/3300"),
		BencodeString("creation date"): BencodeInt(1375363666),
		BencodeString("encoding"):      BencodeString("UTF-8"),
		BencodeString("info"): BencodeDict(map[BencodeString]BencodeValue{
			BencodeString("length"):       BencodeInt(362017),
			BencodeString("name"):         BencodeString("Leaves of Grass by Walt Whitman.epub"),
			BencodeString("piece length"): BencodeInt(16384),
			BencodeString("pieces"):       BencodeString("\x1f\x9c?Y\xbe\xec\a\x97\x15\xecS2Kޅi䠴\xeb\xecB0}L\xe5U{]9d\xc5\xefU\xd3T\xcfJn\xcc{\xf1\xbc\xafy\xd1\x1f\xa5\xe0\xbe\x06Y<\x8f\xaa\xfc\f+\xa2\xcfv\xd7\x1c[\x01Rk#\x00\u007f\x9e\x99)\xbe\xaf\xc5\x15\x1ee\x11\t1\xa1\xb4L!\xbf\x1eh\xb9\x13\x8f\x90I^i\r\xbcU\xf5r\xe4\u0094L\xba\xcf&泮\x8ar)؊\xaf\xa0_a\xea\xaej\xbf?\a\xcbm\xb9g|ƭ\xedMӘ^E\x86'V\u007f\xa7c\x9f\x06_q\xb1\x89T0J\xcacfr\x9e\vGs\xd7z\xe8\f\xaa\x96\xa5$\x80M\xfeK\x9b\xd3ޮ\xf9\x99\xc9\xddQ\x02tgQ\x9d^\xb2V\x1a\xe2\xcc\x01F}\xe5\xf6C\r`\xbc\xba$yv\x92\xef\xa8w\r#\xdf\r\x83\r\x91\xcb5\xb3@z\x88\xba\xa0Y\r\xc8ɪj\x12\x0f'Cg\xdc\xd8g莃8\xc5r\xa0n<\x80\x1b)\xf5\x19\xdfS+>v\xf6p\xcfj\xeeS\x10\u007f=97\x84\x83\xf6\x9c\xf8\x0f\xa5h\xb1\xea\xc5;PaY\xe9\x88ؼ\x16\x92-\x12]w\xd8\x03\xd6R\xc3\xca0p\xc1n\xed\x91r\xabPm \xe5\"\xea?\x1a\xb6t\xb3\xf9#\xd7o\xe8\xf4O\xf3.7,;7ed\xc6\xfb_\r\xbeR\x16O\x03b\x9f\xd12&6\xba\xbb,\x01K}\xaeX-\xa4\x13c\x96Ra\xe6\xce\x12\xb47\x01\xf0\xa8\xc9\xed\x15 \xa7\x0e\xba\x00D\x00\xa2gv_m=\xd5Ǿ\xb5\xbd<u\xf3\xdf*TV\ra\x80\x11G\xfaN\xc7\xcfV\x8ep:\xcb\x04\xe5a\rMV\xdc\xc2B\xd02\x93\xe9Dl\xf5\xe4W\xd8\xeb=\x95\x88\xfd\x90Ƙޛ\r\xad\x92\x98\t\x06\xc0&\xd8\xc1@\x8f\xa0\x8f\xe4\xec")})}),
		"it decodes real files")
}

func TestEncoding(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Encode(BencodeInt(2)), "i2e", "it encodes integers")
	assert.Equal(Encode(BencodeString("girl")), "4:girl", "it encodes strings")
	assert.Equal(Encode(BencodeDict(map[BencodeString]BencodeValue{
		BencodeString("e"): BencodeString("?")})), "d1:e1:?e", "it encodes dicts")
	assert.Equal(Encode(BencodeList([]BencodeValue{BencodeString("woman")})), "l5:womane")
}
