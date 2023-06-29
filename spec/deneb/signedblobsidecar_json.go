// Copyright © 2023 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deneb

import (
	"encoding/json"
	"fmt"

	"github.com/attestantio/go-eth2-client/codecs"
	"github.com/pkg/errors"
)

// signedBlobSidecarJSON is the spec representation of the struct.
type signedBlobSidecarJSON struct {
	Message   *BlobSidecar `json:"message"`
	Signature string       `json:"signature"`
}

// MarshalJSON implements json.Marshaler.
func (s *SignedBlobSidecar) MarshalJSON() ([]byte, error) {
	return json.Marshal(&signedBlobSidecarJSON{
		Message:   s.Message,
		Signature: fmt.Sprintf("%#x", s.Signature),
	})
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *SignedBlobSidecar) UnmarshalJSON(input []byte) error {
	raw, err := codecs.RawJSON(&signedBlobSidecarJSON{}, input)
	if err != nil {
		return err
	}

	s.Message = &BlobSidecar{}
	if err := s.Message.UnmarshalJSON(raw["message"]); err != nil {
		return errors.Wrap(err, "message")
	}

	if err := s.Signature.UnmarshalJSON(raw["signature"]); err != nil {
		return errors.Wrap(err, "signature")
	}

	return nil
}