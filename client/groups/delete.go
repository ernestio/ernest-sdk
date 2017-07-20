/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package groups

import (
	"fmt"
)

// Delete : delete a group
func (g *Groups) Delete(id int) error {
	path := fmt.Sprintf(apiroute+"%d", id)

	resp, err := g.Conn.Delete(path)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
