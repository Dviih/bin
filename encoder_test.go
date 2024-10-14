/*
 *     A tiny format for using binary data
 *     Copyright (C) 2024  Dviih
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Affero General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Affero General Public License for more details.
 *
 *     You should have received a copy of the GNU Affero General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 */

package bin

import (
	"testing"
)

func TestEncoderNil(t *testing.T) {
	t.Parallel()

	s := &stream{}
	encoder := NewEncoder(s)

	if err := encoder.Encode(Nil); err != nil {
		t.Error("failed to encode nil")
	}

	if string(s.Data) != string(expectedNil) {
		t.Errorf("expected %v, received: %v", expectedNil, s.Data)
	}
}

