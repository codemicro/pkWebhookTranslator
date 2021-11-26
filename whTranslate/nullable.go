/*
 *  pkWebhookTranslate, https://github.com/codemicro/pkWebhookTranslate
 *  Copyright (c) 2021 codemicro and contributors
 *
 *  SPDX-License-Identifier: BSD-2-Clause
 */

package whtranslate

import (
	"bytes"
	"encoding/json"
)

var nullBytes = []byte("null")

type nullableString struct {
	Value    string
	HasValue bool
}

func (n *nullableString) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullBytes) {
		n.HasValue = true
		return nil
	}

	if err := json.Unmarshal(b, &n.Value); err != nil {
		return err
	}

	n.HasValue = true
	return nil
}
