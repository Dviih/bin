/*
 *     A tiny binary format
 *     Copyright (C) 2025  Dviih
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

//go:build !dviih_bin_kind_encoding

package bin

import (
	"encoding"
	"github.com/Dviih/bin/kind"
	"reflect"
)

func init() {
	register(65, reflect.TypeFor[encoding.BinaryMarshaler](), kind.EncodingBinary)
	mkind.Alias(65, reflect.TypeFor[encoding.BinaryUnmarshaler]())
}
