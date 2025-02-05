/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.4.7]

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package stader

import (
	"fmt"
)

const (
	colorReset  string = "\033[0m"
	colorYellow string = "\033[33m"
)

// Print a warning about the gas estimate for operations that have multiple transactions
func (sd *Client) PrintMultiTxWarning() {
	fmt.Printf("%sNOTE: This operation requires multiple transactions.\n%s",
		colorYellow,
		colorReset)

}
