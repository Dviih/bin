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

//go:build !dviih_bin_kind_big

package bin

import (
	"github.com/Dviih/bin/kind"
	"math/big"
	"reflect"
)

func init() {
	i := kind.NewHandler(
		func(encoder kind.Encoder, value reflect.Value) error {
			return encoder.Encode(kind.Call(value, "Bytes")[0].Interface())
		},
		func(decoder kind.Decoder, value reflect.Value) error {
			var data []byte

			if err := decoder.Decode(&data); err != nil {
				return err
			}

			kind.Call(value, "SetBytes", reflect.ValueOf(data))
			return nil
		},
	)

	register(67, reflect.TypeFor[big.Int](), i)
	register(68, reflect.TypeFor[big.Float](), kind.Gob)
	register(69, reflect.TypeFor[big.Rat](), kind.Gob)
}
