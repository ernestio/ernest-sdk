/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package client

// Authenticate : authenticate against an ernest instance
func (c *Client) Authenticate() (string, error) {
	return c.Conn.Authenticate()
}
