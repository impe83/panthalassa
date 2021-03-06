package dapp

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	mh "github.com/multiformats/go-multihash"
	ed25519 "golang.org/x/crypto/ed25519"
)

var InvalidSignature = errors.New("failed to verify signature for DApp")

type SV struct {
	Major uint `json:"major"`
	Minor uint `json:"minor"`
	Patch uint `json:"patch"`
}

func (v *SV) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// JSON Representation of published DApp
type Data struct {
	Name           map[string]string
	UsedSigningKey ed25519.PublicKey `storm:"id"`
	Code           []byte
	Image          []byte
	Signature      []byte
	Engine         SV
	Version        int
}

// hash the published DApp
func (r Data) Hash() ([]byte, error) {

	buff := bytes.NewBuffer(nil)

	// sort languages
	var languages []string
	for k := range r.Name {
		languages = append(languages, k)
	}
	sort.Strings(languages)

	// write languages to buffer
	for _, k := range languages {
		// write language code
		if _, err := buff.WriteString(k); err != nil {
			return nil, err
		}
		// write name
		if _, err := buff.WriteString(r.Name[k]); err != nil {
			return nil, err
		}
	}

	// write used signing key into buffer
	if _, err := buff.Write(r.UsedSigningKey); err != nil {
		return nil, err
	}

	// write code into buffer
	if _, err := buff.Write(r.Code); err != nil {
		return nil, err
	}

	// write image into buffer
	if _, err := buff.Write(r.Image); err != nil {
		return nil, err
	}

	// write engine into buffer
	if _, err := buff.Write([]byte(r.Engine.String())); err != nil {
		return nil, err
	}

	// write version as string - see for the specs why
	if _, err := buff.WriteString(strconv.Itoa(int(r.Version))); err != nil {
		return nil, err
	}

	// hash it
	multiHash, err := mh.Sum(buff.Bytes(), mh.SHA2_256, -1)
	if err != nil {
		return nil, err
	}

	return multiHash, nil

}

// verify if this published DApp
// was signed with the attached public key
func (r Data) VerifySignature() (bool, error) {

	hash, err := r.Hash()
	if err != nil {
		return false, err
	}

	return ed25519.Verify(r.UsedSigningKey, hash, r.Signature), nil

}

func (r Data) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func engineVersionToSV(version string) (SV, error) {

	// split the version
	rawSVFragments := strings.Split(version, ".")
	if len(rawSVFragments) != 3 {
		return SV{}, errors.New("invalid version - a version must consist of 3 numbers separated by dots")
	}

	sVFragments := []uint{}

	// convert string fragments into numeric fragments
	for _, f := range rawSVFragments {
		numFragment, err := strconv.Atoi(f)
		if err != nil {
			return SV{}, err
		}
		sVFragments = append(sVFragments, uint(numFragment))
	}

	// make sure we have the right amount of fragments
	if len(sVFragments) != 3 {
		return SV{}, errors.New("invalid version - a version must consist of 3 numbers separated by dots")
	}

	v := SV{
		Major: sVFragments[0],
		Minor: sVFragments[1],
		Patch: sVFragments[2],
	}

	return v, nil

}

type RawData struct {
	Name           map[string]string `json:"name"`
	UsedSigningKey string            `json:"used_signing_key"`
	Code           string            `json:"code"`
	Image          string            `json:"image"`
	Signature      string            `json:"signature"`
	Engine         string            `json:"engine"`
	Version        string            `json:"version"`
}

func ParseJsonToData(b RawData) (Data, error) {

	// unmarshal used signing key
	usedSigningKey, err := hex.DecodeString(b.UsedSigningKey)
	if err != nil {
		return Data{}, err
	}
	if len(usedSigningKey) != 32 {
		return Data{}, fmt.Errorf("invalid length fof signing key %d", len(usedSigningKey))
	}

	// unmarshal signature
	rawSignature, err := hex.DecodeString(b.Signature)
	if err != nil {
		return Data{}, err
	}

	// decode engine version
	sv, err := engineVersionToSV(b.Engine)
	if err != nil {
		return Data{}, err
	}

	// decode image from base64 to bytes
	image, err := base64.StdEncoding.DecodeString(b.Image)
	if err != nil {
		return Data{}, err
	}

	// turn string version into integer
	v, err := strconv.Atoi(b.Version)
	if err != nil {
		return Data{}, err
	}

	return Data{
		Name:           b.Name,
		UsedSigningKey: usedSigningKey,
		Code:           []byte(b.Code),
		Image:          image,
		Signature:      rawSignature,
		Engine:         sv,
		Version:        v,
	}, nil

}
