/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// PolicyDocument : stores policy document data
type PolicyDocument struct {
	ID         int    `json:"id"`
	PolicyID   int    `json:"policy_id"`
	Revision   int    `json:"revision"`
	Username   string `json:"username"`
	Definition string `json:"definition"`
	CreatedAt  string `json:"created_at"`
}
